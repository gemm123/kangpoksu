package model

type Cost struct {
	Value int    `json:"value"`
	Etd   string `json:"etd"`
	Note  string `json:"note"`
}

type Costs struct {
	Service     string `json:"service"`
	Description string `json:"description"`
	Cost        []Cost `json:"cost"`
}

type ResultCost struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Costs []Costs `json:"costs"`
}

type QueryCost struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Weight      int    `json:"weight"`
	Courier     string `json:"courier"`
}

type Details struct {
	CityId     string `json:"city_id"`
	ProvinceId string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

type resultDataCost struct {
	Query              QueryCost    `json:"query"`
	Status             Status       `json:"status"`
	OriginDetails      Details      `json:"origin_details"`
	DestinationDetails Details      `json:"destination_details"`
	Result             []ResultCost `json:"results"`
}

type RespRajaOngkirCost struct {
	RajaOngkir resultDataCost `json:"rajaongkir"`
}
