package main

import (
	"fmt"
	"net/url"

	z "github.com/gravitymir/zadarma-golang/zadarma"
	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func run() error {

	paramsUrlValues := url.Values{
		"start": []string{"2018-10-01 08:00:00"},
		"end":   []string{"2018-10-04 08:00:00"},
	}
	statistics := z.New{
		APIMethod:    "/v1/info/balance/",
		APIUserKey:   "e30e16c201343883f77e",
		APISecretKey: "dbf5606ea4c1f2234201",

		ParamsUrlValues: paramsUrlValues,
	}

	g := new(errgroup.Group)

	g.Go(func() error {
		statistics.Request()
		fmt.Println("APIUserKey: ", statistics.APIUserKey)
		fmt.Println("APISecretKey: ", statistics.APISecretKey)
		fmt.Println("HTTPMethod: ", statistics.HTTPMethod)
		fmt.Println("APIMethod: ", statistics.APIMethod)
		fmt.Println("LinkToAPI: ", statistics.LinkToAPI)
		fmt.Println("ParamsUrlValues: ", statistics.ParamsUrlValues)
		fmt.Println("ParamsMap: ", statistics.ParamsMap)
		fmt.Println("ParamsString: ", statistics.ParamsString)
		fmt.Println("ResponseRaw: ", string(statistics.ResponseRaw))
		fmt.Println("Signature: ", statistics.Signature)
		fmt.Println("SortedParamsString: ", statistics.SortedParamsString)
		fmt.Println("\n->")

		fmt.Println(statistics.Status)
		//fmt.Println(statistics.Error)
		fmt.Println(statistics.Result)
		fmt.Printf("%T\n", statistics.Result["balance"])
		fmt.Println(statistics.Result["balance"])
		fmt.Println(statistics.Result["currency"])

		return statistics.Error
	})

	return g.Wait()
}
