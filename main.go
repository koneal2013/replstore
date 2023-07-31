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
		_, err := fmt.Scan(&cmd)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		switch strings.ToUpper(cmd) {
		case "READ":
			if _, err = fmt.Scan(&key); err != nil {
				fmt.Println("Error:", err)
				continue
			}
			val, err := store.Get(key)
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
			if err = store.Put(key, value); err != nil {
				fmt.Println("Error:", err)
			}
		case "DELETE":
			if _, err = fmt.Scan(&key); err != nil {
				fmt.Println("Error:", err)
				continue
			}
			if err = store.Delete(key); err != nil {
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
