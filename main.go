package main

import (
	"fmt"

	"github.com/srickopp/go-getting-started/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Ricko",
		LastName:  "Samuel",
	}

	fmt.Println(u)
}
