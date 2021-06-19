package model

type AuthReq struct {
	Identification string
	Password string
}

type HealthReq struct {
	Height 		int
	Weight 		int
	FatRatio 	float64
	MuscleRatio float64
	UserID 		uint
}

type PasswordReq struct {
	ID			int
	Account		string
	Password	string
	Level		int
	WebName		string
	Domain		string
	Company 	string
	Remark		string
}

type PassListReq struct {
	Page	int
	Count 	int
	Content string
}