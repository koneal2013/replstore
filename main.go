package main

import (
	"fmt"
	"strings"

	"kv-store/kvs"
)

func main() {
	store := &kvs.InMemoryKeyValueStore{}

	for {
		fmt.Print("> ")
		var cmd, key, value string
		fmt.Scan(&cmd)

		switch strings.ToUpper(cmd) {
		case "READ":
			fmt.Scan(&key)
			val, err := store.Get(key)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println(val)
			}
		case "WRITE":
			fmt.Scan(&key, &value)
			err := store.Put(key, value)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "DELETE":
			fmt.Scan(&key)
			err := store.Delete(key)
			if err != nil {
				fmt.Println("Error:", err)
			}
		case "QUIT":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Unknown command")
		}
	}
}
