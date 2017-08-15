package permissions

/**
* Fident permission keys
**/

// Entry is a single permission entry
type Entry string

const (
	// PermissionAll allows account to perform all requests
	PermissionAll Entry = "fident/management/*"

	// PermissionsReadAll allows account perform all read requests
	PermissionsReadAll Entry = "fident/management/read"

	// PermissionGetLastLoginTimestamps is the permission required to read login timestamps
	PermissionGetLastLoginTimestamps Entry = "fident/management/read/login-timestamps"

	// PermissionGetAccountDetails is the permission required to read account details
	PermissionGetAccountDetails Entry = "fident/management/read/account-details"

	// PermissionGetAllIdentityIDs is the permission required to read all identity IDs
	PermissionGetAllIdentityIDs Entry = "fident/management/read/all-identity-ids"

	// PermissionAddManagementPermissions is the permission required to add management permissions to an identity
	PermissionAddManagementPermissions Entry = "fident/management/write/add-management-permission"

	// PermissionGetManagementPermissions is the permission required to retrieve assigned managemnent permissions for an identity
	PermissionGetManagementPermissions Entry = "fident/management/read/management-permissions"

	// PermissionRemoveManagementPermissions is the permission required to remove management permissions from an identity
	PermissionRemoveManagementPermissions Entry = "fident/management/write/remove-management-permission"
)
