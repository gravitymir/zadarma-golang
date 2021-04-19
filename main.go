package main

import (
	"fmt"
	"net/url"

	Z "github.com/gravitymir/zadarma-golang/zadarma"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("%v", err)
	}
}

func run() error {

	m := map[string]string{
		//"pbx_number": "100",
	}
	l := Z.API{
		APIMethod:       "/v1/pbx/redirection/",
		APIUserKey:      "e30e16c501343883f77e",
		APISecretKey:    "dbf5605ea4c1f2234201",
		ParamsUrlValues: url.Values{
			//"pbx_number": []string{"100"},
		},
		ParamsMap:    m,
		ParamsString: `pbx_number=100`,
	}

	Z.Request(l)
	return nil
}
