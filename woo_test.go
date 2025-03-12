package woocommerce

import (
	"fmt"
	"os"
	"testing"

	"github.com/jmolboy/woocommerce-go/config"
	"github.com/jmolboy/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

var wooClient *WooCommerce

var orderId, noteId int
var mainId, childId int

// Operate data use WooCommerce for golang
func Example() {
	b, err := os.ReadFile("./config/config_test.json")
	if err != nil {
		panic(fmt.Sprintf("Read config error: %s", err.Error()))
	}
	var c config.Config
	err = jsoniter.Unmarshal(b, &c)
	if err != nil {
		panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
	}

	wooClient = NewClient(c)
	// Query an order
	order, err := wooClient.Services.Order.One(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fmt.Sprintf("%#v", order))
	}

	// Query orders
	params := OrdersQueryParams{
		After: "2022-06-10",
	}
	params.PerPage = 100
	for {
		orders, total, totalPages, isLastPage, err := wooClient.Services.Order.All(params)
		if err != nil {
			break
		}
		fmt.Println(fmt.Sprintf("Page %d/%d", total, totalPages))
		// read orders
		for _, order := range orders {
			_ = order
		}
		if err != nil || isLastPage {
			break
		}
		params.Page++
	}
}

func ExampleErrorWrap() {
	err := ErrorWrap(200, "Ok")
	if err != nil {
		return
	}
}

func getOrderId(t *testing.T) {
	t.Log("Execute getOrderId test")
	items, _, _, _, err := wooClient.Services.Order.All(OrdersQueryParams{})
	if err != nil || len(items) == 0 {
		t.FailNow()
	}
	orderId = items[0].ID
	mainId = items[0].ID
}

func TestCreateProd(m *testing.T) {
	b, err := os.ReadFile("./config/config_test.json")
	if err != nil {
		panic(fmt.Sprintf("Read config error: %s", err.Error()))
	}
	var c config.Config
	err = jsoniter.Unmarshal(b, &c)
	if err != nil {
		panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
	}

	wooClient = NewClient(c)

	req := CreateProductRequest{
		Name:             "API调用测试测试图片",
		Slug:             "这是一个slug",
		Type:             "simple",
		Status:           "publish",
		Description:      `测试描述`,
		ShortDescription: "这是商品的简述",
		SKU:              "iphone12Plus手机壳",
		RegularPrice:     88,
		SalePrice:        58,
		TaxStatus:        "none",
		StockQuantity:    10,
		StockStatus:      "instock",
		Weight:           "10g",
		Categories: []entity.ProductCategory{
			{
				ID:   1383,
				Name: "Bags &amp; Backpacks",
			},
		},
		Tags:   []entity.ProductTag{},
		Images: []entity.ProductImage{},
		Attributes: []entity.ProductAttribute{
			{
				ID:   1,
				Name: "Coloe",
				Slug: "pa_color",
				Type: "select",
			},
			{
				ID:   2,
				Name: "Size",
				Slug: "pa_size",
				Type: "size",
			},
		},
		ParentId:          2483,
		DefaultAttributes: []entity.ProductDefaultAttribute{},
		MetaData:          []entity.Meta{},
	}
	prod, err := wooClient.Services.Product.Create(req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("created product: %+v\n", prod)

	// m.Run()
}
