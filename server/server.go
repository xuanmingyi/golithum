package server

import (
	"github.com/xuanmingyi/golithum/g"

	"fmt"
	"log"
	"net/http"
)

func init(){

}

func Start(){
	if !g.Config().Server.Enabled {
		return
	}
	addr := fmt.Sprintf("0.0.0.0:%d", g.Config().Server.Port)

	s := &http.Server{
		Addr: addr,
		MaxHeaderBytes: 1<<30,
	}

	log.Println("server listening at ", addr)
	log.Fatalln(s.ListenAndServe())
}