package Handler

import (
	"fmt"
	"outdoorsy_api/Models"
	"outdoorsy_api/database"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetRentals(c *fiber.Ctx) error {
	rs := getRentals()

	var rentals []Models.RentalResponse

	for _, v := range rs {
		var rental Models.RentalResponse
		var location Models.Location
		location.City = v.Home_city
		location.Country = v.Home_country
		location.Lat = fmt.Sprintf("%f", v.Lat)
		location.Lng = fmt.Sprintf("%f", v.Lng)
		location.State = v.Home_state
		location.Zip = v.Home_zip

		var user Models.Users
		user.First_name = v.First_name
		user.Last_name = v.Last_name
		user.Id = v.User_id

		var price Models.Price
		price.Day = strconv.Itoa(v.Price_per_day)

		rental.ID = strconv.Itoa(v.Id)
		rental.Name = v.Name
		rental.Description = v.Description
		rental.Type = v.Type
		rental.Make = v.Vehicle_make
		rental.Model = v.Vehicle_model
		rental.Year = strconv.Itoa(v.Created.Year())
		rental.Length = fmt.Sprintf("%f", v.Vehicle_length)
		rental.Sleeps = strconv.Itoa(v.Sleeps)
		rental.PrimaryImageURL = v.Primary_image_url
		rental.Location = location
		rental.User = user
		rental.Price = price

		rentals = append(rentals, rental)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": rentals})
}

func GetRentalsWithParams(c *fiber.Ctx) error {

	c.AllParams()
	var params []string
	if c.Params("price_min") != "" {
		params = append(params, "price_min")
	}

	if c.Params("price_max") != "" {
		params = append(params, "price_max")
	}

	if c.Params("limit") != "" {
		params = append(params, "limit")
	}

	if c.Params("sort") != "" {
		params = append(params, "sort")
	}

	if c.Params("ids") != "" {
		params = append(params, "ids")
	}

	if c.Params("near") != "" {
		params = append(params, "near")
	}

	rs := getRentalsWithParams(params)

	var rentals []Models.RentalResponse

	for _, v := range rs {
		var rental Models.RentalResponse
		var location Models.Location
		location.City = v.Home_city
		location.Country = v.Home_country
		location.Lat = fmt.Sprintf("%f", v.Lat)
		location.Lng = fmt.Sprintf("%f", v.Lng)
		location.State = v.Home_state
		location.Zip = v.Home_zip

		var user Models.Users
		user.First_name = v.First_name
		user.Last_name = v.Last_name
		user.Id = v.User_id

		var price Models.Price
		price.Day = strconv.Itoa(v.Price_per_day)

		rental.ID = strconv.Itoa(v.Id)
		rental.Name = v.Name
		rental.Description = v.Description
		rental.Type = v.Type
		rental.Make = v.Vehicle_make
		rental.Model = v.Vehicle_model
		rental.Year = strconv.Itoa(v.Created.Year())
		rental.Length = fmt.Sprintf("%f", v.Vehicle_length)
		rental.Sleeps = strconv.Itoa(v.Sleeps)
		rental.PrimaryImageURL = v.Primary_image_url
		rental.Location = location
		rental.User = user
		rental.Price = price

		rentals = append(rentals, rental)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": rentals})
}

func getRentals() []Models.RentalUser {
	db := database.DB

	var res []Models.RentalUser
	db.Raw("SELECT * FROM rentals LEFT JOIN users ON rentals.user_id = users.id;").Scan(&res)

	return res
}

func getRentalsWithParams(params []string) []Models.RentalUser {
	query := "SELECT * FROM rentals LEFT JOIN users ON rentals.user_id = users.id;"

	for _, v := range params {
		if v == "sort" {
			query = query + " SORT BY " + v
		}
	}

	db := database.DB

	var res []Models.RentalUser
	db.Raw("SELECT * FROM rentals LEFT JOIN users ON rentals.user_id = users.id;").Scan(&res)

	return res
}
