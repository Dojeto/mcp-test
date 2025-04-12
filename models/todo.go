package models

type Todo struct {
	ID          string `bson:"_id,omitempty" json:"id"`
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	Completed   bool   `bson:"completed" json:"completed"`
}