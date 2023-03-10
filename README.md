# triesearch
基于go+内存的轻量级搜索引擎

使用方法：

下载安装  
go get -u github.com/BabyRunPlus/triesearch

第一步：初始化应用  
s := triesearch.NewTrie()

第二步：设置索引深度 深度越小，搜索速度越快  
s.SetDepth(5)

第三步：创建索引与搜索内容  
&emsp;&emsp;list := []string{  
&emsp;&emsp;&emsp;&emsp;"碳中和",  
&emsp;&emsp;&emsp;&emsp;"中国",  
&emsp;&emsp;&emsp;&emsp;"中华人民共和国"  
&emsp;&emsp;}  

&emsp;&emsp;#自动创建全索引搜索引擎  
&emsp;&emsp;s.AutoAddFull(list)   

&emsp;&emsp;#自动创建简单索引搜索引擎  
&emsp;&emsp;s.AutoAdd(list)

&emsp;&emsp;#自定义关键词创建全索引搜索引擎  
&emsp;&emsp;for i := 0; i < len(list); i++ {  
&emsp;&emsp;&emsp;&emsp;s.AddFull("自定义关键词", list[i])  
&emsp;&emsp;}

&emsp;&emsp;#自定义关键词创建简单索引搜索引擎  
&emsp;&emsp;for i := 0; i < len(list); i++ {  
&emsp;&emsp;&emsp;&emsp;s.Add("自定义关键词", list[i])  
&emsp;&emsp;}  

&emsp;&emsp;#数据结构创建结束后建议执行s.GC()用来清除临时数据

搜索：  
r, c := s.Find("碳中和")  
&emsp;&emsp;#r 是否为完全匹配结果 关键词超出深度时返回false  
&emsp;&emsp;#s 搜索出来的内容列表  

全索引与简单索引的区别  
#搜索关键词“中”   
&emsp;&emsp;全索引结果：碳中和、中国、中华人民共和国  
&emsp;&emsp;简单索引结果：中国、中华人民共和国  