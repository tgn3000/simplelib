// export GOPATH=`git rev-parse --show-toplevel`
package main

import (
    "fmt"
    "../src"
)

func main() {
    simpleClass := simplelib.NewSimpleClass()
    result := simpleClass.Hello()
    fmt.Println(result)

    strings := simplelib.NewStringVector()
    simpleClass.HelloString(strings)

    var i int64
    for i = 0; i < strings.Size(); i++ {
        fmt.Println(strings.Get(int(i)))
    }

    bytes := simplelib.NewByteVector()
    simpleClass.HelloBytes(bytes)

    for i = 0; i < bytes.Size(); i++ {
        fmt.Printf("%c", bytes.Get(int(i)))
    }

    c := simplelib.NewSCCGo()
    simpleClass.SetCallBack(c)
    simplelib.DeleteSCCGo(c)

    fmt.Println("")
}