package models

// TextSearchResponse 关键字搜索响应
type TextSearchResponse struct {
	BaseResponse
	Count      IntOrString `json:"count"`      // 结果总数
	Pois       []PoiInfo   `json:"pois"`       // POI列表
	Suggestion *Suggestion `json:"suggestion,omitempty"` // 搜索建议
}

// AroundSearchResponse 周边搜索响应
type AroundSearchResponse struct {
	BaseResponse
	Count      IntOrString `json:"count"`      // 结果总数
	Pois       []PoiInfo   `json:"pois"`       // POI列表
	Suggestion *Suggestion `json:"suggestion,omitempty"` // 搜索建议
}

// SearchByPolygonResponse 多边形搜索响应
type SearchByPolygonResponse struct {
	BaseResponse
	Count      IntOrString `json:"count"`      // 结果总数
	Pois       []PoiInfo   `json:"pois"`       // POI列表
	Suggestion *Suggestion `json:"suggestion,omitempty"` // 搜索建议
}
