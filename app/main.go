package main

import (
	"fmt"

	"gitlab.bukalapak.io/heindewantara/onboarding-go/reqres"
)

func main() {
	content := reqres.Access("http://google.com")

	fmt.Println(content)
}
