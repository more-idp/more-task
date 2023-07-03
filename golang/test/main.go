package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	task "github.com/more-idp/more-task"
)

var aaa = "aa"

func main() {
	log.Println("Xxxxxxx")
	a := task.NewPubSub(&task.ClientConfig{
		RedisServer: "192.168.1.251:6379",
		Name:        "test",
	})
	log.Println(a)
	req := task.Request{
		Topic: "aaaa",
	}
	a.RunTask(&req)

	go func() {
		i := 0
		for {
			r, err := a.Pop()
			if err != nil {
				log.Println("Errir", err)
				continue
			}
			i += 1
			log.Println("xxxxx", i, r)
			time.Sleep(time.Microsecond * 100)
		}

	}()
	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	log.Println("EXIT")
}
