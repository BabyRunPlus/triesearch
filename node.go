package triesearch

// Node
/**
 *  @Author: rym 2022-11-16 09:01:35
 *  @Description: 数据结构节点对象
 */
type Node struct {
	//  Children
	/**
	 *  @Author: rym 2022-11-16 09:09:00
	 *  @Description: 该节点的子节点字典
	 */
	Children map[rune]*Node `json:"children"`

	//  ContentMap
	/**
	 *  @Author: rym 2023-03-09 16:00:51
	 *  @Description: 临时用的map，用来作数据去重
	 */
	ContentMap map[string]string

	//  Char
	/**
	 *  @Author: rym 2022-11-16 09:01:47
	 *  @Description: 该节点的字符
	 */
	Char string

	//  ContentList
	/**
	 *  @Author: rym 2022-11-16 09:10:54
	 *  @Description:
	 */
	Content []string `json:"content"`

	//  Code
	/**
	 *  @Author: rym 2022-11-16 09:08:38
	 *  @Description: 改节点的unicode
	 */
	Code rune
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
