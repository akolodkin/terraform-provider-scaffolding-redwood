/*
 * Redwood User API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RwRolePermission struct {
	RoleId int32 `json:"RoleId"`
	PermissionId string `json:"PermissionId"`
	ApplicationId int32 `json:"ApplicationId,omitempty"`
	IsDeny bool `json:"IsDeny,omitempty"`
}
