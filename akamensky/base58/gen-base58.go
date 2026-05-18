package main

import (
	"fmt"

	"github.com/akamensky/base58"
	"github.com/google/uuid"
)

func main() {
	u := uuid.New()

	encoded := base58.Encode(u[:])
	fmt.Println(encoded)
}
