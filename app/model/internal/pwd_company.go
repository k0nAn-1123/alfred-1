// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

// PwdCompany is the golang structure for table pwd_company.
type PwdCompany struct {
	ID          int    `orm:"ID,primary"  json:"iD"`          //
	CompanyName string `orm:"CompanyName" json:"companyName"` //
	Domain      string `orm:"Domain"      json:"domain"`      //
}
