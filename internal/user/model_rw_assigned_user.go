/*
 * Redwood User API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

type RwAssignedUser struct {
	UserId string `json:"UserId,omitempty"`
	RoleId int32 `json:"RoleId,omitempty"`
	FirstName string `json:"FirstName,omitempty"`
	LastName string `json:"LastName,omitempty"`
	Email string `json:"Email,omitempty"`
	JobTitle string `json:"JobTitle,omitempty"`
	StartDate time.Time `json:"StartDate,omitempty"`
	EndDate time.Time `json:"EndDate,omitempty"`
}
