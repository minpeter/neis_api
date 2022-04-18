package main

import (
	"fmt"

	api "github.com/minpeter/neis_api/API"
)

func main() {
	fmt.Println(api.Meal())
	fmt.Println(api.Time())
	fmt.Println(api.Schedule())
}
