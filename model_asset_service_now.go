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

type AssetServiceNow struct {
	AssetId          string `json:"asset_id"`
	Organization     string `json:"organization,omitempty"`
	Site             string `json:"site,omitempty"`
	DetectedBy       string `json:"detected_by,omitempty"`
	Type_            string `json:"type,omitempty"`
	SysClassName     string `json:"sys_class_name,omitempty"`
	OsVendor         string `json:"os_vendor,omitempty"`
	OsProduct        string `json:"os_product,omitempty"`
	OsVersion        string `json:"os_version,omitempty"`
	HwVendor         string `json:"hw_vendor,omitempty"`
	HwProduct        string `json:"hw_product,omitempty"`
	HwVersion        string `json:"hw_version,omitempty"`
	IpAddress        string `json:"ip_address,omitempty"`
	AddressesScope   string `json:"addresses_scope,omitempty"`
	AddressesExtra   string `json:"addresses_extra,omitempty"`
	MacAddress       string `json:"mac_address,omitempty"`
	MacManufacturer  string `json:"mac_manufacturer,omitempty"`
	NewestMacAge     string `json:"newest_mac_age,omitempty"`
	Macs             string `json:"macs,omitempty"`
	MacVendors       string `json:"mac_vendors,omitempty"`
	Name             string `json:"name,omitempty"`
	Tags             string `json:"tags,omitempty"`
	Domains          string `json:"domains,omitempty"`
	ServiceCount     int64  `json:"service_count,omitempty"`
	ServiceCountTcp  int64  `json:"service_count_tcp,omitempty"`
	ServiceCountUdp  int64  `json:"service_count_udp,omitempty"`
	ServiceCountArp  int64  `json:"service_count_arp,omitempty"`
	ServiceCountIcmp int64  `json:"service_count_icmp,omitempty"`
	LowestTtl        int64  `json:"lowest_ttl,omitempty"`
	LowestRtt        int64  `json:"lowest_rtt,omitempty"`
	Alive            bool   `json:"alive,omitempty"`
	FirstDiscovered  string `json:"first_discovered,omitempty"`
	LastDiscovered   string `json:"last_discovered,omitempty"`
	LastUpdated      string `json:"last_updated,omitempty"`
	Comments         string `json:"comments,omitempty"`
}