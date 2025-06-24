package main

import "fmt"

func main() {
	currentMessageLength := 70
	maxMessageLength := 40 

	fmt.Println("Trying to send the message of length:", currentMessageLength)

	if currentMessageLength <= maxMessageLength{
		fmt.Println("Message Sent")
	}else{
		fmt.Println("Message not sent. Message length too long")
	}

	// shorter way to do this by limiting the scope of the variable used

	if length := 40; length > 1{
		fmt.Println("Length is valid")
	} 
}