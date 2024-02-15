package model

type Result struct {
	CityID     string `json:"city_id"`
	ProvinceID string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type Query struct {
	Province string `json:"province"`
	ID       string `json:"id"`
}

type resultData struct {
	Query  []Query  `json:"query"`
	Status Status   `json:"status"`
	Result []Result `json:"results"`
}

type RespRajaOngkir struct {
	RajaOngkir resultData `json:"rajaongkir"`
}
