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

type Organization struct {
	Id                      string `json:"id"`
	CreatedAt               int64  `json:"created_at,omitempty"`
	UpdatedAt               int64  `json:"updated_at,omitempty"`
	ClientId                string `json:"client_id,omitempty"`
	DownloadToken           string `json:"download_token,omitempty"`
	DownloadTokenCreatedAt  int64  `json:"download_token_created_at,omitempty"`
	Permanent               bool   `json:"permanent,omitempty"`
	Name                    string `json:"name"`
	Description             string `json:"description,omitempty"`
	Inactive                bool   `json:"inactive,omitempty"`
	DeactivatedAt           int64  `json:"deactivated_at,omitempty"`
	ServiceCount            int64  `json:"service_count,omitempty"`
	ServiceCountTcp         int64  `json:"service_count_tcp,omitempty"`
	ServiceCountUdp         int64  `json:"service_count_udp,omitempty"`
	ServiceCountArp         int64  `json:"service_count_arp,omitempty"`
	ServiceCountIcmp        int64  `json:"service_count_icmp,omitempty"`
	AssetCount              int64  `json:"asset_count,omitempty"`
	ExportToken             string `json:"export_token,omitempty"`
	ExportTokenCreatedAt    int64  `json:"export_token_created_at,omitempty"`
	ExportTokenLastUsedAt   int64  `json:"export_token_last_used_at,omitempty"`
	ExportTokenLastUsedBy   string `json:"export_token_last_used_by,omitempty"`
	ExportTokenCounter      int64  `json:"export_token_counter,omitempty"`
	Project                 bool   `json:"project,omitempty"`
	ParentId                string `json:"parent_id,omitempty"`
	ExpirationAssetsStale   int64  `json:"expiration_assets_stale,omitempty"`
	ExpirationAssetsOffline int64  `json:"expiration_assets_offline,omitempty"`
	ExpirationScans         int64  `json:"expiration_scans,omitempty"`
}
