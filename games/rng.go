package games

import (
	"crypto/sha512"
	"encoding/hex"
	"strconv"
)

// returns a number 0 and 9999 based on the provided seeds and nonce
func RNG(clientSeed string, serverSeed string, nonce int64, cursor int) int {
	hash := sha512.New()
	hash.Write([]byte(serverSeed + clientSeed + strconv.FormatInt(nonce, 10) + ":" + strconv.Itoa(cursor)))
	hashBytes := hash.Sum(nil)

	hashHex := hex.EncodeToString(hashBytes)

	var lucky int64
	index := 0

	for {
		if index+5 > len(hashHex) {
			break
		}
		hexSubstr := hashHex[index : index+5]
		parsed, _ := strconv.ParseInt(hexSubstr, 16, 0)

		lucky = parsed

		if lucky < 1000000 {
			break
		}
		index += 5
	}

	return int(lucky % 10000)
}
