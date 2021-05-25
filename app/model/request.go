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