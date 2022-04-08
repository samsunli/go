package main
import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan string)
	
	go ReceiveMessage(channel)
	go SendMessage(channel)

	
	fmt.Scanln();
}

func SendMessage(channel chan string) {
	for {
		channel <- "sending message @" + 
		time.Now().String()
		time.Sleep(5 * time.Second)
	}
}

func ReceiveMessage(channel chan string) {
	for {
		message := <-channel
		fmt.Println(message)
	}
}