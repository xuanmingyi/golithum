package main

import (
    "flag"
    "os"
    "fmt"
    "github.com/xuanmingyi/golithum/server"
    //"github.com/xuanmingyi/golithum/client"
    "github.com/xuanmingyi/golithum/http"
    "github.com/xuanmingyi/golithum/g"
)

func main(){

    cfg := flag.String("c", "cfg.json", "configuration file")
    version := flag.Bool("v", false, "show version")
    //check := flag.Bool("check", false, "check collector")

    flag.Parse()

    if *version {
        fmt.Println(g.VERSION)
        os.Exit(0)
    }

    g.ParseConfig(*cfg)

    if g.Config().Debug {
        g.InitLog("debug")
    }else{
        g.InitLog("info")
    }

    go http.Start()
    go server.Start()
    //go client.Start()

    select {}
}