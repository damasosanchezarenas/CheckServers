package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	init := time.Now()

	channel := make(chan string)

	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, server := range servers {
		go checkServ(server, channel)
	}

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-channel) //Main wait to channels
	}

	countTime := time.Since(init)
	fmt.Printf("Tiempo de Ejecucion: %s/n", countTime)
}

func checkServ(serv string, channel chan string) {
	_, err := http.Get(serv)
	if err != nil {
		channel <- serv + " is not available :( "
	} else {
		channel <- serv + " is available :) "
	}
}
