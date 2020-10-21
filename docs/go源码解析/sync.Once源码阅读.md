# sync.Once 源码阅读

## １.Demo

```
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var once sync.Once

	for i:=0;i<=10;i++{
		go once.Do(func() {
			fmt.Println("hello world")
		})
	}

	time.Sleep(time.Second * 2)
}
```

## 2.介绍

sync.Once是sync包中的一个对象，它只有一个方法Do,这个方法很特殊，在程序运行过程中，无论被多少次调用，只会执行一次，就与结构体的名称一样，once（一次）。那它是如何做的呢？

## 3.使用场景

1. 当程序运行过程中，在会被多次调用的地方却只想执行一次某代码块。就可以全局声明一个once，然后用once.Do()来之行此代码块。

## ４.源码

```
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

import (
	"sync/atomic"
)

// Once is an object that will perform exactly one action.
type Once struct {
	m    Mutex
	done uint32
}

// Do calls the function f if and only if Do is being called for the
// first time for this instance of Once. In other words, given
// 	var once Once
// if once.Do(f) is called multiple times, only the first call will invoke f,
// even if f has a different value in each invocation. A new instance of
// Once is required for each function to execute.
//
// Do is intended for initialization that must be run exactly once. Since f
// is niladic, it may be necessary to use a function literal to capture the
// arguments to a function to be invoked by Do:
// 	config.once.Do(func() { config.init(filename) })
//
// Because no call to Do returns until the one call to f returns, if f causes
// Do to be called, it will deadlock.
//
// If f panics, Do considers it to have returned; future calls of Do return
// without calling f.
//
func (o *Once) Do(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	// Slow-path.
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		f()
	}
}

```

## 5.源码解析

可以看到once结构体中，有两个字段，ｍ是了保证并发安全性的，done是标志是否已经执行过此方法，如果done是１则表示执行过，０表示未执行。

Do方法中，首先通过atomic.LoadUint32(&o.done)，来取得done的值，看是否为１，如果为１就表示已经执行过了，直接返回，未执行则继续执行。

代码很简单，就不啰嗦了，值得注意的是　`defer atomic.StoreUint32(&o.done, 1)`很精髓，为了防止f()方法中panic，无法为done赋值，作者特地使用defer。值得学习。
