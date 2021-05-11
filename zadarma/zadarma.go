package zadarma

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"unicode/utf8"
)

//New is main struct
type New struct {
	APIUserKey         string
	APISecretKey       string
	HTTPMethod         string
	APIMethod          string
	LinkToAPI          string
	ParamsUrlValues    url.Values
	ParamsMap          map[string]string
	ParamsString       string
	SortedParamsString string
	Timeout            uint64
	Signature          string
	ResponseRaw        []byte
	Error              error
	Status             bool
	Result             map[string]interface{}
}

func (z *New) prepareData() {
	if z.HTTPMethod == "" {
		z.HTTPMethod = http.MethodGet
	}
	if z.APIMethod == "" {
		z.Error = errors.New("error: APIMethod is empty string! example: \"/v1/info/balance/\"")
	}

	if z.APIUserKey == "" || utf8.RuneCountInString(z.APIUserKey) != 20 {
		z.Error = errors.New("error: APIUserKey is empty string or length not 20 symbols! example: \"e27e28c201943883f77e\"")
	}
	if z.APISecretKey == "" || utf8.RuneCountInString(z.APISecretKey) != 20 {
		z.Error = errors.New("error: APISecretKey is empty string or length not 20 symbols! example: \"e27e28c201943883f77e\"")
	}

	//1 priority ParamsUrlValues high
	//2 priority ParamsMap medium
	//3 priority ParamsString low
	if len(z.ParamsUrlValues) == 0 { //high priority
		if len(z.ParamsMap) > 0 { //medium priority
			for k, v := range z.ParamsMap {
				z.ParamsUrlValues.Set(k, v)
			}

		} else if z.ParamsString != "" { //low priority
			urlValues, err := url.ParseQuery(z.ParamsString)
			if err != nil {
				z.Error = err
				return
			}
			z.ParamsUrlValues = urlValues
		}
	}

	z.SortedParamsString = z.ParamsUrlValues.Encode() //Encode "sorted by key"
	//https://golang.org/pkg/net/url/#Values.Encode

	if z.LinkToAPI == "" {
		z.LinkToAPI = "https://api.zadarma.com"
	}

	z.createSignature()

}

func (z *New) getHttpRequest() (*http.Request, error) {
	return http.NewRequestWithContext(
		context.Background(),
		z.HTTPMethod,
		z.LinkToAPI+z.APIMethod+"?"+z.SortedParamsString,
		bytes.NewBuffer([]byte("")),
	)
}

//Request is request to API Zadarma "https://api.zadarma.com"
func (z *New) Request() *New {
	z.SortedParamsString = ""
	if z.Timeout == 0 {
		z.Timeout = 5
	}
	z.Signature = ""
	z.ResponseRaw = []byte{}
	z.Error = nil
	z.Status = false
	z.Result = map[string]interface{}{}

	z.prepareData()

	if z.Error != nil {
		return z
	}

	req, err := z.getHttpRequest()

	if err != nil {
		z.Error = err
		return z
	}

	req.Header.Set("Authorization", z.APIUserKey+":"+z.Signature)

	client := &http.Client{Timeout: time.Second * time.Duration(z.Timeout)}

	resp, err := client.Do(req)
	if err != nil {
		z.Error = err
		return z
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		z.Error = err
		return z
	}

	z.ResponseRaw = body

	if err := json.Unmarshal(z.ResponseRaw, &z.Result); err != nil {
		z.Error = err
		return z
	}

	if z.Result["status"] == "success" {
		delete(z.Result, "status")
		z.Status = true
	} else if z.Result["status"] == "error" {

		z.Error = errors.New("error: " + z.Result["message"].(string))
		delete(z.Result, "status")
		delete(z.Result, "message")
	}

	return z
}

func (z *New) createSignature() {
	md5Hash := md5.New()
	md5Hash.Write([]byte(z.SortedParamsString))
	md5Str := hex.EncodeToString(md5Hash.Sum(nil))

	data := z.APIMethod + z.SortedParamsString + md5Str

	mac := hmac.New(sha1.New, []byte(z.APISecretKey))
	mac.Write([]byte(data))
	macStr := hex.EncodeToString(mac.Sum(nil))

	z.Signature = base64.StdEncoding.EncodeToString([]byte(macStr))
}
