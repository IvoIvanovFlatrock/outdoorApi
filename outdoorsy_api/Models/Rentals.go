package Models

import "time"

// gorm.Model definition
type Rental struct {
	Id                uint      `gorm:"primaryKey"`
	User_id           uint      `json:"user_id"`
	User              Users     `json:"user"`
	Name              string    `json:"name"`
	Type_car          string    `json:"type_car"`
	Description       string    `json:"description"`
	Sleeps            int       `json:"sleeps"`
	Price_per_day     int       `json:"price_per_day"`
	Home_city         string    `json:"home_city"`
	Home_state        string    `json:"home_state"`
	Home_zip          string    `json:"home_zip"`
	Home_country      string    `json:"home_country"`
	Vehicle_make      string    `json:"vehicle_make"`
	Vehicle_model     string    `json:"vehicle_model"`
	Vehicle_year      int       `json:"vehicle_year"`
	Vehicle_length    float64   `json:"vehicle_length"`
	Created           time.Time `json:"created"`
	Updated           time.Time `json:"updated"`
	Lat               float64   `json:"lat"`
	Lng               float64   `json:"lng"`
	Primary_image_url string    `json:"primary_image_url"`
}
