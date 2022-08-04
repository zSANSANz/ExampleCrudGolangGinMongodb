package model

type Devision struct {
	*DevisionInput `bson:",inline"`
}

type DevisionInput struct {
	Name string `json:"name" bson:"name"`
}
