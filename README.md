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

* 对于城市列表中的每个城市URL进行解析，获取到用户名和用户URL        

* 用户信息进行解析    

* 分析单任务版本爬虫存在的性能问题    
  对于单任务版本，其对于互谅网上获取网页数据耗时比较到，而此时整个程序需要再次等待，当完全获取到数据后才能继续执行，所以可以将此处进行并发处理，异步化。   

**并发版本**    
以下表述的是实现的过程：    

* 将fetcher和parser合并成一个work动作    
针对上述存在的性能问题，可以将从网页获取数据代码异步化。而对于解析器（parser）,是从本地获取数据，不存在性能问题，但是没获取到一个网页内容都要进行解析，因此可以将获取网页内容（fetcher)和parser合并成一个work动作函数，进行真正网页与数据的处理。可将上述流程图，修改成如下图所示：    
![将fetcher和parser合并成一个work动作images1](https://github.com/Clodfisher/crawler_concurrent/raw/master/readmeimages/images1.jpg)         

* 构建一个调度程序（Scheduler）
  Schedule用于分发request到多个worker中，从而实现多对多的关系。其架构图如下所示：
![构建一个调度程序](https://github.com/Clodfisher/crawler_concurrent/raw/master/readmeimages/images2.jpg)   
注： 其中每个线表示一个chanal;每个方框表示一个goroute；Request表示多个请求   

* Schedule中实现所有的Worker共用一个输入       
![所有的Worker共用一个输入](https://github.com/Clodfisher/crawler_concurrent/raw/master/readmeimages/images3.jpg)   
  这种情况下会出现一个问题，就是锁死chan。    

* 实现Scheduler为每个Request创建一个goroutine      
  主要功能是为每个Request创建一个goroutine，每个goroutine只做一件事情，往worker同一队列分发request。其缺点是控制力比较小，无法控制goroutine，以及无法控制request给那个worker，对于负载均衡无法实现。实现过程如下图所示：    
![所有的Worker共用一个输入](https://github.com/Clodfisher/crawler_concurrent/raw/master/readmeimages/images4.jpg)      
注：goroutine是与Request数量对等，与worker的数量不对等。    

* Schedule自己实现request的分发
  在Schedule中实现Request队列和Worker队列，Request队列用于存储从Schedule过来的Request，Worker队列用于存储worker，Schedule从Request队列获取相应的Request，交给从Worker队列获取的相应worker，从而实现对request和worker的控制，达到分发的目的。其实现过程如下图所示：    
![所有的Worker共用一个输入](https://github.com/Clodfisher/crawler_concurrent/raw/master/readmeimages/images4.jpg)      
   

   


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

或参考：[go get golang.org/x 包失败解决方法](https://blog.csdn.net/alexwoo0501/article/details/73409917)

**正则表达式的基本使用**    

主要包括正则表达式的匹配，以及将匹配串中需要的子串提取出来。

