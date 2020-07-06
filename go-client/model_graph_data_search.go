/*
 * Catalog Services API
 *
 * Catalog Services Web API Documentation
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GraphDataSearch struct {
	Parent string `json:"parent,omitempty"`
	Key string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
	SortByKey bool `json:"sortByKey,omitempty"`
	IncludeDescendents bool `json:"includeDescendents,omitempty"`
	EncodeValues bool `json:"encodeValues,omitempty"`
	TopLevelOnly bool `json:"topLevelOnly,omitempty"`
}