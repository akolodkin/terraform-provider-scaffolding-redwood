/*
 * Redwood User API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type RwPermissionSetPermission struct {
	PermissionSetId int32 `json:"PermissionSetId,omitempty"`
	PermissionId string `json:"PermissionId,omitempty"`
	ApplicationId int32 `json:"ApplicationId,omitempty"`
	IsDeny bool `json:"IsDeny,omitempty"`
}
