package main

import (
	"fmt"

	uuid "github.com/google/uuid"
)

func main() {
	u := uuid.New()
	fmt.Printf("Hello, %s\n", u)
}
