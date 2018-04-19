# crawler_concurrent    

### 创建意图    

该项目用于实现并发版爬虫，对于网站进行有些数据的提取、存储、查找、以及页面展示。    

爬取的网站：http://www.zhenai.com/zhenghun    

### 实现路程    

先实现单任务版本、随后实现并发版本、最终实现分布式版本。 

**单任务版本**    
以下表述的是实现的过程：

* 获取并打印第一页网页的信息。  
  需要安装用于网页字符转换的库，具体操作参考下面的知识点部分：    
```
get -g -v golang.org/x/text    
get -g -v golang.org/x/net
```
* 通过正则表达式获取到第一页网页信息中，所包含的城市名字和URL。  
  对于网页中信息提取的方法包括：    
  1. 使用css选择器。    
  2. 使用xpath。    
  3. 使用正则表达式。    
  
 
* 构思爬虫的算法    
  根据爬取页面的手动点击和总结，可以将对于每个业务获取数据的动作称为解析器（Parser）。    
  对于不同的网站需要不同的解析器，而当前需要爬取的网站其中包含三种解析器分别为：    
  1. 城市列表解析器。    
  2. 城市解析器。    
  3. 用户解析器。    
  
  对于解析器的主要承担的工作有，输入输出。    
  其中输入为：utf-8编码的文本。输出为：Request{需要操作的URL, 对应Parser方法}列表和获取的有用信息item，对于不同的URL需要不同的解析器进行处理。    

* 构思爬虫整体架构
  对于整个程序，重要的是需要一个驱动程序（engine），负责将整个程序进行运转。驱动程序起初需要一个种子（seed）即爬取网闸的请求（request），将其添加到任务队列。 随后驱动程序依次将任务队列中的request请求，发送给收集器（fetcher）, fercher收集器负责从Internet获取到相应请求url的网页信息(text)，将utf-8的网页信息返回给engine。其次engine负责将text提交给解析器进行解析，解析器将解析的结果（ParseResult）返回给engine。最后engine将返回的结果进行拆分，其中需要再次请求Requests的信息放入到任务队列，结果数据（item）进行打印或存储。其整体流程图如下所示：    
  ![整体架构images0](https://github.com/Clodfisher/crawler_concurrent/raw/master/readmeimages/images0.jpg)     

* 对城市列表解析进行测试      
  测试方法：采用表格驱动的方法


### 知识点    

**下载golang.org中的包**
gopm 代替go 下载第三方依赖包    
在国内采用go get有时会下载不到一些网站如golang.org的依赖包。    
可以采用gopm从golang.org一些镜像网站上下载。    
1. 安装gopm    
```
go get -u -v github.com/gpmgo/gopm
执行该命令的用户需要设置GOPATH参数
```
2. gopm get 如果不携带-g采用，会把依赖包下载.vendor目录下面。    
```
采用-g 参数，可以把依赖包下载到GOPATH目录中 
gopm get -g golang.org/x/net   
```

**正则表达式的基本使用**    

主要包括正则表达式的匹配，以及将匹配串中需要的子串提取出来。

或参考：[go get golang.org/x 包失败解决方法](https://blog.csdn.net/alexwoo0501/article/details/73409917)
