package games

func DiceRoll(clientSeed, serverSeed string, nonce int64) float64 {
	return float64(RNG(clientSeed, serverSeed, nonce, 0)) / 100
}
