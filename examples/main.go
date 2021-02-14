package main

import (
    "fmt"
    "github.com/zacg/simplelib/src"
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
    fmt.Println("")

}