package Models

type Users struct {
	Id         int    `gorm:"primaryKey"`
	First_name string `json:"First_name"`
	Last_name  string `json:"Last_name"`
}
