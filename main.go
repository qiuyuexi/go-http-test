package main

import (
	"flag"
	"fmt"
	"go-test/pkg/server"
	"os"
)

var (
	p int
	h bool
	d bool
)

func main() {
	flag.Parse()
	if h {
		flag.Usage()
	} else {
		server.RegisterSignal()
		server.LogPid()
		server.Start(p)
	}
}
func init() {
	flag.BoolVar(&h, "h", false, "帮助")
	flag.BoolVar(&d, "d", false, "后台运行")
	flag.IntVar(&p, "p", 9090, "端口")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `usage [-hd] [-p port] 
options:
`)
		flag.PrintDefaults()
	}
}
