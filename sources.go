package newsapi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Sources
func (c *Client) Sources(params interface{}) (response map[string]interface{}, err error) {
	endpoint := "/v2/sources"
	url := createUrlFromEndpointAndParams(endpoint, params)
	data, fetchErr := getDataFromWeb(url, c)
  return data, fetchErr
}
