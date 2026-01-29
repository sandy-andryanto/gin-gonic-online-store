/**
 * This file is part of the Sandy Andryanto Online Store Website.
 *
 * @author     Sandy Andryanto <sandy.andryanto.blade@gmail.com>
 * @copyright  2025
 *
 * For the full copyright and license information,
 * please view the LICENSE.md file that was distributed
 * with this source code.
 */

package data

import (
	_db "backend/src/config"
	"backend/src/helpers"
	"backend/src/models"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	math "math/rand"
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
)

func RunSeed() {
	CreateUser()
	CreateSetting()
	CreateCategories()
	CreateBrands()
	CreateColours()
	CreatePayment()
	CreateSize()
	CreateProduct()
}

func CreateSetting() {

	var totalRow int64

	db := _db.SetupDB()
	db.Model(&models.Setting{}).Where("id <> 0").Count(&totalRow)

	if totalRow == 0 {

		settings := map[string]string{
			"about_section":   "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut.",
			"com_location":    "West Java, Indonesia",
			"com_phone":       "+62-898-921-8470",
			"com_email":       "sandy.andryanto.blade@gmail.com",
			"com_currency":    "USD",
			"installed":       "1",
			"discount_active": "1",
			"discount_value":  "5",
			"discount_start":  time.Now().Format("2006-01-02 15:04:05"),
			"discount_end":    time.Now().Add(7 * 24 * time.Hour).Format("2006-01-02 15:04:05"),
			"taxes_value":     "10",
			"total_shipment":  "50",
		}

		for key, value := range settings {
			setting := models.Setting{
				KeyName:  key,
				KeyValue: value,
				Status:   1,
			}
			db.Create(&setting)
		}

	}

}

func CreateCategories() {

	var totalRow int64

	db := _db.SetupDB()
	db.Model(&models.Category{}).Where("id <> 0").Count(&totalRow)

	if totalRow == 0 {

		items := map[string]string{
			"Laptop":      "https://5an9y4lf0n50.github.io/demo-images/demo-commerce/product01.png",
			"Smartphone":  "https://5an9y4lf0n50.github.io/demo-images/demo-commerce/product02.png",
			"Camera":      "https://5an9y4lf0n50.github.io/demo-images/demo-commerce/product03.png",
			"Accessories": "https://5an9y4lf0n50.github.io/demo-images/demo-commerce/product04.png",
			"Others":      "https://5an9y4lf0n50.github.io/demo-images/demo-commerce/product05.png",
		}

		i := 1
		for name, image := range items {

			var displayed int16

			if i < 4 {
				displayed = 1
			} else {
				displayed = 0
			}

			category := models.Category{
				Name:        name,
				Image:       sql.NullString{String: image, Valid: true},
				Description: randomdata.Paragraph(),
				Displayed:   uint8(displayed),
				Status:      1,
			}
			db.Create(&category)
			i++
		}
	}

}

func CreateBrands() {

	var totalRow int64

	db := _db.SetupDB()
	db.Model(&models.Brand{}).Where("id <> 0").Count(&totalRow)

	if totalRow == 0 {

		items := []string{"Samsung", "LG", "Sony", "Apple", "Microsoft"}

		for _, item := range items {
			brand := models.Brand{
				Name:        item,
				Description: randomdata.Paragraph(),
				Status:      1,
			}
			db.Create(&brand)
		}
	}

}

func CreateColours() {

	var totalRow int64

	db := _db.SetupDB()
	db.Model(&models.Colour{}).Where("id <> 0").Count(&totalRow)

	if totalRow == 0 {

		colors := map[string]string{
			"#FF0000": "Red",
			"#0000FF": "Blue",
			"#FFFF00": "Yellow",
			"#000000": "Black",
			"#FFFFFF": "White",
			"#666":    "Dark Gray",
			"#AAA":    "Light Gray",
		}

		for name, code := range colors {
			colour := models.Colour{
				Code:        code,
				Name:        name,
				Description: randomdata.Paragraph(),
				Status:      1,
			}
			db.Create(&colour)
		}

	}

}

func CreatePayment() {

	var totalRow int64

	db := _db.SetupDB()
	db.Model(&models.Payment{}).Where("id <> 0").Count(&totalRow)

	if totalRow == 0 {

		items := []string{"Direct Bank Transfer", "Cheque Payment", "Paypal System"}

		for _, item := range items {
			payment := models.Payment{
				Name:        item,
				Description: randomdata.Paragraph(),
				Status:      1,
			}
			db.Create(&payment)
		}
	}

}

func CreateSize() {

	var totalRow int64

	db := _db.SetupDB()
	db.Model(&models.Size{}).Where("id <> 0").Count(&totalRow)

	if totalRow == 0 {

		items := []string{
			"11 to 12 Inches",
			"13 to 14 Inches",
			"15 to 16 Inches",
			"17 to 18 Inches",
		}

		for _, item := range items {
			size := models.Size{
				Name:        item,
				Description: randomdata.Paragraph(),
				Status:      1,
			}
			db.Create(&size)
		}
	}

}

