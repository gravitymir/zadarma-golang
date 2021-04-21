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
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
	"unicode/utf8"
)

type ResponceJSON struct {
	Status   string  `json:"status"`
	Balance  float64 `json:"balance"`
	Currency string  `json:"currency"`
}

type New struct {
	LinkToAPI          string
	HTTPMethod         string
	APIMethod          string
	APIUserKey         string
	APISecretKey       string
	ParamsUrlValues    url.Values
	ParamsMap          map[string]string
	ParamsString       string
	SortedParamsString string
	Signature          string
	ResponseString     string
	ResponseStruct     ResponceJSON
}

func (z *New) prepare() error {
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

	fmt.Printf("%v", z.SortedParamsString)
	if z.LinkToAPI == "" {
		z.LinkToAPI = "https://api.zadarma.com"
	}

	z.createSignature()

	// if z.HTTPMethod == http.MethodGet {
	// 	z.APIMethod = z.APIMethod + "?" + z.SortedParamsString
	// 	z.SortedParamsString = ""
	// }

	return err
}

func (z *New) request() (*http.Request, error) {
	return http.NewRequestWithContext(
		context.Background(),
		z.HTTPMethod,
		z.LinkToAPI+z.APIMethod+"?"+z.SortedParamsString,
		bytes.NewBuffer([]byte("")),
	)
}

//Request to API var z.LinkToAPI https://api.zadarma.com
func (z *New) Go() error {

	if err := z.prepare(); err != nil {
		return err
	}

	req, err := z.request()

	if err != nil {
		return err
	}

	req.Header.Set("Authorization", z.APIUserKey+":"+z.Signature)

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	z.ResponseString = string(body)
	fmt.Printf("%v", z.ResponseString)

	// var inventory ResponceJSON
	// if err := json.Unmarshal([]byte(body), &inventory); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%v", inventory)

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

// func Request(a Zadarma) (string, error) {
// 	prepare_data_to_request(&a)

// 	resp, err := http.Get(z.LinkToAPI)
// 	if err != nil {
// 		return "", err
// 	}
// 	defer resp.Body.Close()

// 	//fmt.Printf("\n%v\n\n", resp.StatusCode)

// 	if resp.StatusCode != http.StatusOK {
// 		bodyBytes, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		bodyString := string(bodyBytes)
// 		fmt.Printf("%v", bodyString)
// 	}

// 	return "asd", err
// }

// func (a *Zadarma) sortParamsString() error {
// 	q, err := url.Parse("?" + a.ParamsString)

// 	if err != nil {
// 		return errors.New("error: ParamsString is wrong format! example: \"name=Bob&a=1&b=s&c=x\"")
// 	}
// 	// for k, v := range q.Query() {
// 	// 	fmt.Println(k, v[0])
// 	// }
// 	a.ParamsUrlValues = q.Query()
// 	return nil
// }
