package main

import (
	"crypto/sha256"
	"fmt"
)
func main() {
    message := "Johnny was here boys"
    hash := sha256.Sum256([]byte(message))
    fmt.Printf("%x\n", hash)
}

