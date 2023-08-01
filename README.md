# In-Memory Key-Value Store with Transaction Support

This project implements an in-memory key-value store with transaction support, designed using the Hexagonal Architecture. This design allows for different storage engines to be implemented and swapped seamlessly, providing high flexibility and adaptability for various storage requirements.

## Features

- **In-Memory Key-Value Store**: A simple in-memory key-value store that supports basic operations like Get, Put, and Delete.
- **Transaction Support**: The ability to start a transaction, perform multiple operations within the transaction, and either commit the changes or abort the transaction.
- **Hexagonal Architecture**: The design of the project follows the Hexagonal Architecture, which allows for easy extension and replacement of the storage engine.

## Usage

The project includes a simple command-line interface for interacting with the key-value store. Here are the available commands:

- `READ <key>`: Retrieves the value for the given key.
- `WRITE <key> <value>`: Stores the value for the given key.
- `DELETE <key>`: Removes the key-value pair for the given key.
- `START`: Starts a new transaction.
- `COMMIT`: Commits the changes made in the current transaction.
- `ABORT`: Aborts the current transaction.
- `QUIT`: Exits the application.

## Installation

To install the project, you need to have Go installed on your machine. You can then clone the repository and build the project:

```bash
git clone https://github.com/openly-hiring/kenston-oneal.git
cd kenston-oneal
go build
```

To run the applicaton:

```base
./kenston-oneal
```

## Testing

The project includes unit tests for the `InMemoryKeyValueStore`, `Transaction`, and `TransactionStack` types. You can run the tests using the `go test` command:

```bash
go test ./...

```

