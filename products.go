package goshopify

import (
	"encoding/json"
	"fmt"
	"time"
)

// GetProducts Returns all products in the Shopify account
func GetProducts() {
	url := urltpl() + "products.json"
	m := make(map[string]string)
	m["limit"] = "2"
	response := queryGet(url, m)
	var products ProductsResponse
	json.Unmarshal(response, &products)
	for _, v := range products.Products {
		fmt.Println(v.Handle)
	}
}

type ProductsResponse struct {
	Products []struct {
		ID                int64     `json:"id"`
		Title             string    `json:"title"`
		BodyHTML          string    `json:"body_html"`
		Vendor            string    `json:"vendor"`
		ProductType       string    `json:"product_type"`
		CreatedAt         time.Time `json:"created_at"`
		Handle            string    `json:"handle"`
		UpdatedAt         time.Time `json:"updated_at"`
		PublishedAt       time.Time `json:"published_at"`
		TemplateSuffix    string    `json:"template_suffix"`
		PublishedScope    string    `json:"published_scope"`
		Tags              string    `json:"tags"`
		AdminGraphqlAPIID string    `json:"admin_graphql_api_id"`
		Variants          []struct {
			ID                   int64       `json:"id"`
			ProductID            int64       `json:"product_id"`
			Title                string      `json:"title"`
			Price                string      `json:"price"`
			Sku                  string      `json:"sku"`
			Position             int         `json:"position"`
			InventoryPolicy      string      `json:"inventory_policy"`
			CompareAtPrice       interface{} `json:"compare_at_price"`
			FulfillmentService   string      `json:"fulfillment_service"`
			InventoryManagement  string      `json:"inventory_management"`
			Option1              string      `json:"option1"`
			Option2              interface{} `json:"option2"`
			Option3              interface{} `json:"option3"`
			CreatedAt            time.Time   `json:"created_at"`
			UpdatedAt            time.Time   `json:"updated_at"`
			Taxable              bool        `json:"taxable"`
			Barcode              string      `json:"barcode"`
			Grams                int         `json:"grams"`
			ImageID              interface{} `json:"image_id"`
			Weight               float64     `json:"weight"`
			WeightUnit           string      `json:"weight_unit"`
			InventoryItemID      int64       `json:"inventory_item_id"`
			InventoryQuantity    int         `json:"inventory_quantity"`
			OldInventoryQuantity int         `json:"old_inventory_quantity"`
			RequiresShipping     bool        `json:"requires_shipping"`
			AdminGraphqlAPIID    string      `json:"admin_graphql_api_id"`
		} `json:"variants"`
		Options []struct {
			ID        int64    `json:"id"`
			ProductID int64    `json:"product_id"`
			Name      string   `json:"name"`
			Position  int      `json:"position"`
			Values    []string `json:"values"`
		} `json:"options"`
		Images []struct {
			ID                int64         `json:"id"`
			ProductID         int64         `json:"product_id"`
			Position          int           `json:"position"`
			CreatedAt         time.Time     `json:"created_at"`
			UpdatedAt         time.Time     `json:"updated_at"`
			Alt               interface{}   `json:"alt"`
			Width             int           `json:"width"`
			Height            int           `json:"height"`
			Src               string        `json:"src"`
			VariantIds        []interface{} `json:"variant_ids"`
			AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
		} `json:"images"`
		Image struct {
			ID                int64         `json:"id"`
			ProductID         int64         `json:"product_id"`
			Position          int           `json:"position"`
			CreatedAt         time.Time     `json:"created_at"`
			UpdatedAt         time.Time     `json:"updated_at"`
			Alt               interface{}   `json:"alt"`
			Width             int           `json:"width"`
			Height            int           `json:"height"`
			Src               string        `json:"src"`
			VariantIds        []interface{} `json:"variant_ids"`
			AdminGraphqlAPIID string        `json:"admin_graphql_api_id"`
		} `json:"image"`
	} `json:"products"`
}
