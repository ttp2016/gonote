package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

func main() {
	A()
}

func A()  {
	B()
}

func B()  {
	C()
}

func C()  {
	D()
}

func D()  {
	for i:=0;i<=4;i++ {
		_,f,line,_ := runtime.Caller(i)
		fmt.Println(f,"----",line)
	}

	debug.PrintStack()
}

// /home/xmge/ws/gonote/src/github.com/xmge/gonote/go标准库/runtime/main.go ---- 27
// /home/xmge/ws/gonote/src/github.com/xmge/gonote/go标准库/runtime/main.go ---- 21
// /home/xmge/ws/gonote/src/github.com/xmge/gonote/go标准库/runtime/main.go ---- 17
// /home/xmge/ws/gonote/src/github.com/xmge/gonote/go标准库/runtime/main.go ---- 13
// /home/xmge/ws/gonote/src/github.com/xmge/gonote/go标准库/runtime/main.go ---- 9
