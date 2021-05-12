package main

import (
	"encoding/json"
	"fmt"

	zApi "github.com/gravitymir/zadarma-golang/zadarma"
	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func run() error {

	g := new(errgroup.Group)

	g.Go(func() error {
		// paramsUrlValues := url.Values{
		// 	"start": []string{"2018-10-01 08:00:00"},
		// 	"end":   []string{"2018-10-04 08:00:00"},
		// }

		balance := zApi.New{
			APIMethod:    "/v1/info/balance/",
			APIUserKey:   "e30e16c201343883f77e",
			APISecretKey: "dbf5606ea4c1f2234201",

			//ParamsUrlValues: paramsUrlValues,
		}

		balance.Request()
		// fmt.Println("APIUserKey: ", balance.APIUserKey)
		// fmt.Println("APISecretKey: ", balance.APISecretKey)
		// fmt.Println("HTTPMethod: ", balance.HTTPMethod)
		// fmt.Println("APIMethod: ", balance.APIMethod)
		// fmt.Println("LinkToAPI: ", balance.LinkToAPI)
		// fmt.Println("ParamsUrlValues: ", balance.ParamsUrlValues)
		// fmt.Println("ParamsMap: ", balance.ParamsMap)
		// fmt.Println("ParamsString: ", balance.ParamsString)
		// fmt.Println("ResponseRaw: ", string(balance.ResponseRaw))
		// fmt.Println("Signature: ", balance.Signature)
		// fmt.Println("SortedParamsString: ", balance.SortedParamsString)
		// fmt.Println("\n->")

		//fmt.Println(balance.Ok)
		//fmt.Println(balance.Result)
		//fmt.Println(balance.Result.Info)
		//fmt.Println(balance.Result.Message)
		return balance.Error
	})

	g.Go(func() error {

		tariff := zApi.New{
			APIMethod:    "/v1/tariff/",
			APIUserKey:   "e30e16c201343883f77e",
			APISecretKey: "dbf5606ea4c1f2234201",
		}

		tariff.Request()
		// fmt.Println("APIUserKey: ", tariff.APIUserKey)
		// fmt.Println("APISecretKey: ", tariff.APISecretKey)
		// fmt.Println("HTTPMethod: ", tariff.HTTPMethod)
		// fmt.Println("APIMethod: ", tariff.APIMethod)
		// fmt.Println("LinkToAPI: ", tariff.LinkToAPI)
		// fmt.Println("ParamsUrlValues: ", tariff.ParamsUrlValues)
		// fmt.Println("ParamsMap: ", tariff.ParamsMap)
		// fmt.Println("ParamsString: ", tariff.ParamsString)
		// fmt.Println("ResponseRaw: ", string(tariff.ResponseRaw))
		// fmt.Println("Signature: ", tariff.Signature)
		// fmt.Println("SortedParamsString: ", tariff.SortedParamsString)
		// fmt.Println("\n->")

		fmt.Println(tariff.Ok)
		fmt.Println(tariff.Result)
		//fmt.Println(tariff.Result["info"])
		type m struct {
			Cost                  float64 `json:"cost"`
			Currency              string  `json:"currency"`
			IsActive              string  `json:"is_active"`
			TariffForNextPeriod   string  `json:"tariff_for_next_period"`
			TariffId              uint64  `json:"tariff_id"`
			TariffIdForNextPeriod uint64  `json:"tariff_id_for_next_period"`
			TariffName            string  `json:"tariff_name"`
		}

		mm := m{}
		mbyte, err := json.Marshal(tariff.Result.Info)
		if err != nil {
			fmt.Println(err)
		}
		if err := json.Unmarshal(mbyte, &mm); err != nil {
			fmt.Println(err)
		}
		fmt.Println(mm)
		return tariff.Error
	})

	return g.Wait()
}
