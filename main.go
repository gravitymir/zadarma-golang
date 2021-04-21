package main

import (
	"fmt"

	z "github.com/gravitymir/zadarma-golang/zadarma"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("%v", err)
	}
}

func run() error {
	// paramsUrlValues := url.Values{
	// 	//"/v1/info/lists/tariffs/": []string{"USD"},
	// }

	// paramsMap := map[string]string{
	// 	//"pbx_number": "100",
	// }

	// paramsString := "" //`pbx_number=100`

	zadarmaBalance := z.New{
		APIMethod:    "/v1/info/balance/",
		APIUserKey:   "e30e16c201343883f77e",
		APISecretKey: "dbf5606ea4c1f2234201",
		// ParamsUrlValues: paramsUrlValues,
		// ParamsMap:       paramsMap,
		// ParamsString:    paramsString,
	}

	if err := zadarmaBalance.Request(); err != nil {
		fmt.Println(err)
		return err
	}

	// paramsUrlValues := url.Values{
	// 	"start": []string{"2018-10-01 08:00:00"},
	// 	"end":   []string{"2018-11-01 08:00:00"},
	// }
	// statistics := z.New{
	// 	APIMethod:       "/v1/statistics/",
	// 	APIUserKey:      "e30e16c201343883f77e",
	// 	APISecretKey:    "dbf5606ea4c1f2234201",
	// 	ParamsUrlValues: paramsUrlValues,
	// }
	// if err := statistics.Request(); err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	// statisticsPbx := Z.Zadarma{
	// 	ZadarmaMethod:    "/v1/statistics/pbx/",
	// 	ZadarmaUserKey:   "e30e16c201343883f77e",
	// 	ZadarmaSecretKey: "dbf5606ea4c1f2234201",
	// 	ParamsUrlValues:  paramsUrlValues,
	// 	ParamsMap:        paramsMap,
	// 	ParamsString:     paramsString,
	// }
	// Z.Request(statisticsPbx)

	// paramsUrlValues = url.Values{
	// 	"pbx_number": []string{"100"},
	// }

	// pbxRedirection := Z.Zadarma{
	// 	ZadarmaMethod:    "/v1/pbx/redirection/",
	// 	ZadarmaUserKey:   "e30e16c201343883f77e",
	// 	ZadarmaSecretKey: "dbf5606ea4c1f2234201",
	// 	ParamsUrlValues:  paramsUrlValues,
	// 	ParamsMap:        paramsMap,
	// 	ParamsString:     paramsString,
	// }
	// Z.Request(pbxRedirection)

	// pbxCallinfo := Z.Zadarma{
	// 	ZadarmaMethod:    "/v1/pbx/callinfo/",
	// 	ZadarmaUserKey:   "e30e16c201343883f77e",
	// 	ZadarmaSecretKey: "dbf5606ea4c1f2234201",
	// }
	// Z.Request(pbxCallinfo)

	// pbxWebhooks := Z.Zadarma{
	// 	ZadarmaMethod:    "/v1/pbx/webhooks/",
	// 	ZadarmaUserKey:   "e30e16c201343883f77e",
	// 	ZadarmaSecretKey: "dbf5606ea4c1f2234201",
	// }
	// Z.Request(pbxWebhooks)

	return nil
}
