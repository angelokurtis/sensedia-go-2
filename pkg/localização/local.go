package localização

type Local struct {
	CódigoDoPaís   string  `json:"country_code"`
	CódigoRegional string  `json:"region_code"`
	Cidade         string  `json:"city"`
	CEP            string  `json:"zip"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
}
