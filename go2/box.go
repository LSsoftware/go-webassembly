package main

import (
    "syscall/js"
    "strconv"
)

func add(this js.Value, i []js.Value) interface{}  {
    i11, _ := strconv.Atoi(i[0].String())
    i22, _ := strconv.Atoi(i[1].String())
    println("first = %s, second = %s", i[0].String(), i[1].String());
    return js.ValueOf(i11 + i22);
}

func subtract(this js.Value, i []js.Value) interface{}  {
    i11, _ := strconv.Atoi(i[0].String())
    i22, _ := strconv.Atoi(i[1].String())
    return js.ValueOf(i11 - i22)
}

func multiplication(this js.Value, i []js.Value) interface{}  {
    i11, _ := strconv.Atoi(i[0].String())
    i22, _ := strconv.Atoi(i[1].String())
    return js.ValueOf(i11 * i22)
}

func division(this js.Value, i []js.Value) interface{}  {
    i11, _ := strconv.Atoi(i[0].String())
    i22, _ := strconv.Atoi(i[1].String())
    return js.ValueOf(i11 / i22)
}

func registerCallbacks() {
    // 转换为js的Global的add方法
    js.Global().Set("add", js.FuncOf(add))
    js.Global().Set("subtract", js.FuncOf(subtract))
    js.Global().Set("multiplication", js.FuncOf(multiplication))
    js.Global().Set("division", js.FuncOf(division))
}

func main() {
    c := make(chan struct{}, 0)
    println("WASM Go Initialized")
    // 注册方法
    registerCallbacks()
    <-c
}