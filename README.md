# Text Messaging CLI Application in Go

A simple command-line interface (CLI) application written in Go for managing text messages between users. This project demonstrates key concepts in Go, including user authentication, message handling, and real-time notifications through goroutines.

## Features

1. **User Registration & Login**: 
   - Allows users to create and authenticate accounts.
   - Credentials are stored in-memory for simplicity.

2. **Send and Receive Messages**:
   - Users can send text messages to other registered users.
   - Messages are timestamped and stored for each receiver.

3. **Inbox View**:
   - Users can view their received messages, along with sender information and timestamps.

4. **Real-Time Notifications**:
   - Uses goroutines to simulate real-time message notifications, alerting users of new messages.

## Getting Started

### Prerequisites

- Go 1.16 or later

### Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd text-messaging-cli
Run the application:

go run text_messaging.go
Usage
Start the application, and follow the menu prompts to:

Register a new user.
Log in with a username and password.
Send a message to another user.
View the inbox to see received messages.
Open multiple terminals to simulate interaction between different users.

Code Overview
1. User Authentication
Handles user registration and login, storing credentials in an in-memory map.
2. Messaging
Manages sending and receiving messages, storing each message along with the sender, receiver, and timestamp.
3. Inbox Management
Allows users to view messages sent to them, showing sender info and message time.
4. Real-Time Notifications
Goroutines are used to simulate real-time notifications, checking for new messages every few seconds.
Main Functions
register(): Registers a new user and stores their credentials.
login(): Authenticates the user based on username and password.
sendMessage(): Sends a message to a specified user.
viewInbox(): Displays received messages in the user's inbox.
startMessageReceiver(): Runs in the background to simulate real-time notifications.
Sample Output
Upon running the application, the following menu is presented:


--- Text Messaging CLI ---
1. Register
2. Login
3. Send Message
4. View Inbox
5. Exit
Choose an option: 
Example of sending and receiving messages:


Enter receiver's username: alice
Enter message content: Hello Alice!
Message sent!
Real-time notification:


[New message] bob: Hi, how are you?

Project Structure
text_messaging.go: The main file, containing all functions and logic for the CLI app.
Potential Extensions
Database Integration: Store users and messages in a database for persistence.
Enhanced Security: Hash passwords for secure authentication.
WebSocket for True Real-Time Communication: Implement WebSocket for authentic real-time message updates.
Error Handling and Validation: Improve validation for inputs and handle errors more robustly.
