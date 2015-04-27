package main

import (
    "fmt"
    "io/ioutil"
)

func main() {
    dat, err := ioutil.ReadFile("E:/go/src/github.com/danhardman/officr/sys/bus/w1/devices/10-000802824e58/w1_slave")

    if err != nil {
        panic(err)
    }

    fmt.Print(string(dat))
}
