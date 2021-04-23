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
	Signature          string
	ResponseRaw        []byte
}

func (z *New) prepareData() error {
	var err error
	if z.HTTPMethod == "" {
		z.HTTPMethod = http.MethodGet
	}
	if z.APIMethod == "" {
		err = errors.New("error: APIMethod is empty string! example: \"/v1/info/balance/\"")
	}

	if z.APIUserKey == "" || utf8.RuneCountInString(z.APIUserKey) != 20 {
		err = errors.New("error: APIUserKey is empty string or length not 20 symbols! example: \"e27e28c201943883f77e\"")
	}
	if z.APISecretKey == "" || utf8.RuneCountInString(z.APISecretKey) != 20 {
		err = errors.New("error: APISecretKey is empty string or length not 20 symbols! example: \"e27e28c201943883f77e\"")
	}

	if len(z.ParamsUrlValues) == 0 {
		if len(z.ParamsMap) > 0 {
			for k, v := range z.ParamsMap {
				z.ParamsUrlValues.Set(k, v)
			}

		} else if z.ParamsString != "" {
			urlValues, err := url.ParseQuery(z.ParamsString)
			if err != nil {
				return err
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

	return err
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
func (z *New) Request() (map[string]interface{}, error) {
	var zMap map[string]interface{}
	if err := z.prepareData(); err != nil {
		return zMap, err
	}

	req, err := z.getHttpRequest()

	if err != nil {
		return zMap, err
	}

	req.Header.Set("Authorization", z.APIUserKey+":"+z.Signature)

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		return zMap, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return zMap, err
	}

	z.ResponseRaw = body

	if err := json.Unmarshal(z.ResponseRaw, &zMap); err != nil {
		return zMap, err
	}

	return zMap, nil
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
