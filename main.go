package main

import (
	"fmt"
	"net/url"

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

	// g.Go(func() error {
	// 	// paramsUrlValues := url.Values{
	// 	// 	"start": []string{"2018-10-01 08:00:00"},
	// 	// 	"end":   []string{"2018-10-04 08:00:00"},
	// 	// }

	// 	balance := zApi.New{
	// 		APIMethod:    "/v1/info/balance/",
	// 		APIUserKey:   "e30e16c201343883f77e",
	// 		APISecretKey: "dbf5606ea4c1f2234201",

	// 		//ParamsUrlValues: paramsUrlValues,
	// 	}
	// 	balance.Request()

	// 	type balanceSt struct {
	// 		Status   string  `json:"status"`
	// 		Message  string  `json:"message"`
	// 		Balance  float64 `json:"balance"`
	// 		Currency string  `json:"currency"`
	// 	}

	// 	balanceData := balanceSt{}
	// 	if err := json.Unmarshal(balance.Response, &balanceData); err != nil {
	// 		return err
	// 	}

	// 	fmt.Println("Balance: ", balanceData)

	// 	return nil
	// })
	// g.Go(func() error {

	// 	tariff := zApi.New{
	// 		APIMethod:    "/v1/tariff/",
	// 		APIUserKey:   "e30e16c201343883f77e",
	// 		APISecretKey: "dbf5606ea4c1f2234201",
	// 	}

	// 	if err := tariff.Request(); err != nil {
	// 		return err
	// 	}

	// 	type tariffSt struct {
	// 		Status  string `json:"status"`
	// 		Message string `json:"message"`
	// 		Info    struct {
	// 			TariffId                  uint64  `json:"tariff_id"`
	// 			TariffName                string  `json:"tariff_name"`
	// 			IsActive                  string  `json:"is_active"`
	// 			Cost                      float64 `json:"cost"`
	// 			Currency                  string  `json:"currency"`
	// 			Used_seconds              uint64  `json:"used_seconds"`
	// 			Used_seconds_mobile       uint64  `json:"used_seconds_mobile"`
	// 			Used_seconds_fix          uint64  `json:"used_seconds_fix"`
	// 			Tariff_id_for_next_period uint64  `json:"tariff_id_for_next_period"`
	// 			Tariff_for_next_period    string  `json:"tariff_for_next_period"`
	// 		} `json:"info"`
	// 	}
	// 	tariffData := tariffSt{}
	// 	if err := json.Unmarshal(tariff.Response, &tariffData); err != nil {
	// 		return err
	// 	}

	// 	fmt.Println("Tariff: ", tariffData)

	// 	return nil
	// })

	g.Go(func() error {

		paramsUrlValues := url.Values{
			"start": []string{"2018-09-01 08:00:00"},
			"end":   []string{"2018-10-04 08:00:00"},
		}

		number_lookup := zApi.New{
			//HTTPMethod:      "POST",
			APIMethod:       "/v1/statistics/",
			APIUserKey:      "e30e16c201343883f77e",
			APISecretKey:    "dbf5606ea4c1f2234201",
			ParamsUrlValues: paramsUrlValues,
		}

		if err := number_lookup.Request(); err != nil {
			return err
		}

		// tariffData := tariffSt{}
		// if err := json.Unmarshal(tariff.Response, &tariffData); err != nil {
		// 	return err
		// }

		fmt.Println("number_lookup: ", string(number_lookup.Response))

		return nil
	})

	return g.Wait()
}
