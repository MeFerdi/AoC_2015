package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func mineCoin(secretKey string) int {
	number := 1

	for {
		data := secretKey + strconv.Itoa(number)

		// Compute the MD5 hash
		hash := md5.Sum([]byte(data))
		hashString := hex.EncodeToString(hash[:])

		// Check if the hash starts with "000000"
		if hashString[:6] == "000000" {
			return number
		}

		number++
	}
}
func main() {
	secretKey := "yzbqklnj"
	result := mineCoin(secretKey)
	fmt.Printf("The lowest number: %d\n", result)
}