func CreateProduct() {

	var totalRow int64

	db := _db.SetupDB()
	db.Model(&models.Product{}).Where("id <> 0").Count(&totalRow)

	var sizes []models.Size
	db.Find(&sizes)

	var colours []models.Colour
	db.Find(&colours)

	if totalRow == 0 {

		images := []string{
			"https://5an9y4lf0n50.github.io/demo-images/demo-commerce/product01.png",
			"https://5an9y4lf0n50.github.io/demo-images/demo-commerce/product02.png",
			"https://5an9y4lf0n50.github.io/demo-images/demo-commerce/product03.png",
			"https://5an9y4lf0n50.github.io/demo-images/demo-commerce/product04.png",
			"https://5an9y4lf0n50.github.io/demo-images/demo-commerce/product05.png",
			"https://5an9y4lf0n50.github.io/demo-images/demo-commerce/product06.png",
			"https://5an9y4lf0n50.github.io/demo-images/demo-commerce/product07.png",
			"https://5an9y4lf0n50.github.io/demo-images/demo-commerce/product08.png",
			"https://5an9y4lf0n50.github.io/demo-images/demo-commerce/product09.png",
		}

		for i := 1; i <= 9; i++ {

			math.Seed(time.Now().UnixNano())
			image := images[math.Intn(len(images))]

			var categories []models.Category
			db.Order("RAND()").Limit(3).Find(&categories)

			var reviewers []models.User
			db.Order("RAND()").Limit(5).Find(&reviewers)

			var brand models.Brand
			db.Order("RAND()").Limit(1).First(&brand)

			product := models.Product{
				Image:       sql.NullString{String: image, Valid: true},
				BrandId:     brand.Id,
				Sku:         fmt.Sprintf("P%03d", i),
				Name:        fmt.Sprintf("Product %03d", i),
				Price:       float64(helpers.RandomInt(100, 999)),
				TotalOrder:  uint16(helpers.RandomInt(100, 1000)),
				TotalRating: uint16(helpers.RandomInt(100, 1000)),
				Description: randomdata.Paragraph(),
				Details:     randomdata.Paragraph(),
				PublishedAt: func(t time.Time) *time.Time { return &t }(time.Now()),
				Categories:  categories,
				Status:      1,
			}
			db.Create(&product)

			for _, reviewer := range reviewers {
				rr := models.ProductReview{
					ProductId: product.Id,
					UserId:    reviewer.Id,
					Rating:    uint16(helpers.RandomInt(0, 100)),
					Review:    randomdata.Paragraph(),
					Status:    1,
				}
				db.Create(&rr)
			}

			for j := 1; j <= 3; j++ {
				math.Seed(time.Now().UnixNano())
				imageOther := images[math.Intn(len(images))]
				pi := models.ProductImage{
					ProductId: product.Id,
					Path:      imageOther,
					Sort:      uint16(j),
					Status:    1,
				}
				db.Create(&pi)
			}

			for _, size := range sizes {
				for _, colour := range colours {
					inv := models.ProductInventory{
						ProductId: product.Id,
						SizeId:    size.Id,
						ColourId:  colour.Id,
						Stock:     uint16(helpers.RandomInt(1, 50)),
						Status:    1,
					}
					db.Create(&inv)
				}
			}

		}
	}

}

func CreateUser() {

	var totalRow int64

	db := _db.SetupDB()
	db.Model(&models.User{}).Where("id <> 0").Count(&totalRow)

	if totalRow == 0 {
		for i := 1; i <= 10; i++ {

			bytes := make([]byte, 32)
			if _, err := rand.Read(bytes); err != nil {
				panic(err.Error())
			}

			key := hex.EncodeToString(bytes)
			encrypted := helpers.Encrypt("Qwerty123!", key)

			min := 1
			max := 2
			gender := math.Intn(max-min+1) + min
			firstName := ""
			genderChar := ""

			if gender == 1 {
				genderChar = "M"
				firstName = faker.FirstNameMale()
			} else {
				genderChar = "F"
				firstName = faker.FirstNameFemale()
			}

			user := models.User{
				Email:     randomdata.Email(),
				Phone:     randomdata.PhoneNumber(),
				FirstName: sql.NullString{String: firstName, Valid: true},
				LastName:  sql.NullString{String: faker.LastName(), Valid: true},
				Gender:    sql.NullString{String: genderChar, Valid: true},
				Country:   sql.NullString{String: randomdata.Country(randomdata.FullCountry), Valid: true},
				City:      sql.NullString{String: randomdata.City(), Valid: true},
				Address:   sql.NullString{String: randomdata.Address(), Valid: true},
				Salt:      key,
				Password:  encrypted,
				Status:    1,
			}
			db.Create(&user)

			token := uuid.New().String()
			auth := models.Authentication{
				UserId:     int64(user.Id),
				AuthType:   "email-confirm",
				Credential: user.Email,
				Token:      token,
				Status:     1,
			}
			db.Create(&auth)

		}
	}

}
