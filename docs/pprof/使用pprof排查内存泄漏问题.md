# 使用pprof排查内存泄漏问题.md

## 一、发现问题

首先通过监控工具查看到某个项目的机器内存在部署之后总是不断上涨，但是用户量并不多，很明显是内存泄漏的问题。

![监控](http://i.xmge.top/img/1594483747-22608778-0cce52fadabd6180.png)

## 二、如何解决

项目中引入了以下代码，自然可以通过 pprof 工具进行分析。

```go
package main

import (
	"net/http"
	_"net/http/pprof"
)

func init() {
	go func() {
		http.ListenAndServe(":8000",nil)
	}()
}
```

## 三、解决步骤

解决问题尝试了很多种方案，最后方案如下：

### 1、通过 pprof 工具获取内存相差较大的两个时间点 heap 数据。

> curl localhost:8000/debug/pprof/heap > heap.base

等待一段时间，通过 htop 可以查看到内存又涨了很多，然后再采集内存情况

> curl localhost:8000/debug/pprof/heap > heap.current


### 2、通过 go tool pprof 工具比较两个内存的情况，找到是什么对象多创建了

> go tool pprof -http=:8080 -base heap.base heap.current

选择当前分配的对象（insue_objects）：

![option](http://i.xmge.top/img/1594483805-22608778-ebb76935450dee24.png)


得到如图所示：

![pprof-heap-base](http://i.xmge.top/img/1594483785-22608778-5d02a9aa64fd309f.png )



图中可以看出 withdraw_record.GetByUserId.FindAndCount()  在这段时间创建了 624110 个对象，于是怀疑是这里出了问题，于是去查看代码：

```go
  count, err = statement.Where("user_id = ? ", userId).FindAndCount(&records)
```

发现 statement 这个数据库 session 没有关闭，怀疑是因为没有关闭造成的问题

### 3、测试服问题复现

既然怀疑是这里的问题，然后就写了个 for 循环，不断地请求嫌疑接口，通过htop 发现，内存果然`蹭蹭蹭`网上涨，问题复现成功

### 4、将嫌疑代码修复了

修复就很简单了，加了一个 `defer statement.Close()`

### 5、部署修复后的代码到测试服务器验证

代码修复后，部署到测试服上，再用 for 循环去测，发现内存不再上涨，到此应该算是问题解决

### 6、查找项目中有没有类似代码并加以改正

## 四、注意点

### 1、不要在同一台机器一边跑项目，一边压测，否则两个程序都跑不满。
