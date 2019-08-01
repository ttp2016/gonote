![banner](images/banner.png)<br>
[![License](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://github.com/xmge/gonote/blob/master/LICENSE)
[![blog](https://img.shields.io/badge/Author-Blog-7AD6FD.svg)](https://github.com/xmge)
[![Open Source Love](https://badges.frapsoft.com/os/v2/open-source.png?v=103)](https://github.com/xmge)
[![GitHub stars](https://img.shields.io/github/stars/xmge/gonote.svg?label=Stars)](https://github.com/xmge/gonote) 
[![GitHub forks](https://img.shields.io/github/forks/xmge/gonote.svg?label=Fork)](https://github.com/xmge/gonote)
[![All Contributors](https://img.shields.io/badge/all_contributors-5-orange.svg?style=flat-square)](#contributors) 

> 莫听穿林打叶声，何妨吟啸且徐行。竹杖芒鞋轻胜马，谁怕？ 一蓑烟雨任平生。 料峭春风吹酒醒，微冷，山头斜照却相迎。回首向来萧瑟处，归去，也无风雨也无晴。
---

## 1.每日一题

1、写出下面代码输出内容。

```
package main

import (
	"fmt"
)

func main() {
	defer_call()
}

func defer_call() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}
```

## 2.专题讲解
> 在这里我们了解go，学习go，调侃go ~~

 未完待续

## 3.开源项目
> 在这里我们了解go，学习go，调侃go ~~

 未完待续

## 4.面试题
> 以考代练，预祝各位考生取得优异成绩！fighting

- [1.卷一](go面试题/卷1.md)<br>
- [2.卷二](go面试题/卷2.md)<br>
- [3.卷三](go面试题/卷3.md)<br>
- [4.卷四](go面试题/卷4.md)<br>



