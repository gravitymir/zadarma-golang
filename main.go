package main

import (
	"encoding/json"
	"fmt"

	"golang.org/x/sync/errgroup"

	zApi "github.com/gravitymir/zadarma-golang/zadarma"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("%v\n", err)
	}
}

func run() error {

	g := new(errgroup.Group)

	g.Go(func() error {

		statistics := zApi.New{
			APIMethod:    "/v1/statistics/",
			APIUserKey:   "e30e16c201343883f77e",
			APISecretKey: "dbf5606ea4c1f2234201",
			ParamsMap: map[string]string{
				"end":   "2018-10-04 08:00:00",
				"start": "2018-09-01 08:00:00",
			},
		}

		data := []byte{}

		if err := statistics.Request(&data); err != nil {
			return err
		}

		targetData := zApi.CatchStatistics{}

		if err := json.Unmarshal(data, &targetData); err != nil {
			return err
		}

		fmt.Println(targetData)
		return nil
	})

	return g.Wait()
}
