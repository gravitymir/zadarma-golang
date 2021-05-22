package zadarma

import (
	"encoding/json"
	"errors"
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

	data := []byte{}

	if err := number_lookup.Request(&data); err != nil {
		t.Error(err)
	}

	jt := jsonTarget{}

	if err := json.Unmarshal(data, &jt); err != nil {
		t.Error(err)
	}
	if jt.Status != "success" && jt.Status != "error" {
		t.Error(errors.New("Error jt.Status != success, error"))
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

	data := []byte{}

	if err := number_lookup.Request(&data); err != nil {
		t.Error(err)
	}
	jt := jsonTarget{}
	if err := json.Unmarshal(data, &jt); err != nil {
		t.Error(err)
	}
	if jt.Status != "success" && jt.Status != "error" {
		t.Error(errors.New("Error jt.Status != success, error"))
	}
}

func TestParamsMap(t *testing.T) {

	paramsMap := map[string]string{
		"end":   "2018-10-04 08:00:00",
		"start": "2018-09-01 08:00:00",
	}

	statistics := New{
		APIMethod:    "/v1/statistics/",
		APIUserKey:   "e30e16c201343883f77e",
		APISecretKey: "dbf5606ea4c1f2234201",
		ParamsMap:    paramsMap, //medium priority, if not set ParamsUrlValues
	}

	data := []byte{}

	if err := statistics.Request(&data); err != nil {
		t.Error(err)
	}
	if statistics.Signature != "ZmI3YzA1MmJkNzQyYmE1YzI3OWY3OTRjNTlmZDFkMTE2MmI4YmI0Yg==" {
		t.Error(errors.New("Error wrong statistics.Signature"))
	}
	jt := jsonTarget{}
	if err := json.Unmarshal(data, &jt); err != nil {
		t.Error(err)
	}
	if jt.Status != "success" && jt.Status != "error" {
		t.Error(errors.New("Error jt.Status != success, error"))
	}
}
