package main

import (
	"bufio"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	log.Println("Starting chat service...")
	// enter code for invite people to chat room here

	log.Println("Enter code: ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading input:", err)
	}
	log.Printf("You entered: %s", input)

	log.Println("Chat service started successfully")
}
