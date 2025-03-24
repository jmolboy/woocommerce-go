package woocommerce

import (
	"encoding/json"
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
		Name:             "API调用测试分组",
		Slug:             "iphone",
		Type:             "variable",
		Status:           "publish",
		Description:      `<img src="https://iuufu-pic.oss-cn-qingdao.aliyuncs.com/images/2025-01/b5d3d25e176e43856863460026048138.png"/>`,
		ShortDescription: "这是商品的简述",
		SKU:              "iphone手机壳-黑夜传说2",
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
		Tags: []entity.ProductTag{},
		Images: []entity.ProductImage{
			{
				Src: "https://iuufu-pic.oss-cn-qingdao.aliyuncs.com/images/2025-01/b5d3d25e176e43856863460026048138.png",
			},
			{
				Src: "https://iuufu-pic.oss-cn-qingdao.aliyuncs.com/images/2025-01/b5d3d25e176e43856863460026048138.png",
			},
		},
		Attributes: []entity.ProductAttribute{
			{
				ID:        1,
				Name:      "Color",
				Variation: true,
				Visible:   true,
				Slug:      "pa_color",
				Options: []string{
					"pink", "black", "red", "yellow",
				},
			},
			{
				ID:        2,
				Name:      "Spec",
				Variation: true,
				Slug:      "pa_size",
				Visible:   true,
				Options: []string{
					"iphone15", "iphone16", "iphone17", "iphone18",
				},
			},
		},
		ParentId:          0,
		DefaultAttributes: []entity.ProductDefaultAttribute{},
		MetaData:          []entity.Meta{},
		GroupedProducts:   []int{
			// 2483, 2485, 2490,
		},
	}
	prod, err := wooClient.Services.Product.Create(req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("created product: %+v\n", prod)

	// m.Run()
}

func TestVarProd(m *testing.T) {
	b, err := os.ReadFile("./config/config_me.json")
	if err != nil {
		panic(fmt.Sprintf("Read config error: %s", err.Error()))
	}
	var c config.Config
	err = jsoniter.Unmarshal(b, &c)
	if err != nil {
		panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
	}

	wooClient = NewClient(c)
	prodId := 2777

	req := CreateProductVariationRequest{
		Description:   "测试描述",
		SKU:           "dc6833-5-2-4-2",
		RegularPrice:  2,
		SalePrice:     1,
		Status:        "publish",
		TaxStatus:     "taxable",
		ManageStock:   true,
		StockQuantity: 10,
		StockStatus:   "instock",
		Weight:        1,
		Image: &entity.ProductImage{
			Src: "https://iuufu-pic.oss-cn-qingdao.aliyuncs.com/images/2025-01/b5d3d25e176e43856863460026048138.png",
		},
		Attributes: []entity.ProductVariationAttribute{
			{
				ID:     1,
				Name:   "Color",
				Option: "blue",
			},

			{
				ID:     3,
				Name:   "Spec",
				Option: "iphone13",
			},
		},
		Dimension: &entity.ProductDimension{
			Length: "1cm",
			Width:  "1cm",
			Height: "1cm",
		},
	}
	prod, err := wooClient.Services.ProductVariation.Create(prodId, req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("created product: %+v\n", prod)
}

func TestVarDecode(t *testing.T) {
	jsonRaw := ``
	var prodVar entity.ProductVariation
	err := json.Unmarshal([]byte(jsonRaw), &prodVar)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}

func TestShipZoneMethods(t *testing.T) {
	b, err := os.ReadFile("./config/config_me.json")
	if err != nil {
		panic(fmt.Sprintf("Read config error: %s", err.Error()))
	}
	var c config.Config
	err = jsoniter.Unmarshal(b, &c)
	if err != nil {
		panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
	}

	wooClient = NewClient(c)
	zones, err := wooClient.Services.ShippingZoneMethod.All(2)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	fmt.Printf("zones is:%+v\n", zones)

}

func TestShipZone(m *testing.T) {
	b, err := os.ReadFile("./config/config_me.json")
	if err != nil {
		panic(fmt.Sprintf("Read config error: %s", err.Error()))
	}
	var c config.Config
	err = jsoniter.Unmarshal(b, &c)
	if err != nil {
		panic(fmt.Sprintf("Parse config file error: %s", err.Error()))
	}

	wooClient = NewClient(c)
	zones, err := wooClient.Services.ShippingZone.All()
	if err != nil {
		panic(err)
	}

	for _, v := range zones {
		zoneLocations, err := wooClient.Services.ShippingZoneLocation.All(v.ID)
		if err != nil {
			continue
		}
		fmt.Printf("zone: %+v, locations: %+v\n\n\n", v.Name, zoneLocations)
	}

}
