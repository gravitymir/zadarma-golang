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
	Error              error
}

func (z *New) prepareData() {
	if z.APIMethod == "" {
		z.Error = errors.New("error: APIMethod is empty! example: \"/v1/info/balance/\"")
	}
	if z.APIUserKey == "" || utf8.RuneCountInString(z.APIUserKey) != 20 {
		z.Error = errors.New("error: APIUserKey is empty or length != 20 runes! example: \"e27e28c201943883f77e\"")
	}
	if z.APISecretKey == "" || utf8.RuneCountInString(z.APISecretKey) != 20 {
		z.Error = errors.New("error: APISecretKey is empty or length != 20 runes! example: \"e27e28c201943883f77e\"")
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
			urlValues, err := url.ParseQuery(z.ParamsString)
			if err != nil {
				z.Error = err
				return
			}
			z.ParamsUrlValues = urlValues
		}
	}
	//Encode "sorted by key"
	z.SortedParamsString = z.ParamsUrlValues.Encode()
	//https://golang.org/pkg/net/url/#Values.Encode
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
func (z *New) Request() {
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
	z.Error = nil
	z.Response = []byte{}

	z.prepareData()
	z.createSignature()

	req, err := z.getHttpRequest()

	if err != nil {
		z.Error = err
	}

	req.Header.Set("Authorization", z.APIUserKey+":"+z.Signature)

	//Timeout request
	client := &http.Client{Timeout: time.Second * time.Duration(z.Timeout)}

	r, err := client.Do(req) //request
	if err != nil {
		z.Error = err
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		z.Error = err
	}
	z.Response = body
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
