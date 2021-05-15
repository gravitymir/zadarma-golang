package zadarma

import (
	"encoding/json"
	"testing"
)

const (
	UserKey   = "e30e16c201343883f77e"
	SecretKey = "dbf5606ea4c1f2234201"
)

type jsonTarget struct {
	Status string `json:"status"`
}

func TestRequest(t *testing.T) {
	number_lookup := New{
		APIMethod:    "/v1/info/balance/",
		APIUserKey:   UserKey,
		APISecretKey: SecretKey,
	}

	responceData, err := number_lookup.Request()
	if err != nil {
		t.Error(err)
	}

	targetData := jsonTarget{}
	if err := json.Unmarshal(responceData, &targetData); err != nil {
		t.Error(err)
	}
	if targetData.Status != "success" && targetData.Status != "error" {
		t.Error(err)
	}
}

func TestParamsString(t *testing.T) {
	paramsString := "start=2018-09-01 08:00:00&end=2018-10-04 08:00:00"

	number_lookup := New{
		APIMethod:    "/v1/statistics/",
		APIUserKey:   UserKey,
		APISecretKey: SecretKey,
		ParamsString: paramsString,
	}

	responceData, err := number_lookup.Request()
	if err != nil {
		t.Error(err)
	}
	targetData := jsonTarget{}
	if err := json.Unmarshal(responceData, &targetData); err != nil {
		t.Error(err)
	}
	if targetData.Status != "success" && targetData.Status != "error" {
		t.Error(err)
	}
}

func TestParamsMap(t *testing.T) {
	// paramsUrlValues := url.Values{
	// 	"start": []string{"2018-09-01 08:00:00"},
	// 	"end":   []string{"2018-10-04 08:00:00"},
	// }

	paramsMap := map[string]string{
		"end":   "2018-10-04 08:00:00",
		"start": "2018-09-01 08:00:00",
	}

	//paramsString := "start=2018-09-01 08:00:00&end=2018-10-04 08:00:00"

	statistics := New{
		APIMethod:    "/v1/statistics/",
		APIUserKey:   "e30e16c201343883f77e",
		APISecretKey: "dbf5606ea4c1f2234201",
		ParamsMap:    paramsMap, //medium priority, if not set ParamsUrlValues
	}

	responceData, err := statistics.Request()

	if err != nil {
		t.Error(err)
	}
	if statistics.Signature != "ZmI3YzA1MmJkNzQyYmE1YzI3OWY3OTRjNTlmZDFkMTE2MmI4YmI0Yg==" {
		t.Error(err)
	}
	targetData := jsonTarget{}
	if err := json.Unmarshal(responceData, &targetData); err != nil {
		t.Error(err)
	}
	if targetData.Status != "success" && targetData.Status != "error" {
		t.Error(err)
	}
}
