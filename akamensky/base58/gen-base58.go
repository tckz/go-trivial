package main

import (
	"fmt"
	"uuid"

	"github.com/akamensky/base58"
)

func main() {
	u := uuid.New()

	encoded := base58.Encode(u[:])
	fmt.Println(encoded)
}
