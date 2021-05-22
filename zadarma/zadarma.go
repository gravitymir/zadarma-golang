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
}

//Request is request to API Zadarma "https://api.zadarma.com"
func (z *New) Request(slb *[]byte) error {
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
	if len(z.ParamsUrlValues) == 0 {
		z.ParamsUrlValues = url.Values{}
	}

	if sps, err := prepareData(z); err != nil {
		return err
	} else {
		z.SortedParamsString = sps
	}

	z.Signature = createSignature(z)

	httpRequest, err := getHttpRequest(z)

	if err != nil {
		return err
	}

	httpRequest.Header.Set("Authorization", z.APIUserKey+":"+z.Signature)

	//Timeout request
	client := &http.Client{Timeout: time.Millisecond * time.Duration(z.Timeout)}

	httpResponse, err := client.Do(httpRequest) //request
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	*slb, err = ioutil.ReadAll(httpResponse.Body)

	if err != nil {
		return err
	}
	//https://play.golang.org/p/An7jG5xl2W
	// rv := reflect.ValueOf(responseBody)

	// var sl = struct {
	// 	addr uintptr
	// 	len  int
	// 	cap  int
	// }{rv.Pointer(), rv.Len(), rv.Len()}

	// *b = *(*[]byte)(unsafe.Pointer(&sl))

	return nil
}

func prepareData(z *New) (string, error) {
	if z.APIMethod == "" {
		return "", errors.New("error: APIMethod is empty! example: \"/v1/info/balance/\"")
	}
	if z.APIUserKey == "" || utf8.RuneCountInString(z.APIUserKey) != 20 {
		return "", errors.New("error: APIUserKey is empty or length != 20 runes! example: \"e27e28c201943883f77e\"")
	}
	if z.APISecretKey == "" || utf8.RuneCountInString(z.APISecretKey) != 20 {
		return "", errors.New("error: APISecretKey is empty or length != 20 runes! example: \"e27e28c201943883f77e\"")
	}
	//priority of incoming parameters
	//high ParamsUrlValues
	//medium ParamsMap
	//low ParamsString
	if len(z.ParamsUrlValues) == 0 { //  <---high priority
		if len(z.ParamsMap) > 0 { //  <---medium priority
			for k, v := range z.ParamsMap {
				z.ParamsUrlValues.Set(k, v)
			}
		} else if z.ParamsString != "" { //  <---low priority
			if urlValues, err := url.ParseQuery(z.ParamsString); err != nil {
				return "", err
			} else {
				z.ParamsUrlValues = urlValues
			}
		}
	}
	//Encode "sorted by key"
	//https://golang.org/pkg/net/url/#Values.Encode
	return z.ParamsUrlValues.Encode(), nil
}

func getHttpRequest(z *New) (*http.Request, error) {
	return http.NewRequestWithContext( // maybe need http.NewRequest()
		context.Background(),
		z.HTTPMethod,
		z.LinkToAPI+z.APIMethod+"?"+z.SortedParamsString, //URL to API
		bytes.NewBuffer([]byte("")),
	)
}

func createSignature(z *New) string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(z.SortedParamsString))
	md5Str := hex.EncodeToString(md5Hash.Sum(nil))

	data := z.APIMethod + z.SortedParamsString + md5Str

	mac := hmac.New(sha1.New, []byte(z.APISecretKey))
	mac.Write([]byte(data))
	macStr := hex.EncodeToString(mac.Sum(nil))

	return base64.StdEncoding.EncodeToString([]byte(macStr))
}
