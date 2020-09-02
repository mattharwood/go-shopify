package goshopify

import (
	"encoding/json"
	"fmt"
)

func GetMetafields() {
	url := urltpl() + "metafields.json"
	m := make(map[string]string)
	m["limit"] = "2"
	response := queryGet(url, m)
	var products ProductsResponse
	json.Unmarshal(response, &products)
	for _, v := range products.Products {
		fmt.Println(v.Handle)
	}
}

func GetProductMetafield(handle, mfnamspace, mfkey string) {
	url := urltpl() + "products/" + handle + "/metafields.json"
	m := make(map[string]string)
	response := queryGet(url, m)
	var products ProductsResponse
	json.Unmarshal(response, &products)
	for _, v := range products.Products {
		fmt.Println(v.Handle)
	}
}

func SetProductMetafield(handle, namespace, key, value, valuetype string) {
	url := urltpl() + "products/" + handle + "/metafields.json"
	body := `
	{
		"metafield": {
		  "namespace": ` + namespace + `,
		  "key": ` + key + `,
		  "value": ` + value + `,
		  "value_type": ` + valuetype + `
		}
	  }
	  
	`
	queryPost(url, body)
}

type MetafieldResponse struct {
	Metafields []struct {
		AdminGraphqlAPIID string `json:"admin_graphql_api_id"`
		CreatedAt         string `json:"created_at"`
		Description       string `json:"description"`
		ID                int64  `json:"id"`
		Key               string `json:"key"`
		Namespace         string `json:"namespace"`
		OwnerID           int64  `json:"owner_id"`
		OwnerResource     string `json:"owner_resource"`
		UpdatedAt         string `json:"updated_at"`
		Value             string `json:"value"`
		ValueType         string `json:"value_type"`
	} `json:"metafields"`
}

type Metafield struct {
	Namespace string
	Key       string
	Value     string
	ValueType string
}
