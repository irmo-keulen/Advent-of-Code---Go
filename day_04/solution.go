package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func checkAdventCoin(str string) bool {
	hash := md5.Sum([]byte(str))
	return strings.HasPrefix(hex.EncodeToString(hash[:]), "0000000")
}

func main() {
	input := "iwrupvqb"
	for i := 1; ; i++ {
		if checkAdventCoin(input + strconv.Itoa(i)) {
			fmt.Printf("%s%d\n", input, i)
			break
		}
	}
}
