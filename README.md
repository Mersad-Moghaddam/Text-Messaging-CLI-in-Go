# Text Messaging CLI Application

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
