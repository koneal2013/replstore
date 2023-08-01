package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"kv-store/kvs"
)

type Command struct {
	Action string
	Params []string
}

func main() {
	txStack := &kvs.TransactionStack{}
	store := kvs.NewInMemoryKVStore()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		command := parseInput(input)

		target := chooseTarget(store, txStack)

		executeCommand(command, target, txStack)
	}
}

func parseInput(input string) Command {
	split := strings.Fields(strings.ToUpper(input))
	return Command{
		Action: split[0],
		Params: split[1:],
	}
}

func chooseTarget(store kvs.KeyValueStore, txStack *kvs.TransactionStack) kvs.KeyValueStore {
	tx, err := txStack.Current()
	if err != nil {
		return store
	}
	return tx
}

func executeCommand(command Command, target kvs.KeyValueStore, txStack *kvs.TransactionStack) {
	switch command.Action {
	case "READ":
		executeRead(target, command.Params)
	case "WRITE":
		executeWrite(target, command.Params)
	case "DELETE":
		executeDelete(target, command.Params)
	case "START":
		txStack.Push(kvs.NewTransaction(target))
	case "COMMIT":
		commitTransaction(txStack)
	case "ABORT":
		abortTransaction(txStack)
	case "QUIT":
		os.Exit(0)
	default:
		fmt.Println("Unknown command")
	}
}

func getKeyValueFromParams(params []string) (key string, value string, err error) {
	if len(params) < 2 {
		err = fmt.Errorf("not enough parameters")
		return "", "", err
	}

	return params[0], params[1], nil
}

func executeRead(target kvs.KeyValueStore, params []string) {
	if len(params) < 1 {
		fmt.Println("Error: not enough parameters")
	} else {
		val, err := target.Get(params[0])
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println(val)
		}
	}
}

func executeWrite(target kvs.KeyValueStore, params []string) {
	key, value, err := getKeyValueFromParams(params)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		err = target.Put(key, value)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func executeDelete(target kvs.KeyValueStore, params []string) {
	if len(params) < 1 {
		fmt.Println("Error: not enough parameters")
	} else {
		err := target.Delete(params[0])
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func commitTransaction(txStack *kvs.TransactionStack) {
	tx, err := txStack.Current()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		tx.Commit()
		err = txStack.Pop()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func abortTransaction(txStack *kvs.TransactionStack) {
	_, err := txStack.Current()
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		err = txStack.Pop()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
}
