package models

// DrivingResponse 驾车路径规划响应
type DrivingResponse struct {
	BaseResponse
	Route DrivingRoute `json:"route"`
}

// DrivingRoute 驾车路径规划结果
type DrivingRoute struct {
	Origin      string        `json:"origin"`
	Destination string        `json:"destination"`
	Paths       []DrivingPath `json:"paths"`
	Taxi        TaxiInfo      `json:"taxi"`
}

// TaxiInfo 出租车信息
type TaxiInfo struct {
	Distance int `json:"distance"`
	Duration int `json:"duration"`
	Cost     int `json:"cost"`
}

// DrivingPath 驾车路径规划路径
type DrivingPath struct {
	Distance     string        `json:"distance"`
	Duration     string        `json:"duration"`
	Steps        []DrivingStep `json:"steps"`
	Strategy     string        `json:"strategy"`
	Tolls        int           `json:"tolls"`
	TollDistance string        `json:"toll_distance"`
}

// DrivingStep 驾车路径规划步骤
type DrivingStep struct {
	Instruction string `json:"instruction"`
	Distance    int    `json:"distance"`
	Duration    int    `json:"duration"`
	Action      string `json:"action"`
	Road        string `json:"road"`
	Polyline    string `json:"polyline"`
}

// WalkingResponse 步行路径规划响应
type WalkingResponse struct {
	BaseResponse
	Route WalkingRoute `json:"route"`
}

// WalkingRoute 步行路径规划结果
type WalkingRoute struct {
	Paths []WalkingPath `json:"paths"`
}

// WalkingPath 步行路径规划路径
type WalkingPath struct {
	Distance string        `json:"distance"`
	Duration string        `json:"duration"`
	Steps    []WalkingStep `json:"steps"`
}

// WalkingStep 步行路径规划步骤
type WalkingStep struct {
	Instruction string `json:"instruction"`
	Distance    int    `json:"distance"`
	Duration    int    `json:"duration"`
	Action      string `json:"action"`
	Road        string `json:"road"`
}
