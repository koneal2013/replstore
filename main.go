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
	txStack := &kvs.TransactionStack{}
	store := kvs.NewInMemoryKVStore()

	for {
		fmt.Print("> ")

		var values cmdValues

		cmd, err := readCmd()
		if err != nil {
			fmt.Println("Error:", err)
		}

		values.cmd = cmd

		target := getTarget(store, txStack)

		processCommand(values, target, txStack)
	}
}

// readCmd reads a command line.
func readCmd() (cmd string, err error) {
	_, err = fmt.Scan(&cmd)

	return strings.ToUpper(cmd), err
}

// getTarget gets the target for the operation.
func getTarget(store kvs.KeyValueStore,
	txStack *kvs.TransactionStack) kvs.KeyValueStore {
	tx, err := txStack.Current()
	if err != nil {
		return store
	}

	return tx
}

// processCommand processes a command.
func processCommand(values cmdValues, target kvs.KeyValueStore,
	txStack *kvs.TransactionStack) {
	switch values.cmd {
	case "READ":
		readCommand(target)
	case "WRITE":
		writeCommand(target)
	case "DELETE":
		deleteCommand(target)
	case "START":
		startCommand(target, txStack)
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
func readCommand(target kvs.KeyValueStore) {
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

// writeCommand writes a key value pair.
func writeCommand(target kvs.KeyValueStore) {
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
func deleteCommand(target kvs.KeyValueStore) {
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
func startCommand(kvStore kvs.KeyValueStore, txStack *kvs.TransactionStack) {
	txStack.Push(kvs.NewTransaction(kvStore))
}

// commitCommand commits a transaction.
func commitCommand(txStack *kvs.TransactionStack) {
	tx, err := txStack.Current()
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	tx.Commit()

	err = txStack.Pop()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// abortCommand aborts a transaction.
func abortCommand(txStack *kvs.TransactionStack) {
	_, err := txStack.Current()
	if err != nil {
		fmt.Println("Error:", err)

		return
	}

	err = txStack.Pop()
	if err != nil {
		fmt.Println("Error:", err)
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
