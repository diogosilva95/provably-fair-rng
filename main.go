package main

import (
	"fmt"
	"provablyfair/games"
)

var clientSeed = "client_seed"
var serverSeed = "server_seed"

func main() {
	nonce := int64(12345)

	n := games.DiceRoll(clientSeed, serverSeed, nonce)
	fmt.Println("Dice roll:", n)
}
