package main

import (
	"fmt"

	"github.com/giovannirossini/curso/basic/variables/packages/phone"
)

// UserName é o nome do usuário do sistema
var UserName = "Giovanni Rossini"

func main() {
	fmt.Printf("Name: %s\r\n", UserName)
	fmt.Printf("Phone: (%d) %s\r\n", phone.DDD, phone.Number)
}
