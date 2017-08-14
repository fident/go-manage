package permissions

/**
* Fident permission keys
**/

type entry int

const (
	// PermissionAll allows account to perform all requests
	PermissionAll = "fident/management/*"

	// PermissionsReadAll allows account perform all read requests
	PermissionsReadAll = "fident/management/read"

	// PermissionGetLastLoginTimestamps is the permission required to read login timestamps
	PermissionGetLastLoginTimestamps = "fident/management/read/login-timestamps"

	// PermissionGetAccountDetails is the permission required to read account details
	PermissionGetAccountDetails = "fident/management/read/account-details"

	// PermissionGetAllIdentityIDs is the permission required to read all identity IDs
	PermissionGetAllIdentityIDs = "fident/management/read/all-identity-ids"

	// PermissionAddManagementPermissions is the permission required to add management permissions to an identity
	PermissionAddManagementPermissions = "fident/management/write/add-management-permission"

	// PermissionGetManagementPermissions is the permission required to retrieve assigned managemnent permissions for an identity
	PermissionGetManagementPermissions = "fident/management/read/management-permissions"

	// PermissionRemoveManagementPermissions is the permission required to remove management permissions from an identity
	PermissionRemoveManagementPermissions = "fident/management/write/remove-management-permission"
)
