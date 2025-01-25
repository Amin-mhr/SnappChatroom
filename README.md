# SnappChatroom

This repository contains the client and server components for the SnappChatroom application, designed as part of the Snapp company interview task. The application is structured to run the client and server separately, with the server being deployed using Docker Compose.

## Architecture Overview

### Server
The server is built using a **Hexagonal Architecture**, which promotes separation of concerns and flexibility in the system. It uses **PostgreSQL** as the database and **NATS** for messaging. The server is designed to handle WebSocket connections from clients, process incoming messages, and manage chat rooms.

### Client
The client is a command-line application written in Go.The client connects to the server via WebSocket to send and receive messages.

## How It Works

1. **Client-Server Communication**: The client sends messages to the server through a WebSocket connection. The server listens for incoming messages and processes them accordingly.

2. **Message Handling**: When a message is received, the server places it on the appropriate channel corresponding to the chat room. This ensures that messages are routed correctly to all participants in the chat room.

3. **Database and Messaging**: The server uses PostgreSQL to store relevant data and NATS for efficient message brokering between different components of the system.

## Running the Application

### Server
for starting server you need to pass json to app that it handle in docker-compose and if you want to change any value of this the sample-config.json is in ./server/sample-config.json.
```json
{
  "server": {
    "httpPort": 
  },
  "nats": {
    "hostPort": ""
  },
  "db": {
    "host": "",
    "port": ,
    "database": "",
    "schema": "",
    "user": "",
    "password": ""
  }
}

```

start the server, use Docker Compose:
```bash
docker-compose up
```

### Client
To start the client, use below command (in ./client/cmd):
```bash
go run main.go
```

## Additional Notes

1. **Docker Network Configuration**: Ensure that the Docker network is properly configured to allow communication between the server, database, and NATS. If you encounter connection issues, check the Docker network settings.

2. **Logging**: The server logs important events and errors to the console. You can redirect these logs to a file or a logging service for better monitoring and debugging.

3. **Security Considerations**:
    - Always keep your database credentials and other sensitive information secure. Avoid committing sensitive data to version control.
    - Use environment variables or secure vaults to manage sensitive configurations.

4. **Scaling**: The application is designed to handle multiple clients. If you need to scale the server, consider using a load balancer and multiple instances of the server.

5. **WebSocket Connections**: Ensure that your network allows WebSocket connections. Some firewalls or proxies might block WebSocket traffic.

6. **Database Migrations**: If you need to make changes to the database schema, use migration tools to manage these changes. This ensures that your database schema is always in sync with your application code.

7. **Client Customization**: The client application can be customized to include additional features such as message formatting, user authentication, and more. Refer to the client documentation for details.

8. **Troubleshooting**: If you encounter any issues, check the logs for error messages. Common issues include incorrect database credentials, NATS connection problems, and WebSocket connection failures.

9. **Performance Monitoring**: Consider integrating performance monitoring tools to track the health and performance of your application in real-time.


