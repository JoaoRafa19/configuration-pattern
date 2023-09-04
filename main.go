package main

import (
	"fmt"
)

type OptFunc func (*Opts)


func defaultOpts() Opts {
	return Opts{
		MaxConn : 10,
		Id: "default",
		Tls: false,
	}
}

//Configurations options
type Opts struct {
	MaxConn int
	Id string
	Tls bool
}

type Server struct {
	Opts
}

func withId(id string) OptFunc {

	return func (opts * Opts) {
		opts.Id = id 
	}
}

func withMaxCon(n int) OptFunc {
	return func (opts *Opts){
		opts.MaxConn = n
	}
}

func withTLS(opts *Opts){
	opts.Tls = true
}

func newServer(opts ...OptFunc) *Server {
	o:= defaultOpts()

	for _, fn := range opts {
		fn(&o)
	}

	return &Server{
		Opts:	o,
	}
		
}


func main() {
	s:= newServer( withMaxCon(99), withId("server 01"))
	fmt.Printf("Server : %+v\n", s)
}
