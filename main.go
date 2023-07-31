package main

import (
	"fmt"
	"strings"

	"kv-store/kvs"
)

func main() {
	store := &kvs.InMemoryKeyValueStore{}
	txStack := &kvs.TransactionStack{}

	for {
		fmt.Print("> ")
		var cmd, key, value string
		_, err := fmt.Scan(&cmd)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Determine the target for the operation: either the store or the current transaction.
		var target kvs.KeyValueStore
		if txStack.Current() != nil {
			target = txStack.Current()
		} else {
			target = store
		}

		switch strings.ToUpper(cmd) {
		case "READ":
			if _, err = fmt.Scan(&key); err != nil {
				fmt.Println("Error:", err)
				continue
			}
			val, err := target.Get(key)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(val)
			}
		case "WRITE":
			if _, err = fmt.Scan(&key, &value); err != nil {
				fmt.Println("Error:", err)
				continue
			}
			if err = target.Put(key, value); err != nil {
				fmt.Println("Error:", err)
			}
		case "DELETE":
			if _, err = fmt.Scan(&key); err != nil {
				fmt.Println("Error:", err)
				continue
			}
			if err = target.Delete(key); err != nil {
				fmt.Println("Error:", err)
			}
		case "START":
			txStack.Push()
		case "COMMIT":
			if txStack.Current() != nil {
				txStack.Current().Commit()
				txStack.Pop()
			}
		case "ABORT":
			if txStack.Current() != nil {
				txStack.Pop()
			}
		case "QUIT":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}
