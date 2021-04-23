package main

import (
	"fmt"
	"net/url"

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

	// zadarmaBalance := z.New{
	// 	APIMethod:    "/v1/info/balance/",
	// 	APIUserKey:   "e30e16c201343883f77e",
	// 	APISecretKey: "dbf5606ea4c1f2234201",
	// 	// ParamsUrlValues: paramsUrlValues,
	// 	// ParamsMap:       paramsMap,
	// 	// ParamsString:    paramsString,
	// }

	// if err := zadarmaBalance.Request(); err != nil {
	// 	fmt.Println(err)
	// 	return err
	// }

	paramsUrlValues := url.Values{
		"start": []string{"2018-10-01 08:00:00"},
		"end":   []string{"2018-10-04 08:00:00"},
	}
	statistics := z.New{
		APIMethod:    "/v1/info/balance/",
		APIUserKey:   "e30e16c201343883f72e",
		APISecretKey: "dbf5606ea4c1f2234701",

		ParamsUrlValues: paramsUrlValues,
	}
	statMap, err := statistics.Request()
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("APIUserKey: ", statistics.APIUserKey)
	fmt.Println("APISecretKey: ", statistics.APISecretKey)
	fmt.Println("HTTPMethod: ", statistics.HTTPMethod)
	fmt.Println("APIMethod: ", statistics.APIMethod)
	fmt.Println("LinkToAPI: ", statistics.LinkToAPI)
	fmt.Println("\nParamsUrlValues: ", statistics.ParamsUrlValues)
	fmt.Println("\nParamsMap: ", statistics.ParamsMap)
	fmt.Println("\nParamsString: ", statistics.ParamsString)
	fmt.Println("\nResponseRaw string: ", string(statistics.ResponseRaw))
	fmt.Println("\nSignature: ", statistics.Signature)
	fmt.Println("\nSortedParamsString: ", statistics.SortedParamsString)
	fmt.Println("\n\n->")

	if statMap["status"] == "success" {
		fmt.Printf("status: \033[1;32m%s\033[0m\n", statMap["status"])
		for k, v := range statMap {
			if k == "status" {
				continue
			}
			fmt.Printf("%v: %v\n", k, v)
		}
	} else {
		fmt.Printf("status: \033[1;31m%s\033[0m\n", statMap["status"])
		fmt.Printf("Error: %v\n", statMap["message"])
	}

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
