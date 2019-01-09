package http

import (
	"github.com/xuanmingyi/golithum/g"

	"fmt"
	"log"
	"net/http"
)

func init(){
}

func Start(){
	if !g.Config().Http.Enabled {
		return
	}
	addr := fmt.Sprintf("0.0.0.0:%d", g.Config().Http.Port)

	s := &http.Server{
		Addr: addr,
		MaxHeaderBytes: 1<<30,
	}

	log.Println("http listening at ", addr)
	log.Fatalln(s.ListenAndServe())
}