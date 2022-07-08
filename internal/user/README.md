# Go API client for swagger

No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to */*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ApplicationApi* | [**CreateApplication**](docs/ApplicationApi.md#createapplication) | **Post** /api/application | 
*ApplicationApi* | [**DeleteApplicationById**](docs/ApplicationApi.md#deleteapplicationbyid) | **Delete** /api/application/{applicationId} | 
*ApplicationApi* | [**GetApplication**](docs/ApplicationApi.md#getapplication) | **Get** /api/application/{query} | 
*ApplicationApi* | [**GetApplications**](docs/ApplicationApi.md#getapplications) | **Get** /api/applications | 
*ApplicationApi* | [**UpdateApplication**](docs/ApplicationApi.md#updateapplication) | **Put** /api/application | 
*GroupApi* | [**GetApplicationsByGroupId**](docs/GroupApi.md#getapplicationsbygroupid) | **Get** /api/group/{groupId}/applications | 
*GroupApi* | [**GetGroupById**](docs/GroupApi.md#getgroupbyid) | **Get** /api/group/{groupId} | 
*GroupApi* | [**GetGroupPermissionTreeByGroupId**](docs/GroupApi.md#getgrouppermissiontreebygroupid) | **Get** /api/group/{groupId}/permissionTree | 
*GroupApi* | [**GetPermissionsByGroupAndApplication**](docs/GroupApi.md#getpermissionsbygroupandapplication) | **Get** /api/group/{groupId}/permissions/{applicationId} | 
*GroupApi* | [**GetPermissionsByGroupIdAndApplicationKey**](docs/GroupApi.md#getpermissionsbygroupidandapplicationkey) | **Get** /api/group/{groupId}/permissions/{applicationKey} | 
*GroupApi* | [**SearchGroups**](docs/GroupApi.md#searchgroups) | **Get** /api/groups/search | 
*PermissionApi* | [**CreatePermission**](docs/PermissionApi.md#createpermission) | **Post** /api/permission | 
*PermissionApi* | [**DeletePermissionById**](docs/PermissionApi.md#deletepermissionbyid) | **Delete** /api/permission/{permissionId}/{applicationId} | 
*PermissionApi* | [**DisablePermissionById**](docs/PermissionApi.md#disablepermissionbyid) | **Patch** /api/permission/{permissionId}/{applicationId}/disable | 
*PermissionApi* | [**EnablePermission**](docs/PermissionApi.md#enablepermission) | **Patch** /api/permission/{permissionId}/{applicationId}/enable | 
*PermissionApi* | [**GetPermissionById**](docs/PermissionApi.md#getpermissionbyid) | **Get** /api/permission/{permissionId}/{applicationId} | 
*PermissionApi* | [**GetPermissionsByApplicationId**](docs/PermissionApi.md#getpermissionsbyapplicationid) | **Get** /api/permissions/{applicationId} | 
*PermissionApi* | [**UpdatePermission**](docs/PermissionApi.md#updatepermission) | **Put** /api/permission | 
*PermissionSetApi* | [**AddPermissionToPermissionSet**](docs/PermissionSetApi.md#addpermissiontopermissionset) | **Post** /api/permission-set/permission | 
*PermissionSetApi* | [**CreatePermissionSet**](docs/PermissionSetApi.md#createpermissionset) | **Post** /api/permission-set | 
*PermissionSetApi* | [**DeletePermissionSetById**](docs/PermissionSetApi.md#deletepermissionsetbyid) | **Delete** /api/permission-set/{permissionSetId} | 
*PermissionSetApi* | [**DisablePermissionSetById**](docs/PermissionSetApi.md#disablepermissionsetbyid) | **Patch** /api/permission-set/{permissionSetId}/disable | 
*PermissionSetApi* | [**EnablePermissionSetById**](docs/PermissionSetApi.md#enablepermissionsetbyid) | **Patch** /api/permission-set/{permissionSetId}/enable | 
*PermissionSetApi* | [**GetPermissionSetById**](docs/PermissionSetApi.md#getpermissionsetbyid) | **Get** /api/permission-set/{permissionSetId} | 
*PermissionSetApi* | [**GetPermissionSetsByApplicationId**](docs/PermissionSetApi.md#getpermissionsetsbyapplicationid) | **Get** /api/permission-sets/{applicationId} | 
*PermissionSetApi* | [**GetPermissionsByPermissionSetId**](docs/PermissionSetApi.md#getpermissionsbypermissionsetid) | **Get** /api/permission-set/{permissionSetId}/permissions | 
*PermissionSetApi* | [**RemovePermissionFromPermissionSet**](docs/PermissionSetApi.md#removepermissionfrompermissionset) | **Delete** /api/permission-set/{permissionSetId}/permission/{permissionId} | 
*PermissionSetApi* | [**UpdatePermissionSet**](docs/PermissionSetApi.md#updatepermissionset) | **Put** /api/permission-set | 
*RoleApi* | [**AddGroupToRole**](docs/RoleApi.md#addgrouptorole) | **Post** /api/role/group | 
*RoleApi* | [**AddPermissionSetToRole**](docs/RoleApi.md#addpermissionsettorole) | **Post** /api/role/permission-set | 
*RoleApi* | [**AddPermissionToRole**](docs/RoleApi.md#addpermissiontorole) | **Post** /api/role/permission | 
*RoleApi* | [**AddUserToRole**](docs/RoleApi.md#addusertorole) | **Post** /api/role/user | 
*RoleApi* | [**CreateRole**](docs/RoleApi.md#createrole) | **Post** /api/role | 
*RoleApi* | [**DeleteRoleById**](docs/RoleApi.md#deleterolebyid) | **Delete** /api/role/{roleId} | 
*RoleApi* | [**DisableRoleById**](docs/RoleApi.md#disablerolebyid) | **Patch** /api/role/{roleId}/disable | 
*RoleApi* | [**EnableRoleById**](docs/RoleApi.md#enablerolebyid) | **Patch** /api/role/{roleId}/enable | 
*RoleApi* | [**GetGroupsByRoleId**](docs/RoleApi.md#getgroupsbyroleid) | **Get** /api/role/{roleId}/groups | 
*RoleApi* | [**GetPermissionSetsByRoleId**](docs/RoleApi.md#getpermissionsetsbyroleid) | **Get** /api/role/{roleId}/permission-sets | 
*RoleApi* | [**GetPermissionsByRoleId**](docs/RoleApi.md#getpermissionsbyroleid) | **Get** /api/role/{roleId}/permissions | 
*RoleApi* | [**GetRoleById**](docs/RoleApi.md#getrolebyid) | **Get** /api/role/{roleId} | 
*RoleApi* | [**GetRolesByApplicationId**](docs/RoleApi.md#getrolesbyapplicationid) | **Get** /api/roles/{applicationId} | 
*RoleApi* | [**GetUsersByRoleId**](docs/RoleApi.md#getusersbyroleid) | **Get** /api/role/{roleId}/users | 
*RoleApi* | [**RemoveGroupFromRole**](docs/RoleApi.md#removegroupfromrole) | **Delete** /api/role/{roleId}/group/{groupId} | 
*RoleApi* | [**RemovePermissionFromRole**](docs/RoleApi.md#removepermissionfromrole) | **Delete** /api/role/{roleId}/permission/{permissionId} | 
*RoleApi* | [**RemovePermissionSetFromRole**](docs/RoleApi.md#removepermissionsetfromrole) | **Delete** /api/role/{roleId}/permission-set/{permissionSetId} | 
*RoleApi* | [**RemoveUserFromRole**](docs/RoleApi.md#removeuserfromrole) | **Delete** /api/role/{roleId}/user/{userId} | 
*RoleApi* | [**UpdateRole**](docs/RoleApi.md#updaterole) | **Put** /api/role | 
*TenantApi* | [**CreateTenant**](docs/TenantApi.md#createtenant) | **Post** /api/tenants | 
*TenantApi* | [**DeleteTenant**](docs/TenantApi.md#deletetenant) | **Delete** /api/tenants/{id} | 
*TenantApi* | [**GetTenant**](docs/TenantApi.md#gettenant) | **Get** /api/tenants/{id} | 
*TenantApi* | [**GetTenantByShortName**](docs/TenantApi.md#gettenantbyshortname) | **Get** /api/tenants/shortname/{shortName} | 
*TenantApi* | [**GetTenants**](docs/TenantApi.md#gettenants) | **Get** /api/tenants | 
*TenantApi* | [**UpdateTenant**](docs/TenantApi.md#updatetenant) | **Put** /api/tenants | 
*UserApi* | [**GetApplicationsByUserId**](docs/UserApi.md#getapplicationsbyuserid) | **Get** /api/user/{userId}/applications | 
*UserApi* | [**GetGroupsByUserId**](docs/UserApi.md#getgroupsbyuserid) | **Get** /api/user/{userId}/groups | 
*UserApi* | [**GetLoadRunnderIdByAdUserId**](docs/UserApi.md#getloadrunnderidbyaduserid) | **Get** /api/user/loadRunner/{adUserId}/id | 
*UserApi* | [**GetLoadRunnderIdByEmail**](docs/UserApi.md#getloadrunnderidbyemail) | **Get** /api/user/loadRunner/{email}/id | 
*UserApi* | [**GetLoadRunnerUserByUserId**](docs/UserApi.md#getloadrunneruserbyuserid) | **Get** /api/user/loadRunner/{userId} | 
*UserApi* | [**GetLoadRunnerUsersByUserIds**](docs/UserApi.md#getloadrunnerusersbyuserids) | **Post** /api/users/loadRunner | 
*UserApi* | [**GetPermissionsByUserAndApplication**](docs/UserApi.md#getpermissionsbyuserandapplication) | **Get** /api/user/{userId}/permissions/{applicationId} | 
*UserApi* | [**GetPermissionsByUserAndApplicationKey**](docs/UserApi.md#getpermissionsbyuserandapplicationkey) | **Get** /api/user/{userId}/permissions/{applicationKey} | 
*UserApi* | [**GetUserById**](docs/UserApi.md#getuserbyid) | **Get** /api/user/{userId} | 
*UserApi* | [**GetUserGroupsByEmail**](docs/UserApi.md#getusergroupsbyemail) | **Get** /api/user/email/groups | 
*UserApi* | [**GetUserPermissionTreeByUserId**](docs/UserApi.md#getuserpermissiontreebyuserid) | **Get** /api/user/{userId}/permissionTree | 
*UserApi* | [**GetUsersByUserIds**](docs/UserApi.md#getusersbyuserids) | **Post** /api/users | 
*UserApi* | [**SearchLoadRunnerUsers**](docs/UserApi.md#searchloadrunnerusers) | **Get** /api/users/loadRunner/search | 
*UserApi* | [**SearchUsers**](docs/UserApi.md#searchusers) | **Get** /api/users/search | 
*UserTenantApi* | [**CreateUserTenant**](docs/UserTenantApi.md#createusertenant) | **Post** /api/users/tenants | 
*UserTenantApi* | [**DisableUserTenant**](docs/UserTenantApi.md#disableusertenant) | **Delete** /api/users/adUsers/{adUserId}/tenants/{tenantId} | 
*UserTenantApi* | [**EnableUserTenant**](docs/UserTenantApi.md#enableusertenant) | **Put** /api/users/adUsers/{adUserId}/tenants/{tenantId}/enable | 
*UserTenantApi* | [**GetUserTenants**](docs/UserTenantApi.md#getusertenants) | **Get** /api/users/tenants | 
*UserTenantApi* | [**GetUserTenantsByADUserId**](docs/UserTenantApi.md#getusertenantsbyaduserid) | **Get** /api/users/adUsers/{adUserId}/tenants | 

## Documentation For Models

 - [IHashFunction](docs/IHashFunction.md)
 - [RwApplication](docs/RwApplication.md)
 - [RwAssignedGroup](docs/RwAssignedGroup.md)
 - [RwAssignedPermission](docs/RwAssignedPermission.md)
 - [RwAssignedPermissionSet](docs/RwAssignedPermissionSet.md)
 - [RwAssignedUser](docs/RwAssignedUser.md)
 - [RwGroup](docs/RwGroup.md)
 - [RwLoadRunnerAdUser](docs/RwLoadRunnerAdUser.md)
 - [RwLoadRunnerUser](docs/RwLoadRunnerUser.md)
 - [RwPermission](docs/RwPermission.md)
 - [RwPermissionReportItem](docs/RwPermissionReportItem.md)
 - [RwPermissionSet](docs/RwPermissionSet.md)
 - [RwPermissionSetPermission](docs/RwPermissionSetPermission.md)
 - [RwRole](docs/RwRole.md)
 - [RwRoleGroup](docs/RwRoleGroup.md)
 - [RwRolePermission](docs/RwRolePermission.md)
 - [RwRolePermissionSet](docs/RwRolePermissionSet.md)
 - [RwRoleUser](docs/RwRoleUser.md)
 - [RwUser](docs/RwUser.md)
 - [RwUserType](docs/RwUserType.md)
 - [StringRwPermissionReportItemITreeNode](docs/StringRwPermissionReportItemITreeNode.md)
 - [Tenant](docs/Tenant.md)
 - [UserTenant](docs/UserTenant.md)

## Documentation For Authorization

## Authentication
## Authorization
- **Type**: API key 

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextAPIKey, sw.APIKey{
	Key: "APIKEY",
	Prefix: "Bearer", // Omit if not necessary.
})
r, err := client.Service.Operation(auth, args)
```

## Author

