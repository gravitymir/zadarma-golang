![GitHub last commit](https://img.shields.io/github/last-commit/gravitymir/zadarma-golang?logo=github)
![Github Repository Size](https://img.shields.io/github/repo-size/gravitymir/zadarma-golang?logo=github)
[![Github forks](https://img.shields.io/github/forks/gravitymir/zadarma-golang?logo=github)](https://github.com/gravitymir/zadarma-golang/network/members)
![Lines of code](https://img.shields.io/tokei/lines/github.com/gravitymir/zadarma-golang?logo=github)
[![GitHub open issues](https://img.shields.io/github/issues/gravitymir/zadarma-golang?logo=github)](https://github.com/gravitymir/zadarma-golang/issues)
[![GitHub closed issues](https://img.shields.io/github/issues-closed/gravitymir/zadarma-golang?logo=github)](https://github.com/gravitymir/zadarma-golang/issues)

[![made-with-Go](https://img.shields.io/badge/Zadarma-Go-00aed8?logo=go)](https://pkg.go.dev/search?q=zadarma)
[![Golang go.mod Go version](https://img.shields.io/github/go-mod/go-version/gravitymir/zadarma-golang?label=mod&logo=go)](https://pkg.go.dev/search?q=zadarma)
[![GoReportCard](https://goreportcard.com/badge/github.com/gravitymir/zadarma-golang)](https://goreportcard.com/report/github.com/gravitymir/zadarma-golang)

[![GitHub Repo stars](https://img.shields.io/github/stars/gravitymir/zadarma-golang?label=zadarma-golang&logo=github&color=505050&logoColor=fff)](https://github.com/gravitymir/zadarma-golang)
[![GitHub User's stars](https://img.shields.io/github/stars/gravitymir?label=gravitymir&logo=github&color=505050&logoColor=fff)](https://github.com/gravitymir)
[![Linkedin](https://img.shields.io/badge/AndreySukhodeev-linkedin?logo=linkedin&color=4d4d4d&logoColor=0B66C3)](https://www.linkedin.com/in/andrey-sukhodeev-108a131b9/)
[![Linkedin](https://img.shields.io/badge/Gravitymir-telegram?logo=telegram&color=4d4d4d&logoColor=1F97D5)](https://t.me/gravitymir)

![Zadarma Golang](https://raw.githubusercontent.com/gravitymir/zadarma-golang/master/zadarma_go.jpg)

# zadarma-golang
Library which help you work with API Zadarma (v1)

``` go
package main

import (
	"encoding/json"
	"errgroup"
	"fmt"

	zApi "github.com/gravitymir/zadarma-golang/zadarma"
	//"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("%v\n", err)
	}
}

type jsonTarget struct {
	Status string `json:"status"`
}

func run() error {

	g := new(errgroup.Group)

	g.Go(func() error {

		paramsMap := map[string]string{
			"end":   "2018-10-04 08:00:00",
			"start": "2018-09-01 08:00:00",
		}

		statistics := zApi.New{
			APIMethod:    "/v1/statistics/",
			APIUserKey:   "e30e16c201343883f77e",
			APISecretKey: "dbf5606ea4c1f2234201",
			ParamsMap:    paramsMap, //medium priority, if not set ParamsUrlValues
		}

		responceData, err := statistics.Request()

		if err != nil {
			return err
		}

		targetData := jsonTarget{}
		if err := json.Unmarshal(responceData, &targetData); err != nil {
			return err
		}

		return nil
	})

	return g.Wait()
}
```
