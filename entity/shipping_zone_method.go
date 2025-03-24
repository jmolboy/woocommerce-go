package entity

// ShippingZoneMethod shipping zone method properties
type ShippingZoneMethod = ShippingMethod

// ShippingZoneMethodSetting shipping zone method setting properties
type ShippingZoneMethodSetting struct {
	Title         ShippingZoneMethodSettingItem      `json:"title"`
	TaxStatus     ShippingZoneMethodSettingItem_Tax  `json:"tax_status"`
	Cost          ShippingZoneMethodSettingItem      `json:"cost"`
	ClassCosts    ShippingZoneMethodSettingItem      `json:"class_costs"`
	ClassCost1418 ShippingZoneMethodSettingItem      `json:"class_cost_1418"`
	NoClassCost   ShippingZoneMethodSettingItem      `json:"no_class_cost"`
	Type          ShippingZoneMethodSettingItem_Type `json:"type"`
}

type ShippingZoneMethodSettingItem struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Default     string `json:"default"`
	Tip         string `json:"tip"`
	PlaceHolder string `json:"place_holder"`
}

type ShippingZoneMethodSettingItem_Tax struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Default     string `json:"default"`
	Tip         string `json:"tip"`
	PlaceHolder string `json:"place_holder"`
	Options     struct {
		Taxable string `json:"taxable"`
		None    string `json:"none"`
	} `json:"options"`
}

type ShippingZoneMethodSettingItem_Type struct {
	ID          string `json:"id"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Value       string `json:"value"`
	Default     string `json:"default"`
	Tip         string `json:"tip"`
	PlaceHolder string `json:"place_holder"`
	Options     struct {
		Class string `json:"class"`
		Order string `json:"order"`
	} `json:"options"`
}
