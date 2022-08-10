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

type ScanTemplate struct {
	Id              string            `json:"id"`
	ClientId        string            `json:"client_id,omitempty"`
	OrganizationId  string            `json:"organization_id,omitempty"`
	AgentId         string            `json:"agent_id,omitempty"`
	SiteId          string            `json:"site_id,omitempty"`
	CruncherId      string            `json:"cruncher_id,omitempty"`
	CreatedAt       int64             `json:"created_at,omitempty"`
	CreatedBy       string            `json:"created_by,omitempty"`
	CreatedByUserId string            `json:"created_by_user_id,omitempty"`
	UpdatedAt       int64             `json:"updated_at,omitempty"`
	Type_           string            `json:"type,omitempty"`
	Status          string            `json:"status,omitempty"`
	Error_          string            `json:"error,omitempty"`
	Params          map[string]string `json:"params,omitempty"`
	Stats           ModelMap          `json:"stats,omitempty"`
	Hidden          bool              `json:"hidden,omitempty"`
	ParentId        string            `json:"parent_id,omitempty"`
	Recur           bool              `json:"recur,omitempty"`
	RecurFrequency  string            `json:"recur_frequency,omitempty"`
	StartTime       int64             `json:"start_time,omitempty"`
	RecurLast       int64             `json:"recur_last,omitempty"`
	RecurNext       int64             `json:"recur_next,omitempty"`
	RecurLastTaskId string            `json:"recur_last_task_id,omitempty"`
	Name            string            `json:"name,omitempty"`
	Description     string            `json:"description,omitempty"`
	GracePeriod     string            `json:"grace_period,omitempty"`
	SourceId        string            `json:"source_id,omitempty"`
	TemplateId      string            `json:"template_id,omitempty"`
	SizeSite        int64             `json:"size_site,omitempty"`
	SizeData        int64             `json:"size_data,omitempty"`
	SizeResults     int64             `json:"size_results,omitempty"`
	HostedZoneId    string            `json:"hosted_zone_id,omitempty"`
	LinkedTaskCount int32             `json:"linked_task_count,omitempty"`
	Global          bool              `json:"global"`
	Acl             ModelMap          `json:"acl"`
}