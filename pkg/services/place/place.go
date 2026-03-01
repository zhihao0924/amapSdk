package place

import (
	"context"
	"errors"
	"fmt"

	"github.com/zhihao0924/amapSdk/pkg/common"
	"github.com/zhihao0924/amapSdk/pkg/core"
	"github.com/zhihao0924/amapSdk/pkg/models"
)

// Service POI搜索服务
type Service struct {
	http   *core.HTTPClient
	logger common.Logger
}

// New 创建POI搜索服务
func New(http *core.HTTPClient, logger common.Logger) *Service {
	return &Service{
		http:   http,
		logger: logger,
	}
}

// TextSearchOptions 关键字搜索选项
type TextSearchOptions struct {
	Keywords   string `json:"keywords"`   // 关键字
	City       string `json:"city"`       // 指定查询城市
	CityLimit  string `json:"citylimit"`  // 是否限制城市
	Types      string `json:"types"`      // POI分类
	Children   string `json:"children"`   // 是否按照层级展示子节点
	Offset     string `json:"offset"`     // 每页记录数
	Page       string `json:"page"`       // 当前页数
	Extensions string `json:"extensions"` // 返回结果控制
}

// Validate 验证选项
func (o *TextSearchOptions) Validate() error {
	if o.Keywords == "" {
		return errors.New("keywords is required")
	}
	return nil
}

// TextSearch 关键字搜索POI
func (s *Service) TextSearch(opts *TextSearchOptions) (*models.TextSearchResponse, error) {
	if err := opts.Validate(); err != nil {
		s.logger.Error("TextSearch options validation failed: %v", err)
		return nil, fmt.Errorf("invalid options: %w", err)
	}

	params := map[string]string{
		"keywords":   opts.Keywords,
		"city":       opts.City,
		"citylimit":  opts.CityLimit,
		"types":      opts.Types,
		"children":   opts.Children,
		"offset":     opts.Offset,
		"page":       opts.Page,
		"extensions": opts.Extensions,
	}

	var resp models.TextSearchResponse
	err := s.http.Get(context.Background(), "/place/text", params, &resp)
	if err != nil {
		s.logger.Error("TextSearch request failed: %v", err)
		return nil, err
	}

	s.logger.Info("TextSearch success: keywords=%s, count=%d", opts.Keywords, resp.Count)
	return &resp, nil
}

// AroundSearchOptions 周边搜索选项
type AroundSearchOptions struct {
	Location   string `json:"location"`   // 中心点坐标
	Keywords   string `json:"keywords"`   // 关键字
	Type       string `json:"type"`       // POI分类
	Radius     string `json:"radius"`     // 搜索半径
	Offset     string `json:"offset"`     // 每页记录数
	Page       string `json:"page"`       // 当前页数
	Sortrule   string `json:"sortrule"`   // 排序规则
	Extensions string `json:"extensions"` // 返回结果控制
}

// Validate 验证选项
func (o *AroundSearchOptions) Validate() error {
	if o.Location == "" {
		return errors.New("location is required")
	}
	return nil
}

// AroundSearch 周边POI搜索
func (s *Service) AroundSearch(opts *AroundSearchOptions) (*models.AroundSearchResponse, error) {
	if err := opts.Validate(); err != nil {
		s.logger.Error("AroundSearch options validation failed: %v", err)
		return nil, fmt.Errorf("invalid options: %w", err)
	}

	params := map[string]string{
		"location":   opts.Location,
		"keywords":   opts.Keywords,
		"type":       opts.Type,
		"radius":     opts.Radius,
		"offset":     opts.Offset,
		"page":       opts.Page,
		"sortrule":   opts.Sortrule,
		"extensions": opts.Extensions,
	}

	var resp models.AroundSearchResponse
	err := s.http.Get(context.Background(), "/place/around", params, &resp)
	if err != nil {
		s.logger.Error("AroundSearch request failed: %v", err)
		return nil, err
	}

	s.logger.Info("AroundSearch success: location=%s, count=%d", opts.Location, resp.Count)
	return &resp, nil
}

// AroundSearchByLocation 使用Location对象进行周边搜索
func (s *Service) AroundSearchByLocation(loc *common.Location, keywords string, radius int) (*models.AroundSearchResponse, error) {
	return s.AroundSearch(&AroundSearchOptions{
		Location: loc.String(),
		Keywords: keywords,
		Radius:   fmt.Sprintf("%d", radius),
	})
}

// SearchByPolygon 多边形搜索选项
type SearchByPolygonOptions struct {
	Polygon    string `json:"polygon"`    // 多边形坐标点
	Keywords   string `json:"keywords"`   // 关键字
	Types      string `json:"types"`      // POI分类
	Offset     string `json:"offset"`     // 每页记录数
	Page       string `json:"page"`       // 当前页数
	Extensions string `json:"extensions"` // 返回结果控制
}

// Validate 验证选项
func (o *SearchByPolygonOptions) Validate() error {
	if o.Polygon == "" {
		return errors.New("polygon is required")
	}
	return nil
}

// SearchByPolygon 多边形搜索
func (s *Service) SearchByPolygon(opts *SearchByPolygonOptions) (*models.SearchByPolygonResponse, error) {
	if err := opts.Validate(); err != nil {
		s.logger.Error("SearchByPolygon options validation failed: %v", err)
		return nil, fmt.Errorf("invalid options: %w", err)
	}

	params := map[string]string{
		"polygon":    opts.Polygon,
		"keywords":   opts.Keywords,
		"types":      opts.Types,
		"offset":     opts.Offset,
		"page":       opts.Page,
		"extensions": opts.Extensions,
	}

	var resp models.SearchByPolygonResponse
	err := s.http.Get(context.Background(), "/place/polygon", params, &resp)
	if err != nil {
		s.logger.Error("SearchByPolygon request failed: %v", err)
		return nil, err
	}

	s.logger.Info("SearchByPolygon success: count=%d", resp.Count)
	return &resp, nil
}
