## 【标准库系列】runtime.Caller()

## 1.runtime.Caller() 详解

Caller 方法反应的是堆栈信息中某堆栈帧所在文件的绝对路径和语句所在文件的行数。而 skip 表示的是从下往上数第几个堆栈帧。如果要打印全部堆栈信息可以直接使用 `debug.PrintStack()` 来实现。

源码：

```go
// Caller reports file and line number information about function invocations on
// the calling goroutine's stack. The argument skip is the number of stack frames
// to ascend, with 0 identifying the caller of Caller.  (For historical reasons the
// meaning of skip differs between Caller and Callers.) The return values report the
// program counter, file name, and line number within the file of the corresponding
// call. The boolean ok is false if it was not possible to recover the information.
func Caller(skip int) (pc uintptr, file string, line int, ok bool) {
	// Make room for three PCs: the one we were asked for,
	// what it called, so that CallersFrames can see if it "called"
	// sigpanic, and possibly a PC for skipPleaseUseCallersFrames.
	var rpc [3]uintptr
	if callers(1+skip-1, rpc[:]) < 2 {
		return
	}
	var stackExpander stackExpander
	callers := stackExpander.init(rpc[:])
	// We asked for one extra, so skip that one. If this is sigpanic,
	// stepping over this frame will set up state in Frames so the
	// next frame is correct.
	callers, _, ok = stackExpander.next(callers, true)
	if !ok {
		return
	}
	_, frame, _ := stackExpander.next(callers, true)
	pc = frame.PC
	file = frame.File
	line = frame.Line
	return
}
```

demo：

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	A()                                 // 9
}

func A()  {
	B()                                 // 13
}

func B()  {
	C()                                 // 17
}

func C()  {
	D()                                 // 21
}

func D()  {
	for i:=0;i<=4;i++ {
		_,f,line,_ := runtime.Caller(i) // 27
		fmt.Println(f,"----",line)
	}

    debug.PrintStack()
}


// /home/xmge/ws/gonote/src/github.com/xmge/gonote/go标准库/runtime/main.go ---- 27
// /home/xmge/ws/gonote/src/github.com/xmge/gonote/go标准库/runtime/main.go ---- 21
// /home/xmge/ws/gonote/src/github.com/xmge/gonote/go标准库/runtime/main.go ---- 17
// /home/xmge/ws/gonote/src/github.com/xmge/gonote/go标准库/runtime/main.go ---- 13
// /home/xmge/ws/gonote/src/github.com/xmge/gonote/go标准库/runtime/main.go ---- 9

//	runtime/debug.PrintStack()
//	/home/xmge/go/go1.11/src/runtime/debug/stack.go:16 +0x22
//	main.D()
//	/home/xmge/ws/gonote/src/github.com/xmge/gonote/go标准库/runtime/main.go:31 +0x11f
//	main.C()
//	/home/xmge/ws/gonote/src/github.com/xmge/gonote/go标准库/runtime/main.go:22 +0x20
//	main.B()
//	/home/xmge/ws/gonote/src/github.com/xmge/gonote/go标准库/runtime/main.go:18 +0x20
//	main.A()
//	/home/xmge/ws/gonote/src/github.com/xmge/gonote/go标准库/runtime/main.go:14 +0x20
//	main.main()
//	/home/xmge/ws/gonote/src/github.com/xmge/gonote/go标准库/runtime/main.go:10 +0x20
```

## 2.runtime.Caller() 使用场景

1.写开源框架时，想引入自己的静态文件，但是又不知道自己的项目放在哪个路径下，可以通过返回的 file 参数来找到文件位置，然后使用它。
例如 [seelog](https://github.com/xmge/seelog/blob/master/server.go) 中引入 index.html 的方式

```go
// 输出page
func showPage(writer http.ResponseWriter, page string, data interface{}) {
	_, currentfile, _, _ := runtime.Caller(0) // 忽略错误
	filename := path.Join(path.Dir(currentfile), page)
	t, err := template.ParseFiles(filename)
	if err != nil {
		log.Println(err)
	}
	t.Execute(writer, data)
}
```