package newsapi

// Fetch Top headlines
func (c *Client) TopHeadlines(params interface{}) (map[string]interface{}, error) {
  endpoint := "/v2/top-headlines"
	url := createUrlFromEndpointAndParams(endpoint, params)
	data, fetchErr := getDataFromWeb(url, c)
  return data, fetchErr
}
