package node

/*
 * @brief: 节点原数据
 * @param: nodeID: 节点id
 * @param: nodePort: 节点端口
 */
type NodeMetaData struct {
	nodeID   int
	nodePort string
}

/*
 * @brief: 节点结构体
 * @param: metaData: 节点元数据
 * @param: cache: 缓存数据
 */
type Node struct {
	MetaData NodeMetaData
	Cache    map[string][]string
}

/*
 * @brief: 创建一个节点
 * @param: id: 节点id
 * @param: port: 节点端口
 * @return: *Node: 节点指针
 */
func NewNode(id int, port string) *Node {
	return &Node{
		MetaData: NodeMetaData{nodeID: id, nodePort: port},
		Cache:    make(map[string][]string),
	}
}

/*
 * @brief: 添加一条cache
 * @param: key: cache的key
 * @param: value: cache的value
 * @return: int: 添加成功返回1，失败返回0
 */
func (n *Node) AddCache(key string, value []string) int {
	if _, ok := n.Cache[key]; ok {
		return 0
	} else {
		n.Cache[key] = value
		return 1
	}
}

/*
 * @brief: 删除一条cache
 * @param: key: cache的key
 * @return: int: 删除成功返回1，失败返回0
 */
func (n *Node) DelCache(key string) int {
	if _, ok := n.Cache[key]; ok {
		delete(n.Cache, key)
		return 1
	} else {
		return 0
	}
}

/*
 * @brief: 获取一条cache
 * @param: key: cache的key
 * @return: []string: 获取成功返回value，失败返回nil
 */
func (n *Node) GetCache(key string) []string {
	if value, ok := n.Cache[key]; ok {
		return value
	} else {
		return nil
	}
}

/*
 * @brief: 修改cache内容
 * @param: key: cache的key
 * @param: value: cache的value
 * @return: int: 修改成功返回1，失败返回0
 */
func (n *Node) SetCache(key string, value []string) int {
	n.Cache[key] = value
	return 1
}
