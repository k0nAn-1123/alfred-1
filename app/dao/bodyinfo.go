// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"alfred/app/dao/internal"
)

// bodyinfoDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type bodyinfoDao struct {
	internal.BodyinfoDao
}

var (
	// Bodyinfo is globally public accessible object for table bodyinfo operations.
	Bodyinfo = bodyinfoDao{
		internal.Bodyinfo,
	}
)

// Fill with you ideas below.