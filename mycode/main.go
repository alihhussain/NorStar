package main

import (
	"fmt"

	"./dog"
)

func main() {
	humanYears := 7
	dogYears := dog.Years(humanYears)

	fmt.Println(dogYears)

}
