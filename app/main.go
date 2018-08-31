package main

import (
	"fmt"

	"github.com/dewantara_hz/onboarding-go/reqres"
)

func main() {
	content := reqres.Access("http://google.com")

	fmt.Println(content)
}
