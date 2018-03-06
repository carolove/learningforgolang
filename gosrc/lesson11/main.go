package main

import (
	"fmt"

	fb "github.com/huandu/facebook"
)

func main() {
	res, _ := fb.Get("/538744468", fb.Params{
		"fields":       "first_name",
		"access_token": "a-valid-access-token",
	})
	fmt.Println("here is my facebook first name:", res["first_name"])
}
