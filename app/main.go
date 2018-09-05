package main

import (
	"fmt"

	"gitlab.bukalapak.io/heindewantara/onboarding-go.git/reqres"
)

func main() {
	content := reqres.Access("http://google.com")

	fmt.Println(content)
}
