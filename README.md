# tiresearch
基于go+内存的轻量级搜索引擎

使用方法：

下载安装 go get -u github.com/BabyRunPlus/tiresearch

第一步：初始化应用
s := tiresearch.NewTire()

第二步：设置索引深度 深度越小，搜索速度越快
s.SetDepth(5)

第三步：创建索引与搜索内容
  list := []string{
		"碳中和",
		"中信证券",
		"中国平安",
		"中信证券今天大涨",
	}
  #自动创建全索引搜索引擎
  s.AutoAddFull(list)
  
  #自动创建简单索引搜索引擎
  s.AutoAdd(list)
  
  #自定义关键词创建全索引搜索引擎
  for i := 0; i < len(list); i++ {
		s.AddFull("自定义关键词", list[i])
	}
  
  #自定义关键词创建简单索引搜索引擎
  for i := 0; i < len(list); i++ {
		s.Add("自定义关键词", list[i])
	}

搜索：
r, c := s.Find("大")
  #r 是否为完全匹配结果  例如：搜索关键词“碳中和”
  #s 搜索出来的内容列表
