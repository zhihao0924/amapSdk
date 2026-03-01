package models

// TextSearchResponse POI搜索响应
type TextSearchResponse struct {
	BaseResponse
	Count      int         `json:"count"`
	Pois       []PoiInfo   `json:"pois"`
	Suggestion *Suggestion `json:"suggestion,omitempty"`
}

// AroundSearchResponse 周边搜索响应
type AroundSearchResponse struct {
	BaseResponse
	Count      int         `json:"count"`
	Pois       []PoiInfo   `json:"pois"`
	Suggestion *Suggestion `json:"suggestion,omitempty"`
}

// SearchByPolygonResponse 多边形搜索响应
type SearchByPolygonResponse struct {
	BaseResponse
	Count      int         `json:"count"`
	Pois       []PoiInfo   `json:"pois"`
	Suggestion *Suggestion `json:"suggestion,omitempty"`
}
