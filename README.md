# Go API client for Rumble Network Discovery

Rumble Network Discovery API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.4
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientExperimentalCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import sw "./openapi"
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), sw.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identifield by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), sw.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), sw.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://console.rumble.run/api/v1.0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ExportApi* | [**ExportAssetSyncCreatedJSON**](docs/ExportApi.md#exportassetsynccreatedjson) | **Get** /export/org/assets/sync/created/assets.json | Exports the asset inventory in a sync-friendly manner using created_at as a checkpoint.
*ExportApi* | [**ExportAssetSyncUpdatedJSON**](docs/ExportApi.md#exportassetsyncupdatedjson) | **Get** /export/org/assets/sync/updated/assets.json | Exports the asset inventory in a sync-friendly manner using updated_at as a checkpoint.
*ExportApi* | [**ExportAssetsCSV**](docs/ExportApi.md#exportassetscsv) | **Get** /export/org/assets.csv | Asset inventory as CSV.
*ExportApi* | [**ExportAssetsJSON**](docs/ExportApi.md#exportassetsjson) | **Get** /export/org/assets.json | Exports the asset inventory.
*ExportApi* | [**ExportAssetsJSONL**](docs/ExportApi.md#exportassetsjsonl) | **Get** /export/org/assets.jsonl | Asset inventory as JSON line-delimited.
*ExportApi* | [**ExportAssetsNmapXML**](docs/ExportApi.md#exportassetsnmapxml) | **Get** /export/org/assets.nmap.xml | Asset inventory as Nmap-style XML.
*ExportApi* | [**ExportServicesCSV**](docs/ExportApi.md#exportservicescsv) | **Get** /export/org/services.csv | Service inventory as CSV.
*ExportApi* | [**ExportServicesJSON**](docs/ExportApi.md#exportservicesjson) | **Get** /export/org/services.json | Service inventory as JSON.
*ExportApi* | [**ExportServicesJSONL**](docs/ExportApi.md#exportservicesjsonl) | **Get** /export/org/services.jsonl | Service inventory as JSON line-delimited.
*ExportApi* | [**ExportSitesCSV**](docs/ExportApi.md#exportsitescsv) | **Get** /export/org/sites.csv | Site list as CSV.
*ExportApi* | [**ExportSitesJSON**](docs/ExportApi.md#exportsitesjson) | **Get** /export/org/sites.json | Export all sites.
*ExportApi* | [**ExportSitesJSONL**](docs/ExportApi.md#exportsitesjsonl) | **Get** /export/org/sites.jsonl | Site list as JSON line-delimited.
*ExportApi* | [**ExportWirelessCSV**](docs/ExportApi.md#exportwirelesscsv) | **Get** /export/org/wireless.csv | Wireless inventory as CSV.
*ExportApi* | [**ExportWirelessJSON**](docs/ExportApi.md#exportwirelessjson) | **Get** /export/org/wireless.json | Wireless inventory as JSON.
*ExportApi* | [**ExportWirelessJSONL**](docs/ExportApi.md#exportwirelessjsonl) | **Get** /export/org/wireless.jsonl | Wireless inventory as JSON line-delimited.
*OrganizationApi* | [**CreateScan**](docs/OrganizationApi.md#createscan) | **Put** /org/sites/{site_id}/scan | Create a scan task for a given site.
*OrganizationApi* | [**CreateSite**](docs/OrganizationApi.md#createsite) | **Put** /org/sites | Create a new site.
*OrganizationApi* | [**GetAgent**](docs/OrganizationApi.md#getagent) | **Get** /org/agents/{agent_id} | Get details for a single agent.
*OrganizationApi* | [**GetAgents**](docs/OrganizationApi.md#getagents) | **Get** /org/agents | Get all agents.
*OrganizationApi* | [**GetAsset**](docs/OrganizationApi.md#getasset) | **Get** /org/assets/{asset_id} | Get asset details.
*OrganizationApi* | [**GetAssets**](docs/OrganizationApi.md#getassets) | **Get** /org/assets | Get all assets.
*OrganizationApi* | [**GetKey**](docs/OrganizationApi.md#getkey) | **Get** /org/key | Get API key details.
*OrganizationApi* | [**GetOrganization**](docs/OrganizationApi.md#getorganization) | **Get** /org | Get organization details.
*OrganizationApi* | [**GetService**](docs/OrganizationApi.md#getservice) | **Get** /org/services/{service_id} | Get service details.
*OrganizationApi* | [**GetServices**](docs/OrganizationApi.md#getservices) | **Get** /org/services | Get all services.
*OrganizationApi* | [**GetSite**](docs/OrganizationApi.md#getsite) | **Get** /org/sites/{site_id} | Get site details.
*OrganizationApi* | [**GetSites**](docs/OrganizationApi.md#getsites) | **Get** /org/sites | Get all sites.
*OrganizationApi* | [**GetTask**](docs/OrganizationApi.md#gettask) | **Get** /org/tasks/{task_id} | Get task details.
*OrganizationApi* | [**GetTaskChangeReport**](docs/OrganizationApi.md#gettaskchangereport) | **Get** /org/tasks/{task_id}/changes | Returns a temporary task change report data url.
*OrganizationApi* | [**GetTaskScanData**](docs/OrganizationApi.md#gettaskscandata) | **Get** /org/tasks/{task_id}/data | Returns a temporary task scan data url.
*OrganizationApi* | [**GetTasks**](docs/OrganizationApi.md#gettasks) | **Get** /org/tasks | Get all tasks (last 1000).
*OrganizationApi* | [**GetWirelessLAN**](docs/OrganizationApi.md#getwirelesslan) | **Get** /org/wirelesss/{wireless_id} | Get wireless LAN details.
*OrganizationApi* | [**GetWirelessLANs**](docs/OrganizationApi.md#getwirelesslans) | **Get** /org/wireless | Get all wireless LANs.
*OrganizationApi* | [**HideTask**](docs/OrganizationApi.md#hidetask) | **Post** /org/tasks/{task_id}/hide | Signal that a completed task should be hidden.
*OrganizationApi* | [**ImportScanData**](docs/OrganizationApi.md#importscandata) | **Put** /org/sites/{site_id}/import | Import a scan data file into a site.
*OrganizationApi* | [**RemoveAgent**](docs/OrganizationApi.md#removeagent) | **Delete** /org/agents/{agent_id} | Remove and uninstall an agent.
*OrganizationApi* | [**RemoveAsset**](docs/OrganizationApi.md#removeasset) | **Delete** /org/assets/{asset_id} | Remove an asset.
*OrganizationApi* | [**RemoveService**](docs/OrganizationApi.md#removeservice) | **Delete** /org/services/{service_id} | Remove a service.
*OrganizationApi* | [**RemoveSite**](docs/OrganizationApi.md#removesite) | **Delete** /org/sites/{site_id} | Remove a site and associated assets.
*OrganizationApi* | [**RemoveWirelessLAN**](docs/OrganizationApi.md#removewirelesslan) | **Delete** /org/wirelesss/{wireless_id} | Remove a wireless LAN.
*OrganizationApi* | [**StopTask**](docs/OrganizationApi.md#stoptask) | **Post** /org/tasks/{task_id}/stop | Signal that a task should be stopped or canceled.
*OrganizationApi* | [**UpdateAgentSite**](docs/OrganizationApi.md#updateagentsite) | **Patch** /org/agents/{agent_id} | Update the site associated with agent.
*OrganizationApi* | [**UpdateAssetComments**](docs/OrganizationApi.md#updateassetcomments) | **Patch** /org/assets/{asset_id}/comments | Update asset comments.
*OrganizationApi* | [**UpdateAssetTags**](docs/OrganizationApi.md#updateassettags) | **Patch** /org/assets/{asset_id}/tags | Update asset tags.
*OrganizationApi* | [**UpdateOrganization**](docs/OrganizationApi.md#updateorganization) | **Patch** /org | Update organization details.
*OrganizationApi* | [**UpdateSite**](docs/OrganizationApi.md#updatesite) | **Patch** /org/sites/{site_id} | Update a site definition.
*OrganizationApi* | [**UpdateTask**](docs/OrganizationApi.md#updatetask) | **Patch** /org/tasks/{task_id} | Update task parameters.
*OrganizationApi* | [**UpgradeAgent**](docs/OrganizationApi.md#upgradeagent) | **Post** /org/agents/{agent_id}/update | Force an agent to update and restart.
*PublicApi* | [**GetLatestAgentVersion**](docs/PublicApi.md#getlatestagentversion) | **Get** /releases/agent/version | Returns latest agent version.
*PublicApi* | [**GetLatestPlatformVersion**](docs/PublicApi.md#getlatestplatformversion) | **Get** /releases/platform/version | Returns latest platform version.
*PublicApi* | [**GetLatestScannerVersion**](docs/PublicApi.md#getlatestscannerversion) | **Get** /releases/scanner/version | Returns latest scanner version.


## Documentation For Models

 - [APIKey](docs/APIKey.md)
 - [Agent](docs/Agent.md)
 - [AgentSiteID](docs/AgentSiteID.md)
 - [Asset](docs/Asset.md)
 - [AssetComments](docs/AssetComments.md)
 - [AssetTags](docs/AssetTags.md)
 - [AssetsWithCheckpoint](docs/AssetsWithCheckpoint.md)
 - [ComponentVersion](docs/ComponentVersion.md)
 - [OrgOptions](docs/OrgOptions.md)
 - [Organization](docs/Organization.md)
 - [ScanOptions](docs/ScanOptions.md)
 - [Service](docs/Service.md)
 - [Site](docs/Site.md)
 - [SiteOptions](docs/SiteOptions.md)
 - [Task](docs/Task.md)
 - [URL](docs/URL.md)
 - [Wireless](docs/Wireless.md)


## Documentation For Authorization



### bearerAuth


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@rumble.run

