package main

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func Coin(secretKey string) int {
	number := 1

	for {
		// Combine the secret key with the current number
		data := secretKey + strconv.Itoa(number)

		// Compute the MD5 hash
		hash := md5.Sum([]byte(data))
		hashString := hex.EncodeToString(hash[:])

		// Check if the hash starts with "00000"
		if hashString[:5] == "00000" {
			return number
		}

		// Increment the number for the next iteration
		number++
	}
}

// func main() {
// 	secretKey := "yzbqklnj"
// 	result := mineAdventCoin(secretKey)
// 	fmt.Printf("The lowest number that produces an MD5 hash starting with five zeroes is: %d\n", result)
// }
