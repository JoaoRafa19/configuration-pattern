package main

import (
	"fmt"
)

//Configurations options
type Opts struct {
	MaxConn int
	Id string
	Tls bool
}

type Server struct {
	Opts
}

func newServer(opts Opts) *Server {
return &Server{
	Opts:	opts,
}
		
}


func main() {
	s:= newServer(Opts{})
	fmt.Println("Server : %+v", s)
}
