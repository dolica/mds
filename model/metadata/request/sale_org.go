package metadata

type ScopeRequest struct {
	Latitude  float32 `json:"lat"`
	Langitude float32 `json:"lang"`
	Scope     int     `json:"scope"`
}
