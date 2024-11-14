package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

// User struct for storing user credentials
type User struct {
	Username string
	Password string
}

// Message struct for storing messages between users
type Message struct {
	Sender   string
	Receiver string
	Content  string
	Time     time.Time
}

// In-memory storage
var (
	users    = map[string]string{} // map[username]password
	messages = []Message{}
	mutex    sync.Mutex
)

var loggedInUser *User

// Register user
func register() {
	var username, password string
	fmt.Print("Enter a username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter a password: ")
	fmt.Scanln(&password)

	mutex.Lock()
	defer mutex.Unlock()

	if _, exists := users[username]; exists {
		fmt.Println("User already exists")
		return
	}

	users[username] = password
	fmt.Println("Registration successful!")
}

// Login user
func login() {
	var username, password string
	fmt.Print("Enter your username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)

	mutex.Lock()
	defer mutex.Unlock()

	if storedPassword, exists := users[username]; exists && storedPassword == password {
		loggedInUser = &User{Username: username}
		fmt.Println("Login successful!")
		return
	}
	fmt.Println("Invalid username or password")
}

// Send message to another user
func sendMessage() {
	if loggedInUser == nil {
		fmt.Println("Please login first.")
		return
	}

	var receiver, content string
	fmt.Print("Enter receiver's username: ")
	fmt.Scanln(&receiver)
	fmt.Print("Enter message content: ")
	reader := bufio.NewReader(os.Stdin)
	content, _ = reader.ReadString('\n')

	mutex.Lock()
	defer mutex.Unlock()

	msg := Message{
		Sender:   loggedInUser.Username,
		Receiver: receiver,
		Content:  strings.TrimSpace(content),
		Time:     time.Now(),
	}
	messages = append(messages, msg)
	fmt.Println("Message sent!")
}

// View inbox
func viewInbox() {
	if loggedInUser == nil {
		fmt.Println("Please login first.")
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	fmt.Println("\n--- Inbox ---")
	for _, msg := range messages {
		if msg.Receiver == loggedInUser.Username {
			fmt.Printf("[%s] %s: %s\n", msg.Time.Format("2006-01-02 15:04:05"), msg.Sender, msg.Content)
		}
	}
	fmt.Println("--------------")
}

// Real-time message receiver (simulated with continuous check)
func startMessageReceiver() {
	for {
		time.Sleep(5 * time.Second)
		if loggedInUser != nil {
			mutex.Lock()
			for _, msg := range messages {
				if msg.Receiver == loggedInUser.Username {
					fmt.Printf("\n[New message] %s: %s\n", msg.Sender, msg.Content)
				}
			}
			mutex.Unlock()
		}
	}
}

// Menu and main function
func main() {
	go startMessageReceiver()

	for {
		fmt.Println("\n--- Text Messaging CLI ---")
		fmt.Println("1. Register")
		fmt.Println("2. Login")
		fmt.Println("3. Send Message")
		fmt.Println("4. View Inbox")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			register()
		case 2:
			login()
		case 3:
			sendMessage()
		case 4:
			viewInbox()
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}
