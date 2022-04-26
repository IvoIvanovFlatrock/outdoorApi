package Models

type RentalResponse struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	Type            string   `json:"type"`
	Make            string   `json:"make"`
	Model           string   `json:"model"`
	Year            string   `json:"year"`
	Length          string   `json:"length"`
	Sleeps          string   `json:"sleeps"`
	PrimaryImageURL string   `json:"primary_image_url"`
	Location        Location `json:"location"`
	User            Users    `json:"user"`
	Price           Price    `json:"price"`
}

type Location struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
	Lat     string `json:"lat"`
	Lng     string `json:"lng"`
}

type Price struct {
	Day string `json:"day"`
}
