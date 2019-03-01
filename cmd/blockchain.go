package main

import (
	"blockchain/pkg/models"
	"flag"
	"fmt"
	"os"
)

var data string
var add = flag.NewFlagSet("add", flag.ExitOnError)

func init() {
	add.StringVar(&data, "data", "", "data to be stored in the block")
}

func main() {
	chain, err := models.BlockchainFromDb("chain", "chain.db")
	if err != nil {
		fmt.Println("Couldn't open database")
	}
	switch os.Args[1] {
	case "add":
		if err := add.Parse(os.Args[2:]); err != nil {
			add.Usage()
			return
		}
		if data == "" {
			add.Usage()
			return

		}
		chain.AddBlock([]byte(data))
		err := chain.ToDb("chain", "chain.db")
		if err != nil {
			fmt.Println("Couldn't save data to db")
		}
	case "print":
		fmt.Println(chain.String())
	}
}
