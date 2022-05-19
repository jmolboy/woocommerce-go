package entity

type Shipping struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Company   string `json:"company,omitempty"`
	Address1  string `json:"address_1,omitempty"`
	Address2  string `json:"address_2,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Postcode  string `json:"postcode,omitempty"`
	Country   string `json:"country,omitempty"`
}

type ShippingLine struct {
	ID          int        `json:"id"`
	MethodTitle string     `json:"method_title"`
	MethodId    string     `json:"method_id"`
	Total       string     `json:"total"`
	TotalTax    string     `json:"total_tax"`
	Taxes       []Taxes    `json:"taxes"`
	MetaData    []MetaData `json:"meta_data"`
}
