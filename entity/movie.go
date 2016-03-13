package entity

type Movie struct {
	Id string `json:"id" bson:"_id"`
	Name string `json:"name"`
	Category string `json:"category"`
}
