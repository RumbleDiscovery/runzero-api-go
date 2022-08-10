/*
 * runZero API
 *
 * runZero API
 *
 * API version: 3.0.0
 * Contact: support@runzero.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package runzeroapi

type ScanTemplateOptions struct {
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	Params      map[string]string `json:"params,omitempty"`
	Global      bool              `json:"global"`
	Acl         ModelMap          `json:"acl"`
}
