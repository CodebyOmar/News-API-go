package newsapi

import (
	"net/http"
	"reflect"
	"net/url"
	"log"
	"strconv"
)

const host = "https://newsapi.org"

type Client struct {
	apikey string
	httpclient *http.Client
}

var httpclient = &http.Client{}

func NewClient(apiKey string) *Client {
	key := apiKey

  return &Client {
  	key,
  	httpclient,
	}
}

func createUrlFromEndpointAndParams (endpoint string, params interface{}) string{
	t := reflect.TypeOf(params)
	v :=  reflect.ValueOf(params)

	baseUrl, err := url.Parse(host+endpoint)
	if err != nil { log.Fatal(err) }
	queryString := baseUrl.Query()

	// Append query strings to URL
	for i := 0; i < t.NumField(); i++ {
		varName := t.Field(i).Name
		varType := t.Field(i).Type
		varIndex := t.Field(i).Index
		varValue := reflect.Indirect(v).FieldByIndex(varIndex)
		varKind := varType.Kind();

		switch  varKind {
		case reflect.Int64:
			strval := strconv.FormatInt(varValue.Int(), 10)
			queryString.Add(varName, strval)
		case reflect.String:
			queryString.Add(varName, varValue.String())
		}
	}

	baseUrl.RawQuery = queryString.Encode()
	return baseUrl.String()
}


