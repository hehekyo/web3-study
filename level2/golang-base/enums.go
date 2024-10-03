package main

import "fmt"

type ServerState int

// 使用 iota 关键字定义了几个常量，这是 Go 中生成一系列相关常数的常用方法。
//iota 在每个 const 块中从 0 开始，每次使用时自增。
const (
    StateIdle ServerState = iota
    StateConnected
    StateError
    StateRetrying
)

var stateName = map[ServerState]string{
    StateIdle:      "idle",
    StateConnected: "connected",
    StateError:     "error",
    StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
    return stateName[ss]
}

func main() {
    ns := transition(StateIdle)
    fmt.Println(ns)

    ns2 := transition(ns)
    fmt.Println(ns2)
}

func transition(s ServerState) ServerState {
    switch s {
    case StateIdle:
        return StateConnected
    case StateConnected, StateRetrying:

        return StateIdle
    case StateError:
        return StateError
    default:
        panic(fmt.Errorf("unknown state: %s", s))
    }
}