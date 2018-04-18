# crawler_concurrent    

### 创建意图    

该项目用于实现并发版爬虫，对于网站进行有些数据的提取、存储、查找、以及页面展示。    

爬取的网站：http://www.zhenai.com/zhenghun    

### 实现路程    

先实现单任务版本、随后实现并发版本、最终实现分布式版本。 

**单任务版本**    

* 获取并打印第一页用户的信息信息。  
  需要安装用于网页字符转换的库：    
```
gopm get -g -v golang.org/x/text    
gopm 
```



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
