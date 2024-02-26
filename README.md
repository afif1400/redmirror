# Redmirror

This is a simple implementation of a Redis-like server in Golang. The server listens for TCP connections on localhost:6379 and handles various commands sent by clients. The project consists of multiple Go files:

- `main.go`: Contains the main server implementation, including the network handling and command execution logic.
- `resp.go`: Defines the RESP (REdis Serialization Protocol) reader and writer for parsing incoming commands and serializing responses.
- `handlers.go`: Defines the supported Redis commands and their corresponding handler functions.
- `run.sh`: A shell script to build the project.

## Getting Started

To build and run the server, follow these steps:

1. Ensure you have Go installed on your machine.
2. Clone the repository to your local machine.
3. Navigate to the project directory.
4. Run the `run.sh` script to build the project.

```sh
git clone <repository_url>
cd <repository_directory>
chmod +x build.sh
./run.sh
```

The server will start listening on `localhost:6379` by default.

## Supported Commands

The server supports the following Redis commands:

- `PING`: Ping the server. It returns "PONG".
- `ECHO`: Echo the input string. It returns the provided string.
- `SET`: Set a key-value pair. Optionally, it supports setting an expiration time.
- `GET`: Get the value associated with a key.
- `HSET`: Set the field in a hash stored at the specified key to a value.
- `HGET`: Get the value of a hash field.

For detailed information on how to use each command, please refer to the Redis documentation as this server mimics Redis behavior.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or create a pull request.

## License

This project is licensed under the [MIT License](LICENSE). Feel free to use, modify, and distribute it as per the terms of the license.
