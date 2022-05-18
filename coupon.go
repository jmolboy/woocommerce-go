package woocommerce

import (
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/google/go-querystring/query"
	"github.com/hiscaler/woocommerce-go/entity"
	jsoniter "github.com/json-iterator/go"
)

// https://woocommerce.github.io/woocommerce-rest-api-docs/?php#coupon-properties

type couponService service

type Coupon struct {
	ID                        string            `json:"id"`
	Code                      string            `json:"code"`
	Amount                    string            `json:"amount"`
	DateCreated               string            `json:"date_created"`
	DateCreatedGMT            string            `json:"date_created_gmt"`
	DateModified              string            `json:"date_modified"`
	DateModifiedGMT           string            `json:"date_modified_gmt"`
	DiscountType              string            `json:"discount_type"`
	Description               string            `json:"description"`
	DateExpires               string            `json:"date_expires"`
	DateExpiresGMT            string            `json:"date_expires_gmt"`
	UsageCount                int               `json:"usage_count"`
	IndividualUse             bool              `json:"individual_use"`
	ProductIDs                []int             `json:"product_ids"`
	ExcludedProductIDs        []int             `json:"excluded_product_ids"`
	UsageLimit                int               `json:"usage_limit"`
	UsageLimitPerUser         int               `json:"usage_limit_per_user"`
	LimitUsageToXItems        int               `json:"limit_usage_to_x_items"`
	FreeShipping              bool              `json:"free_shipping"`
	ProductCategories         []int             `json:"product_categories"`
	ExcludedProductCategories []int             `json:"excluded_product_categories"`
	ExcludeSaleItems          bool              `json:"exclude_sale_items"`
	MinimumAmount             string            `json:"minimum_amount"`
	MaximumAmount             string            `json:"maximum_amount"`
	EmailRestrictions         []string          `json:"email_restrictions"`
	UsedBy                    []int             `json:"used_by"`
	MetaData                  []entity.MetaData `json:"meta_data"`
}

type CouponsQueryParams struct {
	QueryParams
	Search  string   `url:"search,omitempty"`
	After   string   `url:"after,omitempty"`
	Before  string   `url:"before,omitempty"`
	Exclude []string `url:"exclude,omitempty"`
	Include []string `url:"include,omitempty"`
	Code    string   `url:"code,omitempty"`
}

func (m CouponsQueryParams) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.OrderBy, validation.When(m.OrderBy != "", validation.In("id", "include", "date", "title", "slug").Error("无效的排序字段"))),
	)
}

// All List all coupons
func (s couponService) All(params CouponsQueryParams) (items []Coupon, isLastPage bool, err error) {
	if err = params.Validate(); err != nil {
		return
	}

	var res []Coupon
	params.TidyVars()
	urlValues, _ := query.Values(params)
	resp, err := s.httpClient.R().SetQueryParamsFromValues(urlValues).Get("/coupons")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		if err = jsoniter.Unmarshal(resp.Body(), &res); err == nil {
			items = res
			isLastPage = len(items) < params.PerPage
		}
	}
	return
}

// One Retrieve a coupon
func (s couponService) One(id int) (item Coupon, err error) {
	resp, err := s.httpClient.R().Get(fmt.Sprintf("/coupons/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Create

type CreateCouponRequest struct {
	Code             string `json:"code"`
	DiscountType     string `json:"discount_type"`
	Amount           string `json:"amount"`
	IndividualUse    bool   `json:"individual_use"`
	ExcludeSaleItems bool   `json:"exclude_sale_items"`
	MinimumAmount    string `json:"minimum_amount"`
}

func (m CreateCouponRequest) Validate() error {
	return nil
}

// Create Create a coupon
func (s couponService) Create(req CreateCouponRequest) (item Coupon, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().Post("/coupons")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

type UpdateCouponRequest = CreateCouponRequest

// Update Update a coupon
func (s couponService) Update(id int, req UpdateCouponRequest) (item Coupon, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().Put(fmt.Sprintf("/coupons/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Delete Delete a coupon
func (s couponService) Delete(id int) (item Coupon, err error) {
	resp, err := s.httpClient.R().Delete(fmt.Sprintf("/coupons/%d", id))
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &item)
	}
	return
}

// Batch update coupons

type BatchCreateCouponRequest = CreateTagRequest
type BatchUpdateCouponRequest struct {
	ID string `json:"id"`
	CreateTagRequest
}

type BatchCouponRequest struct {
	Create []BatchCreateCouponRequest `json:"create"`
	Update []BatchUpdateCouponRequest `json:"update"`
	Delete []int                      `json:"delete"`
}

func (m BatchCouponRequest) Validate() error {
	if len(m.Create) == 0 && len(m.Update) == 0 && len(m.Delete) == 0 {
		return errors.New("无效的请求数据")
	}
	return nil
}

type BatchCouponResult struct {
	Create []Coupon `json:"create"`
	Update []Coupon `json:"update"`
	Delete []Coupon `json:"delete"`
}

// Batch Batch create/update/delete coupons
func (s couponService) Batch(req BatchCouponRequest) (res BatchCouponResult, err error) {
	if err = req.Validate(); err != nil {
		return
	}

	resp, err := s.httpClient.R().SetBody(req).Post("/coupons/batch")
	if err != nil {
		return
	}

	if resp.IsSuccess() {
		err = jsoniter.Unmarshal(resp.Body(), &res)
	}
	return
}
