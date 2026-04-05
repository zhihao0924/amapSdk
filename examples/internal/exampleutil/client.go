package exampleutil

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/zhihao0924/amapSdk"
)

const defaultTimeout = 10 * time.Second

// NewClient creates a demo client from the AMAP_KEY environment variable.
func NewClient() (*amap.Client, error) {
	key := strings.TrimSpace(os.Getenv("AMAP_KEY"))
	if key == "" {
		return nil, fmt.Errorf("missing AMAP_KEY; run `export AMAP_KEY=your_key` first")
	}

	return amap.NewClient(&amap.Config{
		Key:     key,
		Debug:   true,
		Timeout: int(defaultTimeout / time.Second),
	})
}

// NewRequestContext returns a bounded context for example API calls.
func NewRequestContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), defaultTimeout)
}
