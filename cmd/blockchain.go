package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"

	"github.com/KnutZuidema/blockchain/pkg/models"
)

func main() {
	create := argparse.NewParser("create", "create a blockchain")
	creator := create.String("c", "creator", &argparse.Options{Required: true, Help: "creator of the chain"})
	mine := argparse.NewParser("mine", "mine blocks")
	receiver := mine.String("r", "receiver", &argparse.Options{Required: true, Help: "receiver of the reward"})
	amount := mine.Int("a", "amount", &argparse.Options{Default: 1, Help: "amount of blocks to mine"})
	if len(os.Args) < 2 {
		fmt.Println("Need to specify an action")
		return
	}
	switch os.Args[1] {
	case "create":
		if err := create.Parse(os.Args[1:]); err != nil {
			fmt.Println(create.Usage(err))
			return
		}
		chain, err := models.NewBlockchain(*creator)
		if err != nil {
			fmt.Println("Could not create blockchain")
			return
		}
		fmt.Println(chain.String())
	case "mine":
		if err := mine.Parse(os.Args[1:]); err != nil {
			fmt.Println(mine.Usage(err))
			return
		}
		chain, err := models.GetBlockchain()
		if err != nil {
			fmt.Println("Could not open database")
			return
		}
		for i := 0; i < *amount; i++ {
			if err = chain.MineBlock(*receiver); err != nil {
				fmt.Println("Couldn't mine block", err)
			}
		}
	case "print":
		chain, err := models.GetBlockchain()
		if err != nil {
			fmt.Println("Could not open database")
			return
		}
		fmt.Println(chain.String())
	}
}
