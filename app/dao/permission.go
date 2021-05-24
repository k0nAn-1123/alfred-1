// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"alfred/app/dao/internal"
)

// permissionDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type permissionDao struct {
	internal.PermissionDao
}

var (
	// Permission is globally public accessible object for table permission operations.
	Permission = permissionDao{
		internal.Permission,
	}
)

// Fill with you ideas below.
