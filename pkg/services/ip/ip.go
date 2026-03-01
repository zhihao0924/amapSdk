package ip

import (
	"context"
	"fmt"

	"github.com/zhihao0924/amapSdk/pkg/common"
	"github.com/zhihao0924/amapSdk/pkg/core"
	"github.com/zhihao0924/amapSdk/pkg/models"
)

// Service IP定位服务
type Service struct {
	http   *core.HTTPClient
	logger common.Logger
}

// New 创建IP定位服务
func New(http *core.HTTPClient, logger common.Logger) *Service {
	return &Service{
		http:   http,
		logger: logger,
	}
}

// LocationOptions IP定位选项
type LocationOptions struct {
	IP string `json:"ip"` // IP地址，为空时查询当前IP
}

// Location IP定位
func (s *Service) Location(opts *LocationOptions) (*models.IpLocationResponse, error) {
	params := map[string]string{
		"ip": opts.IP,
	}

	var resp models.IpLocationResponse
	err := s.http.Get(context.Background(), "/ip", params, &resp)
	if err != nil {
		s.logger.Error("IP location request failed: %v", err)
		return nil, err
	}

	if resp.IP == "" {
		resp.IP = opts.IP
	}

	s.logger.Info("IP location success: ip=%s, province=%s, city=%s",
		opts.IP, resp.Province, resp.City)
	return &resp, nil
}

// Current 查询当前IP位置
func (s *Service) Current() (*models.IpLocationResponse, error) {
	return s.Location(&LocationOptions{})
}

// GetIPInfo 获取指定IP的信息
func (s *Service) GetIPInfo(ip string) (*models.IpLocationResponse, error) {
	return s.Location(&LocationOptions{IP: ip})
}

// BatchLocation 批量IP定位
func (s *Service) BatchLocation(ips []string) ([]models.IpLocationResponse, error) {
	responses := make([]models.IpLocationResponse, len(ips))
	for i, ip := range ips {
		resp, err := s.Location(&LocationOptions{IP: ip})
		if err != nil {
			s.logger.Error("Batch IP location failed for ip=%s: %v", ip, err)
			return nil, fmt.Errorf("failed to locate ip %s: %w", ip, err)
		}
		responses[i] = *resp
	}
	return responses, nil
}
