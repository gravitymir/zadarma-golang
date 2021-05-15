package main

import (
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
		// 	"start": []string{"2018-09-01 08:00:00"},
		// 	"end":   []string{"2018-10-04 08:00:00"},
		// }

		paramsMap := map[string]string{
			"end":   "2018-10-04 08:00:00",
			"start": "2018-09-01 08:00:00",
		}

		//paramsString := "start=2018-09-01 08:00:00&end=2018-10-04 08:00:00"

		statistics := zApi.New{
			APIMethod:    "/v1/statistics/",
			APIUserKey:   "e30e16c201343883f77e",
			APISecretKey: "dbf5606ea4c1f2234201",
			//ParamsUrlValues: paramsUrlValues, //high priority
			ParamsMap: paramsMap, //medium priority, if not set ParamsUrlValues
			//ParamsString:    paramsString,    //low priority, if not set ParamsUrlValues & ParamsMap
		}

		_, err := statistics.Request()

		if err != nil {
			return err
		}
		if statistics.Signature == "ZmI3YzA1MmJkNzQyYmE1YzI3OWY3OTRjNTlmZDFkMTE2MmI4YmI0Yg==" {

		}

		// tariffData := tariffSt{}
		// if err := json.Unmarshal(tariff.Response, &tariffData); err != nil {
		// 			return err
		// }

		return nil
	})

	return g.Wait()
}
