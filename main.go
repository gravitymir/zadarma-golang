package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	zApi "github.com/gravitymir/zadarma-golang/zadarma"
)

func main() {

	if err := infoBalance(); err != nil {
		fmt.Printf("%v\n", err)
	}
	if err := infoPrice(); err != nil {
		fmt.Printf("%v\n", err)
	}
	if err := statistics(); err != nil {
		fmt.Printf("%v\n", err)
	}
	if err := smsSend(); err != nil {
		fmt.Printf("%v\n", err)
	}

}
func prettyPrint(data interface{}) {
	dataJSON, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("\n%s\n", string(dataJSON))
}

func infoBalance() error {

	balance := zApi.New{
		APIMethod:    "/v1/info/balance/",
		APIUserKey:   "e30e16c201343883f77e",
		APISecretKey: "dbf5606ea4c1f2234201",
	}

	data := []byte{}

	if err := balance.Request(&data); err != nil {
		return err
	}

	catchDataToStruct := struct {
		Status   string  `json:"status"`
		Balance  float32 `json:"balance"`
		Currency string  `json:"currency"`
		Message  string  `json:"message"`
	}{}
	//or catchDataToStruct := zApi.CatchInfoBalance{},  if zApi.(struct) implement

	if err := json.Unmarshal(data, &catchDataToStruct); err != nil {
		return err
	}

	prettyPrint(catchDataToStruct)

	return nil
}

func infoPrice() error {

	priceKazakhstan := zApi.New{
		APIMethod:    "/v1/info/price",
		APIUserKey:   "e30e16c201343883f77e",
		APISecretKey: "dbf5606ea4c1f2234201",
		ParamsString: "number=77270000000",
	}

	dataKaz := []byte{}

	if err := priceKazakhstan.Request(&dataKaz); err != nil {
		return err
	}

	kaz := zApi.CatchInfoPrice{}

	if err := json.Unmarshal(dataKaz, &kaz); err != nil {
		return err
	}
	prettyPrint(kaz)
	//---------------------------------------------------------------

	priceUnitedKingdom := zApi.New{
		APIMethod:    "/v1/info/price",
		APIUserKey:   "e30e16c201343883f77e",
		APISecretKey: "dbf5606ea4c1f2234201",
		ParamsString: "number=380440000000", ///Ukraine Kyiv
		ParamsMap: map[string]string{ //higher priority above the ParamsString
			"number": "442040000000", //United Kingdom London
		},
	}

	dataKingdom := []byte{}

	if err := priceUnitedKingdom.Request(&dataKingdom); err != nil {
		return err
	}

	king := zApi.CatchInfoPrice{} //if zApi.(struct) implement
	if err := json.Unmarshal(dataKingdom, &king); err != nil {
		return err
	}
	prettyPrint(king)
	//---------------------------------------------------------------

	//Params priority
	//higher priority ParamsUrlValues
	//medium priority ParamsMap
	//low priority ParamsString

	//Kyiv or RusFed or London
	priceRusFed := zApi.New{
		APIMethod:    "/v1/info/price",
		APIUserKey:   "e30e16c201343883f77e",
		APISecretKey: "dbf5606ea4c1f2234201",
		ParamsString: "number=380440000000", ///Ukraine Kyiv
		ParamsUrlValues: url.Values{ //higher priority above the ParamsMap
			"number": []string{"74952409106"}, //Moscow
		},
		ParamsMap: map[string]string{ //higher priority above the ParamsString
			"number": "442040000000", //United Kingdom London
		},
	}
	dataRusFed := []byte{}

	if err := priceRusFed.Request(&dataRusFed); err != nil {
		return err
	}

	resp := zApi.CatchInfoPrice{}
	if err := json.Unmarshal(dataRusFed, &resp); err != nil {
		return err
	}
	prettyPrint(resp)

	priceRusFed.ParamsUrlValues = url.Values{ //change number
		"number": []string{"73952407198"}, //Irkutsk
	}
	if err := priceRusFed.Request(&dataRusFed); err != nil { //try again
		return err
	}
	resp = zApi.CatchInfoPrice{}
	if err := json.Unmarshal(dataRusFed, &resp); err != nil {
		return err
	}
	prettyPrint(resp)

	priceRusFed.ParamsUrlValues = url.Values{ //change number
		"number": []string{"78612179165"}, ///Krasnodar
	}
	if err := priceRusFed.Request(&dataRusFed); err != nil { //try again
		return err
	}
	resp = zApi.CatchInfoPrice{} //if zApi.(struct) implement
	if err := json.Unmarshal(dataRusFed, &resp); err != nil {
		return err
	}
	prettyPrint(resp)

	//---------------------------------------------------------------

	return nil
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
		ParamsUrlValues: url.Values{
			"start": []string{"2018-09-01 08:00:00"},
			"end":   []string{"2018-10-04 08:00:00"},
		},
		Timeout: 10000, //Miliseconds
	}

	data := []byte{}

	if err := statistics.Request(&data); err != nil {
		return err
	}

	catchDataToStruct := zApi.CatchStatistics{}
	//or
	//catchDataToStruct := CatchStatistics{}, //if zApi.(struct) not implement

	if err := json.Unmarshal(data, &catchDataToStruct); err != nil {
		return err
	}

	prettyPrint(catchDataToStruct)

	return nil
}

func infoTimezone() error {

	timeZone := zApi.New{
		APIMethod:    "/v1/info/timezone/",
		APIUserKey:   "e30e16c201343883f77e",
		APISecretKey: "dbf5606ea4c1f2234201",
	}

	data := []byte{}

	if err := timeZone.Request(&data); err != nil {
		return err
	}
	fmt.Printf("%+v", string(data))
	return nil
}

//HTTPMethod Post
func smsSend() error {

	sms := zApi.New{
		HTTPMethod:   http.MethodPost, // or "POST"
		APIMethod:    "/v1/sms/send/",
		APIUserKey:   "e30e16c201343883f77e",
		APISecretKey: "dbf5606ea4c1f2234201",
		//params in map
		ParamsMap: map[string]string{
			"number":  "67200000000",
			"message": "Text сообщения\nперенос строки\n1234567890",
		},
		//params in string
		ParamsString: "number=67200000000&message=Text сообщения\nперенос строки\n1234567890",
		//params in url.Values [recommend]
		ParamsUrlValues: url.Values{
			"number": []string{
				"67200000000", //recipient's phone number
			},
			"message": []string{
				"Text сообщения\nперенос строки\n1234567890",
			},
			//"caller_id": []string{"74950000000"}, //[optional]
		},
	}

	data := []byte{}

	if err := sms.Request(&data); err != nil {
		return err
	}

	result := zApi.CatchSmsSend{} //if zApi.CatchStatistics implement
	//or
	result = struct {
		Status   string  `json:"status"`
		Messages int     `json:"messages"`
		Cost     float32 `json:"cost"`
		Currency string  `json:"currency"`
		Message  string  `json:"message"`
	}{}

	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}

	prettyPrint(result)

	return nil
}
