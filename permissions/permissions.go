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
)
