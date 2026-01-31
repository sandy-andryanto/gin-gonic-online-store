/**
 * This file is part of the Sandy Andryanto Online Store Website.
 *
 * @author     Sandy Andryanto <sandy.andryanto.official@gmail.com>
 * @copyright  2025
 *
 * For the full copyright and license information,
 * please view the LICENSE.md file that was distributed
 * with this source code.
 */

package controllers

import (
	helpers "backend/src/helpers"
	models "backend/src/models"
	schema "backend/src/schema"
	"database/sql"
	"math"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ProductResponse struct {
	Id           int64
	Name         string
	Image        sql.NullString
	Description  string
	Details      string
	Price        float64
	PriceOld     float64
	CategoryName string
	IsNewest     bool
	IsDiscount   bool
	TotalRating  float64
}

func HomePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Connected Established !!"})
}

func HomeComponent(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var categories []models.Category
	db.Where("status = 1").Order("name asc").Find(&categories)

	var settings []models.Setting
	db.Find(&settings)

	setting := make(map[string]string)
	for _, row := range settings {
		setting[row.KeyName] = row.KeyValue
	}

	var payload = gin.H{
		"categories": categories,
		"setting":    setting,
	}

	c.JSON(http.StatusOK, payload)
}

func HomePage(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var topProduct models.Product
	db.Where("status = 1 AND published_at <= NOW()").Order("total_rating desc").First(&topProduct)

	var categories []models.Category
	db.Limit(3).Where("status = 1 AND displayed = 1").Order("name asc").Find(&categories)

	var getProducts []models.Product
	db.Preload("Categories").Limit(4).Where("status = 1 AND published_at <= NOW()").Order("id desc").Find(&getProducts)

	var getBestSellers []models.Product
	db.Preload("Categories").Limit(3).Where("status = 1 AND published_at <= NOW()").Order("total_rating desc").Find(&getBestSellers)

	var getTopSellings []models.Product
	db.Preload("Categories").Limit(3).Where("status = 1 AND published_at <= NOW()").Order("total_order desc").Find(&getTopSellings)

	var products []ProductResponse
	for _, p := range getProducts {
		numRandom := uint16(helpers.RandomInt(0, 1))
		products = append(products, ProductResponse{
			Id:           int64(p.Id),
			Name:         p.Name,
			Image:        p.Image,
			Description:  p.Description,
			Details:      p.Details,
			Price:        p.Price,
			PriceOld:     p.Price + (p.Price * 0.05),
			CategoryName: p.Categories[0].Name,
			IsNewest:     numRandom == 1,
			IsDiscount:   numRandom == 0,
			TotalRating:  math.Floor((((float64(p.TotalRating) / float64(topProduct.TotalRating)) * 100) / 20)),
		})
	}

	var bestSellers []ProductResponse
	for _, p2 := range getBestSellers {
		numRandom := uint16(helpers.RandomInt(0, 1))
		bestSellers = append(bestSellers, ProductResponse{
			Id:           int64(p2.Id),
			Name:         p2.Name,
			Image:        p2.Image,
			Description:  p2.Description,
			Details:      p2.Details,
			Price:        p2.Price,
			PriceOld:     p2.Price + (p2.Price * 0.05),
			CategoryName: p2.Categories[0].Name,
			IsNewest:     numRandom == 1,
			IsDiscount:   numRandom == 0,
			TotalRating:  math.Floor((((float64(p2.TotalRating) / float64(topProduct.TotalRating)) * 100) / 20)),
		})
	}

	var topSellings []ProductResponse
	for _, p3 := range getTopSellings {
		numRandom := uint16(helpers.RandomInt(0, 1))
		topSellings = append(topSellings, ProductResponse{
			Id:           int64(p3.Id),
			Name:         p3.Name,
			Image:        p3.Image,
			Description:  p3.Description,
			Details:      p3.Details,
			Price:        p3.Price,
			PriceOld:     p3.Price + (p3.Price * 0.05),
			CategoryName: p3.Categories[0].Name,
			IsNewest:     numRandom == 1,
			IsDiscount:   numRandom == 0,
			TotalRating:  math.Floor((((float64(p3.TotalRating) / float64(topProduct.TotalRating)) * 100) / 20)),
		})
	}

	var payload = gin.H{
		"categories":  categories,
		"products":    products,
		"bestSellers": bestSellers,
		"topSellings": topSellings,
	}

	c.JSON(http.StatusOK, payload)
}

func HomeNewsletter(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var input schema.UserForgotSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(strings.TrimSpace(input.Email)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The email field is required.!"})
		return
	}

	model := models.NewsLetter{
		Email:     input.Email,
		Status:    1,
		IpAddress: c.ClientIP(),
	}
	db.Create(&model)

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "Message sent successfully!"})
}
