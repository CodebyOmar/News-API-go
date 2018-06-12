# News API SDK for Go

Up-to-date news headlines and metadata in JSON from 70+ popular news sites. Powered by NewsAPI.org.

All you need is an API key from [https://newsapi.org](https://newsapi.org).

This library is built to provide convinience in communicating with the NewsAPI, for detailed documentation on
required parameters, check thier [documentation](https://newsapi.org/docs)

## Example Usage
```go
package main

import (
    "github.com/codebyomar/News-API-go"
    "fmt"
)

type params struct {
    country string
    pageSize int64
}

func main() {
    c := newsapi.NewClient("YOUR_API_KEY_HERE")

    p := params {
      country: "ng",
      pageSize: 40,
    }

    c.TopHeadlines(p, func (res map[string]interface{}, err error) {

      fmt.Println("response",res)
      fmt.Println("errr", err)
    })
}
```

## TODO
1. include endpoints from V1 of NewsAPI
2. write Test

