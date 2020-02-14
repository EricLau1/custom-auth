package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func Intn() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r.Intn(9)
}

func GenerateCode() string {
	code := ""

	for i := 0; i < 6; i++ {
		code += fmt.Sprint(Intn())
	}

	return code
}
