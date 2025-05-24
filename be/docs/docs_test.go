// Package docs_test contains tests for the docs package
package docs_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"legend_score/docs"
)

func TestSwaggerInfo(t *testing.T) {
	// Test that SwaggerInfo is initialized with the expected values
	assert.Equal(t, "1.0", docs.SwaggerInfo.Version, "Version should be 1.0")
	assert.Equal(t, "localhost:1323", docs.SwaggerInfo.Host, "Host should be localhost:1323")
	assert.Equal(t, "/api/v1", docs.SwaggerInfo.BasePath, "BasePath should be /api/v1")
	assert.Equal(t, []string{"http"}, docs.SwaggerInfo.Schemes, "Schemes should be [http]")
	assert.Equal(t, "Legend Score API", docs.SwaggerInfo.Title, "Title should be Legend Score API")
	assert.Equal(t, "This is the API documentation for Legend Score application.", docs.SwaggerInfo.Description,
		"Description should match expected value")
	assert.Equal(t, "swagger", docs.SwaggerInfo.InfoInstanceName, "InfoInstanceName should be swagger")
}