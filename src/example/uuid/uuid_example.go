package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"strings"
)

func main() {
	id := uuid.NewV4().String()
	id = strings.Replace(id, "-", "", -1)
	fmt.Println(id)
}
