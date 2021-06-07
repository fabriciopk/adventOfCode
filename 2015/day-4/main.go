package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func solution(data []byte, leadingZeros int) {
	sum := 0
	for {
		sumStr := []byte(strconv.Itoa(sum))
		hash := md5.Sum(append(data, sumStr...))
		hashStr := hex.EncodeToString(hash[:])
		if r := strings.HasPrefix(hashStr, strings.Repeat("0", leadingZeros)); r {
			fmt.Println("AdventCoint Hash: ", sum)
			break
		}
		sum = sum + 1
	}
}

func main() {
	secretKey := []byte("iwrupvqb")
	solution(secretKey, 5)
	solution(secretKey, 6)
}
