package main

import (
	"fmt"
	"time"
)

type Token struct{
	data string
	recipient int
	ttl int
}

func SendToken(tc chan Token, recipient int) {
	tokenState := <-tc
	fmt.Println(tokenState.ttl)
	if tokenState.ttl != 0 && tokenState.recipient != recipient {
		tokenState.ttl -= 1
		tc <- tokenState
		time.Sleep(time.Second * 1)
	} else if tokenState.recipient == recipient {
		fmt.Println(tokenState.data)
	} else {
		fmt.Println("lol, kek, sort of error")
	}
}

func main() {
	var n int
	var token Token
	fmt.Scanf("%s\n", &token.data)
	fmt.Scanf("%d\n", &token.recipient)
	fmt.Scanf("%d\n", &token.ttl)
	fmt.Scanf("%d\n", &n)

	// Запуск потоков
	tokenChan := make(chan Token)
	for i := 0; i < n; i++ {
		go SendToken(tokenChan, i)
	}

	// Запись в первый поток
	tokenChan <- token

	var input string
	fmt.Scanln(&input)
}