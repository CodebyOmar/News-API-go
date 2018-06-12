package newsapi

import (
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

const endpoint  = "/v2/top-headlines"

// Fetch Top headlines
func (c *Client) TopHeadlines(params interface{}, callback func(response map[string]interface{}, err error)) {

		url := createUrlFromEndpointAndParams(endpoint, params)

		if req, err:= http.NewRequest("GET", url, nil); err == nil {
			req.Header.Set("X-Api-Key", c.apikey)

      response, getErr := c.httpclient.Do(req)

			var responseObj map[string]interface{}

			stringified, _ := ioutil.ReadAll(response.Body)

			jsonErr := json.Unmarshal([]byte(stringified), &responseObj)
			if jsonErr != nil {
				panic(jsonErr)
			}

      callback(responseObj, getErr)
		}else {
	  	log.Fatal(err)
		}

}