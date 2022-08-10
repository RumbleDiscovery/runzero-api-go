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

type Group struct {
	Id             string   `json:"id"`
	Description    string   `json:"description,omitempty"`
	Name           string   `json:"name,omitempty"`
	RoleSummary    string   `json:"role_summary,omitempty"`
	UserCount      int64    `json:"user_count,omitempty"`
	CreatedByEmail string   `json:"created_by_email,omitempty"`
	CreatedAt      int64    `json:"created_at,omitempty"`
	UpdatedAt      int64    `json:"updated_at,omitempty"`
	ExpiresAt      int64    `json:"expires_at,omitempty"`
	OrgDefaultRole string   `json:"org_default_role,omitempty"`
	OrgRoles       ModelMap `json:"org_roles,omitempty"`
}