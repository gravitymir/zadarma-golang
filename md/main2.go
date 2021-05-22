package main

import (
	"encoding/json"
	"fmt"

	zApi "github.com/gravitymir/zadarma-golang/zadarma"
)

func main() {
	if err := statistics(); err != nil {
		fmt.Printf("%v\n", err)
	}
	if err := statistics(); err != nil {
		fmt.Printf("%v\n", err)
	}

}
func prettyPrint(data interface{}) {
	dataJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("\n<---\n %s\n--->", string(dataJSON))
}

func statistics() error {

	type CatchStatistics struct {
		Status string `json:"status"`
		Start  string `json:"start"`
		End    string `json:"end"`
		Stats  []struct {
			Id          string  `json:"id"`
			Sip         string  `json:"sip"`
			Callstart   string  `json:"callstart"`
			From        int     `json:"from"`
			To          int     `json:"to"`
			Description string  `json:"description"`
			Disposition string  `json:"disposition"`
			Billseconds int     `json:"billseconds"`
			Cost        float32 `json:"cost"`
			Billcost    float32 `json:"billcost"`
			Currency    string  `json:"currency"`
			Seconds     int     `json:"seconds"`
		}
		Message string `json:"message"` //if status == error
	}

	statistics := zApi.New{
		APIMethod:    "/v1/statistics/",
		APIUserKey:   "e30e16c201343883f77e",
		APISecretKey: "dbf5606ea4c1f2234201",
		ParamsMap: map[string]string{
			"end":   "2018-10-04 08:00:00",
			"start": "2018-09-01 08:00:00",
		},
		Timeout: 400, //Miliseconds
	}

	data := []byte{}

	if err := statistics.Request(&data); err != nil {
		return err
	}

	prettyPrint(statistics)

	catchDataToStruct := zApi.CatchStatistics{}

	if err := json.Unmarshal(data, &catchDataToStruct); err != nil {
		return err
	}

	//fmt.Println(catchDataToStruct)

	return nil
}
