# sync.Atomic 源码阅读

## 1.Demo

```go
package main

import (
	"fmt"
	"time"
	"sync"
	"sync/atomic"
)

func main() {
	test1()
	test2()
}

// count++  并发不安全
func test1()  {
	var wg sync.WaitGroup
	count := 0
	t := time.Now()
	for i := 0 ; i < 100000 ; i++ {
		wg.Add(1)
		go func(i int) {
			count++   //count不是并发安全的
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Printf("test1 花费时间：%d, count的值为：%d \n",time.Now().Sub(t),count)
}

// atomic.AddInt64(&count,1)  //原子操作
func test2()  {
	var wg sync.WaitGroup
	count := int64(0)
	t := time.Now()
	for i := 0 ; i < 100000 ; i++ {
		wg.Add(1)
		go func(i int) {
			atomic.AddInt64(&count,1)  //原子操作
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Printf("test2 花费时间：%d, count的值为：%d \n",time.Now().Sub(t),count)
}
```

结果：

```
28.922ms
count====> 93095
exit
27.9249ms
count====> 100000
exit
```

## 2.介绍

我们调用sync/atomic中的几个函数可以对几种简单的类型进行原子操作。这些类型包括int32,int64,uint32,uint64,uintptr,unsafe.Pointer,共6个。这些函数的原子操作共有5种：增减，存储，载入，交换，比较并交换。

sync/atomic 解决的典型问题就是 i++和CAS（Compare-and-Swap）的线程安全问题，它的实现原理大致是向CPU发送对某一个块内存的LOCK信号，然后就将此内存块加锁，从而保证了内存块操作的原子性。

与Mutex相比，它的优势主要有以下几点：

1. 更高效，因为atomic是直接作用与内存的锁，所以更底层，更高效。在Demo中的用时也可以看出。
2. 更简洁，atomic避免了加锁解锁的过程，一行代码就可以完成这个操作，使代码更简洁，更具有可读性。

## 3.使用场景

sync/atomic 可以在并发场景下对变量进行非侵入式的操作。可以保证并发安全，虽然使用 `sync.Mutex` 可以实现，但是使用`sync/atomic`不仅是轻量级的，而且代码也更加简洁。

## 4.源码

```
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package atomic provides low-level atomic memory primitives
// useful for implementing synchronization algorithms.
//
// These functions require great care to be used correctly.
// Except for special, low-level applications, synchronization is better
// done with channels or the facilities of the sync package.
// Share memory by communicating;
// don't communicate by sharing memory.
//
// The swap operation, implemented by the SwapT functions, is the atomic
// equivalent of:
//
//	old = *addr
//	*addr = new
//	return old
//
// The compare-and-swap operation, implemented by the CompareAndSwapT
// functions, is the atomic equivalent of:
//
//	if *addr == old {
//		*addr = new
//		return true
//	}
//	return false
//
// The add operation, implemented by the AddT functions, is the atomic
// equivalent of:
//
//	*addr += delta
//	return *addr
//
// The load and store operations, implemented by the LoadT and StoreT
// functions, are the atomic equivalents of "return *addr" and
// "*addr = val".
//
package atomic

import (
	"unsafe"
)

// BUG(rsc): On x86-32, the 64-bit functions use instructions unavailable before the Pentium MMX.
//
// On non-Linux ARM, the 64-bit functions use instructions unavailable before the ARMv6k core.
//
// On both ARM and x86-32, it is the caller's responsibility to arrange for 64-bit
// alignment of 64-bit words accessed atomically. The first word in a
// variable or in an allocated struct, array, or slice can be relied upon to be
// 64-bit aligned.

// SwapInt32 atomically stores new into *addr and returns the previous *addr value.
func SwapInt32(addr *int32, new int32) (old int32)

// SwapInt64 atomically stores new into *addr and returns the previous *addr value.
func SwapInt64(addr *int64, new int64) (old int64)

// SwapUint32 atomically stores new into *addr and returns the previous *addr value.
func SwapUint32(addr *uint32, new uint32) (old uint32)

// SwapUint64 atomically stores new into *addr and returns the previous *addr value.
func SwapUint64(addr *uint64, new uint64) (old uint64)

...
```

## 5.源码详解

源码中并没有go语言版本的实现，在此介绍一下 `sync/atomic` 中的方法（只介绍一种类型，其他类型都是一样的）

1.func AddInt64(addr *int64, delta int64) (new int64)

将addr增加delta（如何要减少，直接将delta为负数即可）

```go
func main() {
	var addr int64
	addr = atomic.AddInt64(&addr,1)
	fmt.Println(addr)
}

//1
```

2.func StoreUint64(addr *uint64, val uint64)

为addr赋值为val

```go
func main() {
	var addr int64
	atomic.StoreInt64(&addr,10)
	fmt.Println(addr)
}
```


3.func LoadInt64(addr *int64) (val int64)

加载addr的值

```go
func main() {
	var addr int64 = 10
	addr = atomic.LoadInt64(&addr)
	fmt.Println(addr)
}

// 10
```

4.func SwapInt64(addr *int64, new int64) (old int64) 交换

将addr与new交换，并返回之前addr的值

```go
func main() {
	var i int64  = 10
	old := atomic.SwapInt64(&i,20)
	fmt.Println(old)
	fmt.Println(i)
}

// 10
// 20
```

5.func CompareAndSwapInt64(addr *int64, old, new int64) (swapped bool)

将addr与old进行比较，想等则返回true,不相等返回false,并将addr赋值为new。

```go
func main() {
	var addr int64 = 2
	compare := atomic.CompareAndSwapInt64(&addr,2,1)
	fmt.Println(compare)
	fmt.Println(addr)
}

// true
// 1
```