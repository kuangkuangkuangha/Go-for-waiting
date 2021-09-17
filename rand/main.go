package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	p := []string{"curry", "zk", "cyw"}
	n, _ := rand.Int(rand.Reader, big.NewInt(3))
	fmt.Println(p[n.Int64()])
}
