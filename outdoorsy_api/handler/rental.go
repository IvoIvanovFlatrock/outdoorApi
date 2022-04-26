package Handler

import (
	"fmt"
	"outdoorsy_api/Models"
	"outdoorsy_api/database"
	"strconv"
	"strings"

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

	rs := getRentalsWithParams(c)

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

func getRentalsWithParams(c *fiber.Ctx) []Models.RentalUser {
	query := "SELECT * FROM rentals LEFT JOIN users ON rentals.user_id = users.id"

	and := false

	if c.Query("price_min") != "" {
		query = query + " WHERE rentals.Price_per_day > " + c.Query("price_min")
		and = true
	}

	if c.Query("price_max") != "" {
		queryPriceMax := " WHERE rentals.Price_per_day < " + c.Query("price_max")

		if and {
			queryPriceMax = " AND " + "rentals.Price_per_day < " + c.Query("price_max")
		}

		query = query + queryPriceMax
		and = true
	}

	if c.Query("limit") != "" {
		if and {
			query = query + " AND "
		}

		query = query + " LIMIT " + c.Query("limit")
		and = true
	}

	if c.Query("sort") != "" {
		if and {
			query = query + " AND "
		}

		query = query + " SORT BY " + c.Query("sort")
		and = true
	}

	if c.Query("ids") != "" {
		s := strings.Split(c.Query("ids"), ",")

		for _, v := range s {
			if and {
				query = query + " AND rentals.Id = " + v
			} else {
				query = query + " WHERE  rentals.Id" + v
			}
		}

	}

	if c.Query("near") != "" {
		intVar, _ := strconv.Atoi(c.Query("near"))

		max := intVar + 2
		min := intVar - 2
		query = query + "WHERE HEIGHT between" + strconv.Itoa(min) + "and" + strconv.Itoa(max)
	}

	db := database.DB

	var res []Models.RentalUser
	db.Raw(query).Scan(&res)

	return res
}
