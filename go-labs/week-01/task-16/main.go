package main 

import (
	"fmt"
	"strings"
)

func main() {
	username := "Admin"
	password := "golang2026"

	if strings.EqualFold(username, "admin") && password == "golang2026"{
		fmt.Println("Доступ разрешен")
	} else {
		fmt.Println("Доступ запрещен")
	}
}