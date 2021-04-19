package zadarma

// "bufio"
// 	"bytes"
// 	"context"
// 	"crypto/tls"
// 	"encoding/base64"
// 	"errors"
// 	"fmt"
// 	"io"
// 	"mime"
// 	"mime/multipart"
// 	"net"
// 	"net/http/httptrace"
// 	"net/textproto"
// 	"net/url"
// 	urlpkg "net/url"
// 	"strconv"
// 	"strings"
// 	"sync"

// 	"golang.org/x/net/idna"
import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
	"unicode/utf8"
)

const (
	linkToApiZadarma string = "https://api.zadarma.com"
)

type API struct {
	HTTPMethod         string
	APIMethod          string
	APIUserKey         string
	APISecretKey       string
	ParamsString       string
	ParamsUrlValues    url.Values
	ParamsMap          map[string]string
	sortedParamsString string
	signature          string
}

func (a *API) request() error {
	req, err := http.NewRequest(
		a.HTTPMethod,
		linkToApiZadarma+a.APIMethod,
		bytes.NewBuffer([]byte(a.ParamsUrlValues.Encode())),
	)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", a.APIUserKey+":"+a.signature)

	client := &http.Client{Timeout: time.Second * 10}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}

	fmt.Printf("%s\n", body)

	return nil
}

func (a *API) createSignature() {
	a.sortedParamsString = ""
	md5Hash := md5.New()
	md5Hash.Write([]byte(a.sortedParamsString))
	md5Str := hex.EncodeToString(md5Hash.Sum(nil))

	data := a.APIMethod + a.sortedParamsString + md5Str

	mac := hmac.New(sha1.New, []byte(a.APISecretKey))
	mac.Write([]byte(data))
	macStr := hex.EncodeToString(mac.Sum(nil))

	a.signature = base64.StdEncoding.EncodeToString([]byte(macStr))
}

func (a *API) sortParamsString() error {
	q, err := url.Parse("?" + a.ParamsString)

	if err != nil {
		return errors.New("error: ParamsString is wrong format! example: \"name=Bob&a=1&b=s&c=x\"")
	}
	// for k, v := range q.Query() {
	// 	fmt.Println(k, v[0])
	// }
	a.ParamsUrlValues = q.Query()
	return nil
}

func prepare_data_to_request(a *API) error {
	var err error = nil
	if a.HTTPMethod == "" {
		a.HTTPMethod = http.MethodGet
	}
	if a.APIMethod == "" {
		err = errors.New("error: APIMethod is empty string! example: \"/v1/info/balance/\"")
	}

	if a.APIUserKey == "" || utf8.RuneCountInString(a.APIUserKey) != 20 {
		err = errors.New("error: APIUserKey is empty string or length not 20 symbols! example: \"e27e28c201943883f77e\"")
	}
	if a.APISecretKey == "" || utf8.RuneCountInString(a.APISecretKey) != 20 {
		err = errors.New("error: APISecretKey is empty string or length not 20 symbols! example: \"e27e28c201943883f77e\"")
	}

	if len(a.ParamsUrlValues) > 0 {
		a.sortedParamsString = a.ParamsUrlValues.Encode()

	} else if len(a.ParamsMap) > 0 {
		for k, v := range a.ParamsMap {
			a.ParamsUrlValues.Set(k, v)
		}

	} else if a.ParamsString != "" {
		urlValues, err := url.ParseQuery(a.ParamsString)
		if err != nil {
			return err
		}
		a.ParamsUrlValues = urlValues
	}

	a.createSignature()
	a.request()

	return err
}

func Request(a API) (string, error) {
	prepare_data_to_request(&a)

	resp, err := http.Get(linkToApiZadarma)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	//fmt.Printf("\n%v\n\n", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		bodyString := string(bodyBytes)
		fmt.Printf("%v", bodyString)
	}

	return "asd", err
}
