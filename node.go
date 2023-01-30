package tiresearch

// Node
/**
 *  @Author: rym 2022-11-16 09:01:35
 *  @Description: 数据结构节点对象
 */
type Node struct {
	//  Char
	/**
	 *  @Author: rym 2022-11-16 09:01:47
	 *  @Description: 该节点的字符
	 */
	Char string

	//  Code
	/**
	 *  @Author: rym 2022-11-16 09:08:38
	 *  @Description: 改节点的unicode
	 */
	Code rune

	//  Children
	/**
	 *  @Author: rym 2022-11-16 09:09:00
	 *  @Description: 该节点的子节点字典
	 */
	Children map[rune]*Node `json:"children"`

	//  ContentList
	/**
	 *  @Author: rym 2022-11-16 09:10:54
	 *  @Description:
	 */
	Content []string `json:"content"`
}

// newTrieNode
/**
 *  @Author: rym 2022-11-16 17:18:08
 *  @Description: 初始化节点
 *  @param char
 *  @return *Node
 */
func newTrieNode(char string) *Node {
	return &Node{
		Char:     char,
		Children: make(map[rune]*Node),
		Content:  []string{},
	}
}
