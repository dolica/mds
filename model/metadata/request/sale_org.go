package request

type SaleOrgRequest struct {
	Latitude   string `json:"lat"`
	Langitude  string `json:"lng"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	License    string `json:"license"`
	Scope      string `json:"scope"`
	CreateTime int    `json:"createTime"`
	Expire     int    `json:"expire"`
	Authority  string `json:"authority"`
}
