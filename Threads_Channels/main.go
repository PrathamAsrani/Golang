package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Message struct {
	friends []string
	chats   []string
}

func main() {
	now := time.Now()
	fmt.Println("Routines & Channels")

	id := getUserByName("Pratham")
	fmt.Println("User ID:", id)

	wg := &sync.WaitGroup{}
	ch := make(chan *Message) 

	wg.Add(2) // two goroutines

	go getUserChats(id, ch, wg)
	go getUserFriends(id, ch, wg)

	// using anonymous function in thread so that each thread can communicate for a single resource, # to optimized space complexity
	go func(){
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		log.Println(msg)
	}

	fmt.Println("Execution Time:", time.Since(now))
}

func getUserByName(name string) string {
	time.Sleep(time.Second * 1)
	return "user123" // Return a dummy user ID for demonstration
}

func getUserFriends(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 1)
	
	ch <- &Message{
		friends: []string{
			"Pratham",
			"Niharika",
			"Vedant",
			"Smit",
			"Shrayash",
			"Tiago",
		},
	}
	wg.Done()
}

func getUserChats(id string, ch chan<- *Message, wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)
	
	ch <- &Message{
		chats: []string{
			"Hey",
			"Hello",
			"Hi",
			"Good Morning",
			"Good Afternoon",
			"Good Evening",
			"Good Night",
		},
	}
	wg.Done()
}
