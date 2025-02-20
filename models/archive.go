package models

type AculeiImage struct {
	Id              string   `bson:"id" json:"id" example:"d38a0ec061a460466c253efe9a62cb14"`
	ImageName       string   `bson:"image_name" json:"image_name" example:"TF_ACULEI_01062021-2741.jpg"`
	PredictedAnimal string   `bson:"predicted_animal" json:"predicted_animal" example:"fox"`
	MoonPhase       *string  `bson:"moon_phase" json:"moon_phase" example:"Waning Crescent"`
	Temperature     *float64 `bson:"temperature" json:"temperature" example:"12"`
	Date            *string  `bson:"date" json:"date" example:"2021-06-01 22:47:09"`
	Cam             string   `bson:"cam" json:"cam" example:"CAM5"`
}
