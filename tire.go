package tiresearch

// Tire
/**
 *  @Author: rym 2022-11-16 09:18:42
 *  @Description:Tire Map 数据结构对象
 */
type Tire struct {
	//  Root
	/**
	 *  @Author: rym 2022-11-16 09:14:41
	 *  @Description: 根节点
	 */
	Root *Node

	//  depth
	/**
	 *  @Author: rym 2022-11-16 09:15:01
	 *  @Description: 树的深度 深度越小 搜索越快
	 */
	depth int8
}

// SetDepth
/**
 *  @Author: rym 2022-11-16 09:41:27
 *  @Description: 设置树的最大深度 深度越大 查找速度越慢
 *  @receiver t
 *  @param n
 */
func (t *Tire) SetDepth(n int8) {
	t.depth = n
}

// AddFull
/**
 *  @Author: rym 2022-11-16 17:32:32
 *  @Description: 根据关键词添加节点，并下挂相关内容 对外方法
 *  @Description: 创建详细的树节点 例如：碳中和  通过 碳、碳中、碳中和、中、中和、和 都可以查找到对应的内容列表
 *  @Description: 此方式创建的树，会更多消耗内存
 *  @receiver t
 *  @param keyword
 *  @param content
 */
func (t *Tire) AddFull(keyword, content string) {
	//  根据深度裁剪关键词长度
	cutNum, strLen := calcCutNum(keyword, int(t.depth))

	//  创建树
	for i := 0; i < strLen; i++ {
		t.insert(substr(keyword, i, cutNum), content)
	}
}

// Add
/**
 *  @Author: rym 2022-11-16 17:39:14
 *  @Description: 根据关键词添加节点，并下挂相关内容 对外方法
 *  @Description: 创建普通的树节点 例如： 碳中和 通过 碳、碳中、碳中和 可以查找到对应的内容列表
 *  @receiver t
 *  @param keyword
 *  @param content
 */
func (t *Tire) Add(keyword, content string) {
	//  根据深度裁剪关键词长度
	cutNum, _ := calcCutNum(keyword, int(t.depth))

	//  创建树
	t.insert(substr(keyword, 0, cutNum), content)
}

// Find
/**
 *  @Author: rym 2022-11-16 09:44:22
 *  @Description: 根据关键词，从树中找相关内容列表
 *  @receiver t
 *  @param keyword
 *  @return bool
 *  @return []string
 */
func (t *Tire) Find(keyword string) (bool, []string) {
	cutNum, wordLen := calcCutNum(keyword, int(t.depth))
	keyword = substr(keyword, 0, cutNum)

	node := t.Root
	for _, code := range keyword {
		//  获取对应子节点
		value, ok := node.Children[code]
		if !ok {
			// 不存在则直接返回
			return false, []string{}
		}

		// 否则继续往后遍历
		node = value
	}

	// 找到对应单词
	if wordLen > int(t.depth) {
		//按照深度匹配前缀
		return false, node.Content
	} else {
		//完全匹配
		return true, node.Content
	}

}

// insert
/**
 *  @Author: rym 2022-11-16 09:38:36
 *  @Description: 根据关键词添加节点，并下挂相关内容
 *  @receiver t
 *  @param keyword
 *  @param content
 */
func (t *Tire) insert(keyword, content string) {
	node := t.Root                 // 获取根节点
	for _, code := range keyword { // 以 Unicode 字符遍历该单词
		value, ok := node.Children[code] // 获取 code 编码对应子节点
		if !ok {
			// 不存在则初始化该节点
			value = newTrieNode(string(code))
			// 然后将其添加到子节点字典
			node.Children[code] = value
		}
		// 当前节点指针指向当前子节点
		node = value
		if ok := listSearchString(content, node.Content); ok > -1 {
			node.Content = append(node.Content, content)
		}
	}
}
