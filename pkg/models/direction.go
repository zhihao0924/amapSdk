package models

// DrivingResponse 驾车路径规划响应
type DrivingResponse struct {
	BaseResponse
	Route DrivingRoute `json:"route"` // 路线信息
}

// DrivingRoute 驾车路径规划结果
type DrivingRoute struct {
	Origin      string        `json:"origin"`      // 起点坐标
	Destination string        `json:"destination"` // 终点坐标
	Paths       []DrivingPath `json:"paths"`       // 路径列表
	TaxiCost    string        `json:"taxi_cost"`   // 出租车费用
}

// DrivingPath 驾车路径规划路径
type DrivingPath struct {
	Distance     string        `json:"distance"`      // 总距离（米）
	Duration     string        `json:"duration"`      // 总时长（秒）
	Steps        []DrivingStep `json:"steps"`         // 导航路段
	Strategy     string        `json:"strategy"`      // 路径策略
	Tolls        string        `json:"tolls"`         // 总费用（元）
	TollDistance string        `json:"toll_distance"` // 收费路段距离（米）
}

// DrivingStep 驾车路径规划步骤
type DrivingStep struct {
	Instruction     string     `json:"instruction"`      // 导航指示
	Orientation     string     `json:"orientation"`      // 方向
	Road            string     `json:"road"`             // 道路名称
	Distance        string     `json:"distance"`         // 本段距离（米）
	Tolls           string     `json:"tolls"`            // 本段费用（元）
	TollDistance    string     `json:"toll_distance"`    // 本段收费距离（米）
	TollRoad        FlexString `json:"toll_road"`        // 收费路段
	Duration        string     `json:"duration"`         // 本段时长（秒）
	Polyline        string     `json:"polyline"`         // 轨迹点集合
	Action          FlexString `json:"action"`           // 驾车动作
	AssistantAction FlexString `json:"assistant_action"` // 辅助动作
	Tmcs            []Tmc      `json:"tmcs"`             // 拥堵信息
	Cities          []City     `json:"cities"`           // 城市信息
}

// Tmc 拥堵信息
type Tmc struct {
	Lcode    []string `json:"lcode"`    // 拥堵编码
	Distance string   `json:"distance"` // 拥堵距离
	Status   string   `json:"status"`   // 拥堵状态
	Polyline string   `json:"polyline"` // 拥堵轨迹
}

// City 城市信息
type City struct {
	Name      string     `json:"name"`      // 城市名称
	Citycode  string     `json:"citycode"`  // 城市编码
	Adcode    string     `json:"adcode"`    // 区域编码
	Districts []District `json:"districts"` // 区域列表
}

// District 区域信息
type District struct {
	Name   string `json:"name"`   // 区域名称
	Adcode string `json:"adcode"` // 区域编码
}

// WalkingResponse 步行路径规划响应
type WalkingResponse struct {
	BaseResponse
	Route WalkingRoute `json:"route"` // 路线信息
}

// WalkingRoute 步行路径规划结果
type WalkingRoute struct {
	Paths []WalkingPath `json:"paths"` // 路径列表
}

// WalkingPath 步行路径规划路径
type WalkingPath struct {
	Distance string        `json:"distance"` // 总距离（米）
	Duration string        `json:"duration"` // 总时长（秒）
	Steps    []WalkingStep `json:"steps"`    // 导航路段
}

// WalkingStep 步行路径规划步骤
type WalkingStep struct {
	Instruction string      `json:"instruction"` // 导航指示
	Distance    IntOrString `json:"distance"`    // 本段距离（米）
	Duration    IntOrString `json:"duration"`    // 本段时长（秒）
	Action      FlexString  `json:"action"`      // 步行动作
	Road        FlexString  `json:"road"`        // 道路名称
}
