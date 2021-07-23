package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	store "github.com/jack-kitto/Store.git" // for demo
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x16261c5963f668F5620b26DB365aB66Cc8071B3d")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version) // "1.0"
}
