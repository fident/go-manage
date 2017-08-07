package permissions

/**
* Fident permission keys
**/

type entry int

const (
	// PermissionsReadAll allows account perform all read requests
	PermissionsReadAll = 800

	// PermissionGetLastLoginTimestamps is the permission required to read login timestamps
	PermissionGetLastLoginTimestamps = 801

	// PermissionGetAccountDetails is the permission required to read account details
	PermissionGetAccountDetails = 802
)
