package httpagent

import (
	"SDCS/hash"
	"SDCS/node"
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	pb "SDCS/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	// http端口与rpc端口的对应
	// 9527 -> 30001
	// 9528 -> 30002
	// 9529 -> 30003
	mapPortRpcPort = map[string]string{
		"9527": "30001",
		"9528": "30002",
		"9529": "30003",
	}
)

type HttpAgent struct {
	id      int
	port    string
	rpcPort string
	node    *node.Node
}

type rpcserver struct {
	h *HttpAgent
	pb.UnimplementedCacheServiceServer
}

func (s *rpcserver) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	key := req.Key
	value := s.h.node.GetCache(key)
	return &pb.GetResponse{
		Value: value,
	}, nil
}

func (s *rpcserver) Set(ctx context.Context, req *pb.SetRequest) (*pb.SetResponse, error) {
	key := req.Key
	value := req.Value
	s.h.node.SetCache(key, value)
	return &pb.SetResponse{}, nil
}

func (s *rpcserver) Del(ctx context.Context, req *pb.DelRequest) (*pb.DelResponse, error) {
	key := req.Key
	count := int32(s.h.node.DelCache(key))
	return &pb.DelResponse{
		DelCount: count,
	}, nil
}

func NewHttpAgent(id int, port string, rpcPort string) *HttpAgent {
	return &HttpAgent{
		id:      id,
		port:    port,
		rpcPort: rpcPort,
		node:    node.NewNode(id, port),
	}
}

func startGrpcServer(p string, h *HttpAgent) {
	lis, err := net.Listen("tcp", ":"+p)
	if err != nil {
		log.Fatalf("fail to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCacheServiceServer(s, &rpcserver{
		h: h,
	})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("fail to serve: %v", err)
	}
}

func (h *HttpAgent) StartHttpAgent() {
	engine := gin.Default()
	engine.Handle(http.MethodGet, "/:key", h.getCache)
	engine.Handle(http.MethodPost, "/", h.setCache)
	engine.Handle(http.MethodDelete, "/:key", h.delCache)

	// 协程，启动grpc服务端
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		startGrpcServer(mapPortRpcPort[h.port], h)
	}()

	// server := &http.Server{Handler: engine}
	// l, err := net.Listen("tcp4", ":"+h.port)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = server.Serve(l)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	engine.Run("0.0.0.0:" + h.port)

}

func (h *HttpAgent) getCache(c *gin.Context) {
	key := c.Param("key")
	nodePort, ok := hash.GetCacheNode(key)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "node not found",
		})
	} else {
		if nodePort == h.port {
			if value := h.node.GetCache(key); value != nil {
				//判断value长度
				if len(value) == 1 {
					c.JSON(http.StatusOK, gin.H{
						key: value[0],
					})
					return
				} else {
					c.JSON(http.StatusOK, gin.H{
						key: value,
					})
					return
				}
			} else {
				c.Status(http.StatusNotFound)
			}
		} else {
			conn, err := grpc.Dial("127.0.0.1:"+mapPortRpcPort[nodePort], grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			connect := pb.NewCacheServiceClient(conn)
			r, err := connect.Get(context.Background(), &pb.GetRequest{Key: key})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			if r.Value != nil {
				if len(r.Value) == 1 {
					c.JSON(http.StatusOK, gin.H{
						key: r.Value[0],
					})
					return
				} else {
					c.JSON(http.StatusOK, gin.H{
						key: r.Value,
					})
					return
				}
			} else {
				c.Status(http.StatusNotFound)
			}

		}
	}
}

func (h *HttpAgent) setCache(c *gin.Context) {
	jsonMap := make(map[string]interface{})
	if err := c.ShouldBindJSON(&jsonMap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid json",
		})
		return
	}
	for key, value := range jsonMap {
		nodePort, ok := hash.GetCacheNode(key)
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "node not found",
			})
		} else {
			valueSlice := []string{}
			if valueStr, ok := value.(string); ok {
				valueSlice = append(valueSlice, valueStr)
			} else {
				for _, val := range value.([]interface{}) {
					valueSlice = append(valueSlice, val.(string))
				}
			}
			if nodePort == h.port {
				if res := h.node.SetCache(key, valueSlice); res == 1 {
					c.Status(http.StatusOK)
				}
			} else {
				conn, err := grpc.Dial("127.0.0.1:"+mapPortRpcPort[nodePort], grpc.WithTransportCredentials(insecure.NewCredentials()))
				if err != nil {
					log.Fatalf("did not connect: %v", err)
				}
				defer conn.Close()
				connect := pb.NewCacheServiceClient(conn)
				_, err = connect.Set(context.Background(), &pb.SetRequest{Key: key, Value: valueSlice})
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}
				c.Status(http.StatusOK)
			}
		}
	}
}

func (h *HttpAgent) delCache(c *gin.Context) {
	key := c.Param("key")
	nodePort, ok := hash.GetCacheNode(key)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "node not found",
		})
	} else {
		if nodePort == h.port {
			res := h.node.DelCache(key)
			c.JSON(http.StatusOK, res)
		} else {
			conn, err := grpc.Dial("127.0.0.1:"+mapPortRpcPort[nodePort], grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			connect := pb.NewCacheServiceClient(conn)
			r, err := connect.Del(context.Background(), &pb.DelRequest{Key: key})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			c.JSON(http.StatusOK, r.DelCount)
		}
	}
}
