package httpagent

import (
	"SDCS/hash"
	"SDCS/node"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpAgent struct {
	id   int
	port string
	node *node.Node
}

func NewHttpAgent(id int, port string) *HttpAgent {
	return &HttpAgent{
		id:   id,
		port: port,
		node: node.NewNode(id, port),
	}
}

func (h *HttpAgent) StartHttpAgent() {
	engine := gin.Default()
	engine.Handle(http.MethodGet, "/:key", h.getCache)
	engine.Handle(http.MethodPost, "/", h.setCache)
	engine.Handle(http.MethodDelete, "/:key", h.delCache)
	engine.Run(":" + h.port)
}

func (h *HttpAgent) getCache(c *gin.Context) {
	key := c.Param("key")
	_, ok := hash.GetCacheNode(key)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "node not found",
		})
	} else {
		//if nodePort == h.port {
		if value := h.node.GetCache(key); value != nil {
			c.JSON(http.StatusOK, gin.H{
				key: value,
			})
		} else {
			c.Status(http.StatusNotFound)
		}
		//}
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
		_, ok := hash.GetCacheNode(key)
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "node not found",
			})
		} else {
			//if nodePort == h.port {
			if valueStr, ok := value.(string); ok {
				valueSlice := []string{valueStr}
				if res := h.node.SetCache(key, valueSlice); res == 1 {
					c.Status(http.StatusOK)
				}
			} else {
				valueSlice := []string{}
				for _, val := range value.([]interface{}) {
					valueSlice = append(valueSlice, val.(string))
				}
				if res := h.node.SetCache(key, valueSlice); res == 1 {
					c.Status(http.StatusOK)
				}
			}
			//}
		}
	}
}

func (h *HttpAgent) delCache(c *gin.Context) {
	key := c.Param("key")
	_, ok := hash.GetCacheNode(key)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "node not found",
		})
	} else {
		//if nodePort == h.port {
		res := h.node.DelCache(key)
		c.JSON(http.StatusOK, res)
		//}
	}
}
