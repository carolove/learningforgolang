package main

import (
	"fmt"

	fb "github.com/huandu/facebook"
)

func main() {
	res, err := fb.Get("/538744468", fb.Params{
		"fields":       "first_name",
		"access_token": "CAAWumgh06egBAOIW1QIJZCX7ZBdcZCX1Jj3grlBVipT8j0KAs8pEGYOKeSkgI7ZBZBlABizkquFGQkdHDDDeE58Bf8GkIu6LZCEzwE4K3HlSovXddEvvYdlHSoJ0ZAwQUELTX0Ip9EvCxTNSxMnPDJ1xjgKpFBu7KjYz7KQjJtdT90AqAOkca5Oh7mLJwLyveMurtaiZAzPhC0LZCE2qRcEuG",
	})

	fmt.Println("here is my facebook first name:", res["first_name"], err)
}
