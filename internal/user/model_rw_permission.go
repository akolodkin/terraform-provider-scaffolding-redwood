/*
 * Redwood User API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RwPermission struct {
	PermissionId string `json:"PermissionId,omitempty"`
	ApplicationId int32 `json:"ApplicationId"`
	Name string `json:"Name"`
	Description string `json:"Description,omitempty"`
	Category string `json:"Category,omitempty"`
	IsUiPermission bool `json:"IsUiPermission,omitempty"`
	IsEnabled bool `json:"IsEnabled,omitempty"`
}