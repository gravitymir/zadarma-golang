package zadarma

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
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
	Response           []byte
}

func (z *New) prepareData() error {
	if z.APIMethod == "" {
		return errors.New("error: APIMethod is empty! example: \"/v1/info/balance/\"")
	}
	if z.APIUserKey == "" || utf8.RuneCountInString(z.APIUserKey) != 20 {
		return errors.New("error: APIUserKey is empty or length != 20 runes! example: \"e27e28c201943883f77e\"")
	}
	if z.APISecretKey == "" || utf8.RuneCountInString(z.APISecretKey) != 20 {
		return errors.New("error: APISecretKey is empty or length != 20 runes! example: \"e27e28c201943883f77e\"")
	}
	//priority of incoming parameters
	//high ParamsUrlValues
	//medium ParamsMap
	//low ParamsString
	if len(z.ParamsUrlValues) == 0 { //high priority
		if len(z.ParamsMap) > 0 { //medium priority
			for k, v := range z.ParamsMap {
				z.ParamsUrlValues.Set(k, v)
			}

		} else if z.ParamsString != "" { //low priority
			if urlValues, err := url.ParseQuery(z.ParamsString); err != nil {
				return err
			} else {
				z.ParamsUrlValues = urlValues
			}
		}
	}
	//Encode "sorted by key"
	z.SortedParamsString = z.ParamsUrlValues.Encode()
	//https://golang.org/pkg/net/url/#Values.Encode
	return nil
}

func (z *New) getHttpRequest() (*http.Request, error) {
	return http.NewRequestWithContext( // maybe need http.NewRequest()
		context.Background(),
		z.HTTPMethod,
		z.LinkToAPI+z.APIMethod+"?"+z.SortedParamsString, //URL to API
		bytes.NewBuffer([]byte("")),
	)
}

//Request is request to API Zadarma "https://api.zadarma.com"
func (z *New) Request() error {
	//new datas after every new request
	//z.SortedParamsString = ""
	//z.Signature = ""
	//z.ResponseRaw = []byte{}
	if z.HTTPMethod == "" {
		z.HTTPMethod = http.MethodGet
	}
	if z.LinkToAPI == "" {
		z.LinkToAPI = "https://api.zadarma.com"
	}
	if z.Timeout == 0 {
		z.Timeout = 5
	}
	z.Response = []byte{}

	if err := z.prepareData(); err != nil {
		return err
	}

	z.createSignature()

	req, err := z.getHttpRequest()

	if err != nil {
		return err
	}

	req.Header.Set("Authorization", z.APIUserKey+":"+z.Signature)

	//Timeout request
	client := &http.Client{Timeout: time.Second * time.Duration(z.Timeout)}

	r, err := client.Do(req) //request
	if err != nil {
		return err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	z.Response = body

	return nil
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
