package goshopify

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func urltpl() string {
	return "https://" + viper.GetString("api-key") + ":" + viper.GetString("pass") + "@" + viper.GetString("api-url")
}

func queryGet(url string, params map[string]string) []byte {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	// ...
	if len(params) > 0 {
		for k, v := range params {
			req.Header.Add(k, v)

		}
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return bodyBytes
}

func queryPost(url, body string) {

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(body))

}
