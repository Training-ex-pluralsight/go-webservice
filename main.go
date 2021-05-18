package main

import (
	"fmt"
	"github.com/pluralsight/webservice/models"
)

func main() {
	u := models.User{ID: 1, FirstName: "Darth", LastName: "Vader"}
	fmt.Println(u)
}
