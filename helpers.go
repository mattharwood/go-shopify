package goshopify

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func SetConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // path to look for the config file in
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

}

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
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
}
