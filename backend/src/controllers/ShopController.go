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

package controllers

import (
	helpers "backend/src/helpers"
	models "backend/src/models"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ResponseCount struct {
	Id    uint
	Name  string
	Total int64
}

func ShopFilter(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var topPrice models.Product
	db.Where("status = 1 AND published_at <= NOW()").Order("price desc").First(&topPrice)

	var minPrice models.Product
	db.Where("status = 1 AND published_at <= NOW()").Order("price asc").First(&minPrice)

	var getTopSellings []models.Product
	db.Preload("Categories").Limit(3).Where("status = 1 AND published_at <= NOW()").Order("total_order desc").Find(&getTopSellings)

	var topProduct models.Product
	db.Where("status = 1 AND published_at <= NOW()").Order("total_rating desc").First(&topProduct)

	var categories []ResponseCount
	var brands []ResponseCount

	db.Raw(`
		SELECT 
			categories.id, 
			categories.name, 
			COUNT(*) AS total
		FROM categories
		INNER JOIN products_categories ON products_categories.category_id = categories.id
		GROUP BY categories.id, categories.name
	`).Scan(&categories)

	db.Raw(`
		SELECT
			brands.id,
			brands.name,
			COUNT(*) AS total
		FROM 
			brands
		INNER JOIN products ON products.brand_id = brands.id
		GROUP BY
			brands.id,
			brands.name
		ORDER BY brands.name ASC
	`).Scan(&brands)

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
		"categories": categories,
		"brands":     brands,
		"tops":       topSellings,
		"maxPrice":   topPrice.Price,
		"minPrice":   minPrice.Price,
	}

	c.JSON(http.StatusOK, payload)
}

func ShopList(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	page := 1
	limit := 9
	order_by := "id"
	order_dir := "desc"
	offset := ((page - 1) * limit)
	var data []models.Product

	var total_all int64
	db.Model(&models.Product{}).Where("status = 1 AND published_at <= NOW()").Count(&total_all)

	var topProduct models.Product
	db.Where("status = 1 AND published_at <= NOW()").Order("total_rating desc").First(&topProduct)

	db = db.Preload("Categories").Where("status = 1 AND published_at <= NOW()")

	if len(strings.TrimSpace(c.Query("category"))) > 0 {
		category, ok := c.GetQuery("category")
		if ok {
			db = db.
				Joins("JOIN products_categories ON products_categories.product_id = products.id").
				Where("products_categories.category_id IN (?)", strings.Split(category, ",")).
				Group("products.id")
		}
	}

	if len(strings.TrimSpace(c.Query("brand"))) > 0 {
		brand, ok := c.GetQuery("brand")
		if ok {
			db = db.Where("products.brand_id IN (?)", strings.Split(brand, ","))
		}
	}

	if len(strings.TrimSpace(c.Query("priceMax"))) > 0 || len(strings.TrimSpace(c.Query("priceMin"))) > 0 {
		order_by = "price"
	}

	if len(strings.TrimSpace(c.Query("limit"))) > 0 {
		limit_, err := strconv.ParseInt(c.Query("limit"), 0, 32)
		if err == nil {
			limit = int(limit_)
		}
	}

	if len(strings.TrimSpace(c.Query("page"))) > 0 {
		page_, err := strconv.ParseInt(c.Query("page"), 0, 32)
		if err == nil {
			page = int(page_)
		}
	}

	if len(strings.TrimSpace(c.Query("order_by"))) > 0 {
		order_by = c.Query("order_by")
	}

	if len(strings.TrimSpace(c.Query("order_dir"))) > 0 {
		order_dir = c.Query("order_dir")
	}

	if len(strings.TrimSpace(c.Query("search"))) > 0 {
		db = db.Where("name LIKE ? OR sku LIKE ? OR description LIKE ?", c.Query("search"), c.Query("search"), c.Query("search"))
	}

	var total_filtered int64
	db.Count(&total_filtered)

	db = db.Limit(limit).Offset(offset).Order(order_by + " " + order_dir).Find(&data)

	var productResult []ProductResponse
	for _, p := range data {

		var categoryNames []string
		for _, cat := range p.Categories {
			categoryNames = append(categoryNames, cat.Name)
		}

		numRandom := uint16(helpers.RandomInt(0, 1))
		productResult = append(productResult, ProductResponse{
			Id:           int64(p.Id),
			Name:         p.Name,
			Image:        p.Image,
			Description:  p.Description,
			Details:      p.Details,
			Price:        p.Price,
			PriceOld:     p.Price + (p.Price * 0.05),
			CategoryName: strings.Join(categoryNames, ", "),
			IsNewest:     numRandom == 1,
			IsDiscount:   numRandom == 0,
			TotalRating:  math.Floor((((float64(p.TotalRating) / float64(topProduct.TotalRating)) * 100) / 20)),
		})
	}

	var payload = gin.H{
		"list":          productResult,
		"totalAll":      total_all,
		"totalFiltered": total_filtered,
		"limit":         limit,
		"page":          page,
	}

	c.JSON(http.StatusOK, payload)
}
