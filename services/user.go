package services

type User struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	Username  string `json:"username" bson:"username,omitempty"`
	Email     string `json:"email" bson:"email,omitempty"`
	Password  string `json:"password" bson:"password,omitempty"`
	CreatedAt string `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at" bson:"updated_at,omitempty"`
}
