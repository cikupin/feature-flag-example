package common

import (
	"github.com/checkr/goflagr"
)

// GetFlagrClient return flagr api client
func GetFlagrClient() *goflagr.APIClient {
	conf := &goflagr.Configuration{
		BasePath:      "http://localhost:18000/api/v1",
		DefaultHeader: make(map[string]string),
		UserAgent:     "Swagger-Codegen/1.1.1/go",
	}
	client := goflagr.NewAPIClient(conf)
	return client
}
