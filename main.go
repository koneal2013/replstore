package main

import (
	"fmt"
	"os"
	"strings"

	"kv-store/kvs"
)

// cmdValues is used to store values of a command line.
type cmdValues struct {
	cmd, key, value string
}

func main() {
	store := &kvs.InMemoryKeyValueStore{}
	txStack := &kvs.TransactionStack{}

	for {
		fmt.Print("> ")

		var values cmdValues
		values.cmd, _ = readCmd()

		target := getTarget(store, txStack)

		processCommand(values, target, txStack)
	}
}

// readCmd reads a command line.
func readCmd() (cmd string, err error) {
	_, err = fmt.Scan(&cmd)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return strings.ToUpper(cmd), err
}

// getTarget gets the target for the operation.
func getTarget(store *kvs.InMemoryKeyValueStore,
	txStack *kvs.TransactionStack) kvs.KeyValueStore {
	if txStack.Current() != nil {
		return txStack.Current()
	}

	return store
}

// processCommand processes a command.
func processCommand(values cmdValues, target kvs.KeyValueStore,
	txStack *kvs.TransactionStack) {
	switch values.cmd {
	case "READ":
		readCommand(values.key, target)
	case "WRITE":
		writeCommand(values.key, values.value, target)
	case "DELETE":
		deleteCommand(values.key, target)
	case "START":
		startCommand(txStack)
	case "COMMIT":
		commitCommand(txStack)
	case "ABORT":
		abortCommand(txStack)
	case "QUIT":
		quitCommand()
	default:
		fmt.Println("Unknown command")
	}
}

// readCommand reads a value.
func readCommand(key string, target kvs.KeyValueStore) {
	key, err := readKey()
	if err != nil {
		return
	}

	val, err := target.Get(key)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(val)
	}
}

// writeCommand writes a keyvalue pair.
func writeCommand(key, value string, target kvs.KeyValueStore) {
	key, value, err := readKeyValue()
	if err != nil {
		return
	}

	err = target.Put(key, value)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// deleteCommand deletes a key.
func deleteCommand(key string, target kvs.KeyValueStore) {
	key, err := readKey()
	if err != nil {
		return
	}

	err = target.Delete(key)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// startCommand starts a new transaction.
func startCommand(txStack *kvs.TransactionStack) {
	txStack.Push()
}

// commitCommand commits a transaction.
func commitCommand(txStack *kvs.TransactionStack) {
	if txStack.Current() != nil {
		txStack.Current().Commit()
		txStack.Pop()
	}
}

// abortCommand aborts a transaction.
func abortCommand(txStack *kvs.TransactionStack) {
	if txStack.Current() != nil {
		txStack.Pop()
	}
}

// quitCommand quits the application.
func quitCommand() {
	fmt.Println("Exiting...")
	os.Exit(0)
}

// readKey reads a key.
func readKey() (key string, err error) {
	_, err = fmt.Scan(&key)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return key, err
}

// readKeyValue reads a key and a value.
func readKeyValue() (key, value string, err error) {
	_, err = fmt.Scan(&key, &value)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return key, value, err
}
