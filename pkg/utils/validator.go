package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io"
	"net/http"
	"reflect"
)

// hasTag checks if the given struct (or pointer to struct) has any field with the specified tag.
func hasTag(obj interface{}, tagName string) bool {
	if obj == nil {
		return false
	}
	val := reflect.ValueOf(obj)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return false
	}

	typ := val.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if _, ok := field.Tag.Lookup(tagName); ok {
			return true
		}
	}
	return false
}

// BindAndValidate is a unified helper function to bind and strictly validate parameters
// from URI, Query, and JSON body based on the request method.
// It enforces strictness by disallowing unknown fields in JSON bodies and unknown query parameters.
// allowedQueryParams should be a list of expected query parameter names (e.g., "page", "page_size").
func BindAndValidate(c *gin.Context, obj interface{}, allowedQueryParams ...string) error {
	// 1. Bind URI parameters (path variables like :id), only if URI params exist in the route AND the struct has uri tags
	if len(c.Params) > 0 && hasTag(obj, "uri") {
		if err := c.ShouldBindUri(obj); err != nil {
			return fmt.Errorf(Translate(err))
		}
	}

	// 2. Bind and strictly validate Query parameters
	if len(allowedQueryParams) > 0 && hasTag(obj, "form") {
		if err := c.ShouldBindQuery(obj); err != nil {
			return fmt.Errorf(Translate(err))
		}

		queryParams := c.Request.URL.Query()
		allowedMap := make(map[string]bool, len(allowedQueryParams))
		for _, p := range allowedQueryParams {
			allowedMap[p] = true
		}

		for param := range queryParams {
			if !allowedMap[param] {
				return fmt.Errorf("unknown query parameter: %s", param)
			}
		}
	}

	// 3. Bind and strictly validate JSON body (for POST requests)
	if c.Request.Method == http.MethodPost && c.Request.ContentLength > 0 && hasTag(obj, "json") {
		// Read the body to a buffer
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			return fmt.Errorf("failed to read request body: %w", err)
		}

		// Restore the body so it can be read again by Gin's validator
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

		// Create a new decoder and enforce strict field checking
		decoder := json.NewDecoder(bytes.NewBuffer(body))
		decoder.DisallowUnknownFields()

		// Attempt to decode. If there are unknown fields, it will error out here.
		if err := decoder.Decode(obj); err != nil {
			return fmt.Errorf("JSON body binding error: %w", err)
		}

		// If strict decoding was successful, run Gin's default validation
		// We must restore the body again before this call.
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		if err := c.ShouldBindWith(obj, binding.JSON); err != nil {
			return fmt.Errorf(Translate(err))
		}
	}

	return nil
}
