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
	"backend/src/models"
	"backend/src/schema"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ProductRequest struct {
	Id    int64
	Name  string
	Image sql.NullString
	Price float64
}

type ProductCartRequest struct {
	Id    int64
	Name  string
	Image sql.NullString
	Price float64
	Qty   uint16
	Total float64
}

type ProductReviewRequest struct {
	Id          int64
	Name        string
	Description string
	Rating      float64
	Percentage  float64
	CreatedAt   time.Time
}

func OrderWishlist(c *gin.Context) {

	auth := c.MustGet("claims").(jwt.MapClaims)
	product_id := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	var user models.User
	if err := db.Where("id = ?", auth["id"]).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	db.Exec("DELETE FROM products_wishlists WHERE product_id = ? AND = ?", product_id, auth["id"])
	db.Exec("INSERT INTO products_wishlists(product_id, user_id) VALUES(?,?) ", product_id, auth["id"])

	Activity := models.Activity{
		UserId:      int64(user.Id),
		Subject:     "Add Wishlist",
		Event:       "Add Product To Wishlist",
		Description: "Your has been added product to your wishlist.",
	}
	db.Create(&Activity)

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "ok"})
}

func OrderGetSession(c *gin.Context) {

	auth := c.MustGet("claims").(jwt.MapClaims)
	db := c.MustGet("db").(*gorm.DB)

	var order models.Order
	db.Where("status = 0 AND user_id = ?", auth["id"]).Order("id desc").First(&order)

	var carts []ProductCartRequest
	var whislists []ProductRequest

	db.Raw(`
		SELECT 
			products.image,
			products.name,
			products.price
		FROM
			products
		INNER JOIN products_wishlists ON products_wishlists.product_id = products.id
		WHERE products_wishlists.user_id = ?
	`, auth["id"]).Scan(&whislists)

	db.Raw(`
		SELECT 
			products.id,
			products.image,
			products.name,
			products.price,
			orders_details.qty,
			orders_details.total
		FROM orders_details
		INNER JOIN products_inventories ON products_inventories.id = orders_details.inventory_id
		INNER JOIN orders ON orders.id = orders_details.order_id
		INNER JOIN products ON products.id = products_inventories.product_id
		WHERE orders.status = 0 AND orders.user_id = ?
	`, auth["id"]).Scan(&carts)

	var payload = gin.H{
		"carts":     carts,
		"order":     order,
		"whislists": whislists,
	}

	c.JSON(http.StatusOK, payload)
}

func OrderCart(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	auth := c.MustGet("claims").(jwt.MapClaims)
	product_id := c.Param("id")

	var user models.User
	if err := db.Where("id = ?", auth["id"]).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var topProduct models.Product
	db.Where("status = 1 AND published_at <= NOW()").Order("total_rating desc").First(&topProduct)

	var getProducts []models.Product
	db.Preload("Categories").Limit(4).Where("status = 1 AND published_at <= NOW() AND id = ?", product_id).Order("id desc").Find(&getProducts)

	var getTopSellings []models.Product
	db.Preload("Categories").Limit(3).Where("status = 1 AND published_at <= NOW() AND id != ? ", product_id).Order("total_order desc").Find(&getTopSellings)

	var images []models.ProductImage
	db.Where("product_id = ? ", product_id).Order("id desc").Find(&images)

	var sizes []models.Size
	db.Where("status = 1").Order("name asc").Find(&sizes)

	var colours []models.Size
	db.Where("status = 1").Order("name asc").Find(&colours)

	var inventories []models.ProductInventory
	db.Where("product_id = ?", product_id).Find(&inventories)

	var products []ProductResponse
	for _, p := range getProducts {

		var categoryNames []string
		for _, cat := range p.Categories {
			categoryNames = append(categoryNames, cat.Name)
		}

		numRandom := uint16(helpers.RandomInt(0, 1))
		products = append(products, ProductResponse{
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

	var topSellings []ProductResponse
	for _, p3 := range getTopSellings {

		var categoryNames []string
		for _, cat := range p3.Categories {
			categoryNames = append(categoryNames, cat.Name)
		}

		numRandom := uint16(helpers.RandomInt(0, 1))
		topSellings = append(topSellings, ProductResponse{
			Id:           int64(p3.Id),
			Name:         p3.Name,
			Image:        p3.Image,
			Description:  p3.Description,
			Details:      p3.Details,
			Price:        p3.Price,
			PriceOld:     p3.Price + (p3.Price * 0.05),
			CategoryName: strings.Join(categoryNames, ", "),
			IsNewest:     numRandom == 1,
			IsDiscount:   numRandom == 0,
			TotalRating:  math.Floor((((float64(p3.TotalRating) / float64(topProduct.TotalRating)) * 100) / 20)),
		})
	}

	var payload = gin.H{
		"images":         images,
		"product":        products[0],
		"productRelated": topSellings,
		"sizes":          sizes,
		"colours":        colours,
		"inventories":    inventories,
		"user":           user,
	}

	c.JSON(http.StatusOK, payload)
}

func OrderListReview(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	product_id := c.Param("id")

	var topRating models.ProductReview
	db.Where("product_id = ?", product_id).Order("rating desc").First(&topRating)

	var reviews []models.ProductReview
	db.Preload("User").Where("product_id = ?", product_id).Order("id desc").Find(&reviews)

	var payload []ProductReviewRequest
	for _, p := range reviews {
		payload = append(payload, ProductReviewRequest{
			Id:          int64(p.Id),
			Name:        p.User.FirstName.String + " " + p.User.LastName.String,
			Description: p.Review,
			Rating:      math.Floor((((float64(p.Rating) / float64(topRating.Rating)) * 100) / 20)),
			Percentage:  math.Floor(((float64(p.Rating) / float64(topRating.Rating)) * 100)),
			CreatedAt:   p.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, payload)
}

func OrderCreateReview(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	product_id := c.Param("id")
	auth := c.MustGet("claims").(jwt.MapClaims)

	var user models.User
	if err := db.Where("id = ?", auth["id"]).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var input schema.ReviewSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var Product models.Product
	db.Where("id = ? ", product_id).Order("total_rating desc").First(&Product)

	ModelReview := models.ProductReview{
		User:    user,
		Product: Product,
		Rating:  uint16(input.Rating),
		Status:  1,
		Review:  input.Review,
	}
	db.Create(&ModelReview)

	Activity := models.Activity{
		User:        user,
		Subject:     "Create new review",
		Event:       "Add review to " + Product.Name,
		Description: "Your has been added new review to " + Product.Name,
	}
	db.Create(&Activity)

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "ok"})
}

func OrderCreateCart(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	product_id := c.Param("id")
	auth := c.MustGet("claims").(jwt.MapClaims)

	var input schema.CreateCartSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var Product models.Product
	db.Where("id = ? ", product_id).Order("total_rating desc").First(&Product)

	var User models.User
	db.Where("id = ? ", auth["id"]).First(&User)

	var Payment models.Payment
	db.Where("status = 1").First(&Payment)

	Order := &models.Order{}
	resultOrder := db.Where("status = 0 AND user_id = ?", auth["id"]).Order("id desc").First(Order)

	Inventory := &models.ProductInventory{}
	resultInventory := db.Where("product_id = ? AND size_id = ? AND colour_id = ?", product_id, input.SizeId, input.ColourId).Order("id desc").First(Inventory)

	Total := Product.Price * float64(input.Qty)

	if errors.Is(resultOrder.Error, gorm.ErrRecordNotFound) {
		// Not found
		NewOrder := models.Order{
			User:          User,
			Payment:       Payment,
			InvoiceNumber: strconv.FormatInt(helpers.NowTicks(), 10),
			TotalItem:     uint16(input.Qty),
			Status:        0,
			Subtotal:      Total,
			TotalPaid:     Total,
		}
		db.Create(&NewOrder)
		Order.Id = NewOrder.Id
	} else {
		// Item Exists
		Order.TotalItem = Order.TotalItem + uint16(input.Qty)
		Order.Subtotal = Order.Subtotal + Total
		Order.TotalPaid = Order.Subtotal + Total
		if err := db.Save(&Order).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update Order"})
			return
		}
	}

	if !errors.Is(resultInventory.Error, gorm.ErrRecordNotFound) {

		DetailOrder := &models.OrderDetail{}
		resultDetailOrder := db.Where("inventory_id = ? AND order_id = ?", Inventory.Id, Order.Id).First(DetailOrder)

		if !errors.Is(resultDetailOrder.Error, gorm.ErrRecordNotFound) {
			DetailOrder.Qty = DetailOrder.Qty + uint16(input.Qty)
			DetailOrder.Total = DetailOrder.Total + Total
			if err := db.Save(&DetailOrder).Error; err != nil {
				c.JSON(500, gin.H{"error": "Failed to update Detail Order"})
				return
			}
		} else {
			NewDetailOrder := models.OrderDetail{
				OrderId:     Order.Id,
				InventoryId: Inventory.Id,
				Price:       Product.Price,
				Status:      1,
				Qty:         uint16(input.Qty),
				Total:       Total,
			}
			db.Create(&NewDetailOrder)
		}
	}

	if Order != nil {
		db.Exec("DELETE FROM orders_carts WHERE order_id = ? AND product_id = ?", Order.Id, product_id)
		db.Exec("INSERT INTO orders_carts(order_id, product_id) VALUES (?,?)", Order.Id, product_id)
	}

	Activity := models.Activity{
		User:        User,
		Subject:     "Add Cart",
		Event:       "Add product to cart with name " + Product.Name,
		Description: "Your has been added product to cart. with name " + Product.Name,
	}
	db.Create(&Activity)

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "ok"})
}

func OrderCheckoutInitial(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	auth := c.MustGet("claims").(jwt.MapClaims)

	var user models.User
	if err := db.Where("id = ?", auth["id"]).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var payments []models.Payment
	db.Where("status = 1").Order("name asc").Find(&payments)

	var order models.Order
	db.Where("status = 0 AND user_id = ?", auth["id"]).Order("id desc").First(&order)

	var carts []ProductCartRequest
	db.Raw(`
		SELECT 
			products.id,
			products.image,
			products.name,
			products.price,
			orders_details.qty,
			orders_details.total
		FROM orders_details
		INNER JOIN products_inventories ON products_inventories.id = orders_details.inventory_id
		INNER JOIN orders ON orders.id = orders_details.order_id
		INNER JOIN products ON products.id = products_inventories.product_id
		WHERE orders.status = 0 AND orders.user_id = ?
	`, auth["id"]).Scan(&carts)

	var discount models.Setting
	db.Where("key_name = ?", "discount_value").Order("id desc").First(&discount)

	var taxes models.Setting
	db.Where("key_name = ?", "taxes_value").Order("id desc").First(&taxes)

	var shipment models.Setting
	db.Where("key_name = ?", "total_shipment").Order("id desc").First(&shipment)

	subtotal := order.Subtotal

	iDiscount, err := strconv.ParseFloat(discount.KeyValue, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	iTaxes, err := strconv.ParseFloat(taxes.KeyValue, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	totalShipment, err := strconv.ParseFloat(shipment.KeyValue, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	totalDiscount := subtotal * (iDiscount / 100)
	totalTaxes := subtotal * (iTaxes / 100)

	order.TotalDiscount = totalDiscount
	order.TotalTaxes = totalTaxes
	order.TotalShipment = totalShipment
	order.TotalPaid = (subtotal + totalTaxes + totalShipment) - totalDiscount

	var payload = gin.H{
		"order":    order,
		"carts":    carts,
		"user":     user,
		"payments": payments,
		"discount": iDiscount,
		"taxes":    iTaxes,
		"shipment": totalShipment,
	}

	c.JSON(http.StatusOK, payload)
}

func OrderCheckout(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	auth := c.MustGet("claims").(jwt.MapClaims)

	var input schema.CheckoutSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := db.Where("id = ?", auth["id"]).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	var order models.Order
	db.Where("status = 0 AND user_id = ?", auth["id"]).Order("id desc").First(&order)

	var details []models.OrderDetail
	db.Where("order_id = ? ", order.Id).Find(&details)

	for _, detail := range details {

		inventory := detail.Inventory
		inventory.Stock = inventory.Stock - detail.Qty
		if err := db.Save(&inventory).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update inventory"})
			return
		}

		product := detail.Inventory.Product
		product.TotalOrder = product.TotalOrder + detail.Qty
		if err := db.Save(&inventory).Error; err != nil {
			c.JSON(500, gin.H{"error": "Failed to update product"})
			return
		}

		db.Exec("DELETE FROM products_wishlists WHERE product_id = ? AND = ?", product.Id, auth["id"])

	}

	var discount models.Setting
	db.Where("key_name = ?", "discount_value").Order("id desc").First(&discount)

	var taxes models.Setting
	db.Where("key_name = ?", "taxes_value").Order("id desc").First(&taxes)

	var shipment models.Setting
	db.Where("key_name = ?", "total_shipment").Order("id desc").First(&shipment)

	subtotal := order.Subtotal

	iDiscount, err := strconv.ParseFloat(discount.KeyValue, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	iTaxes, err := strconv.ParseFloat(taxes.KeyValue, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	totalShipment, err := strconv.ParseFloat(shipment.KeyValue, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	totalDiscount := subtotal * (iDiscount / 100)
	totalTaxes := subtotal * (iTaxes / 100)

	order.Status = 1
	order.PaymentId = input.PaymentId
	order.TotalTaxes = totalTaxes
	order.TotalDiscount = totalDiscount
	order.TotalShipment = totalShipment
	order.TotalPaid = (subtotal + totalTaxes + totalShipment) - totalDiscount
	if err := db.Save(&order).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update order"})
		return
	}

	db.Create(&models.OrderBilling{
		OrderId:     order.Id,
		Name:        "email",
		Description: input.Email,
		Status:      1,
	})

	db.Create(&models.OrderBilling{
		OrderId:     order.Id,
		Name:        "phone",
		Description: input.Phone,
		Status:      1,
	})

	db.Create(&models.OrderBilling{
		OrderId:     order.Id,
		Name:        "first_name",
		Description: input.FirstName,
		Status:      1,
	})

	db.Create(&models.OrderBilling{
		OrderId:     order.Id,
		Name:        "last_name",
		Description: input.LastName,
		Status:      1,
	})

	db.Create(&models.OrderBilling{
		OrderId:     order.Id,
		Name:        "country",
		Description: input.Country,
		Status:      1,
	})

	db.Create(&models.OrderBilling{
		OrderId:     order.Id,
		Name:        "city",
		Description: input.City,
		Status:      1,
	})

	db.Create(&models.OrderBilling{
		OrderId:     order.Id,
		Name:        "zip_code",
		Description: input.ZipCode,
		Status:      1,
	})

	db.Create(&models.OrderBilling{
		OrderId:     order.Id,
		Name:        "address",
		Description: input.Address,
		Status:      1,
	})

	db.Create(&models.OrderBilling{
		OrderId:     order.Id,
		Name:        "notes",
		Description: input.Notes,
		Status:      1,
	})

	db.Exec("DELETE FROM orders_carts WHERE order_id = ? ", order.Id)

	Activity := models.Activity{
		UserId:      int64(user.Id),
		Subject:     "Checkout Order",
		Event:       "Completed Checkout Current Order",
		Description: "Your order has been finished.",
	}
	db.Create(&Activity)

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "ok"})
}

func OrderList(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	auth := c.MustGet("claims").(jwt.MapClaims)

	page := 1
	limit := 10
	order_by := "id"
	order_dir := "desc"
	offset := ((page - 1) * limit)
	var data []models.Order
	var total_filtered int64
	var total_all int64

	db.Model(&models.Order{}).Where("user_id = ? ", auth["id"]).Count(&total_all)
	total_filtered = total_all

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
		db = db.Where("invoice_number LIKE ? ", c.Query("search"))
		db.Model(&models.Order{}).Where("invoice_number LIKE ? ", c.Query("search")).Where("user_id = ? ", auth["id"]).Count(&total_filtered)
	}

	db = db.Limit(limit).Offset(offset).Order(order_by + " " + order_dir).Find(&data)

	var payload = gin.H{
		"list":          data,
		"totalAll":      total_all,
		"totalFiltered": total_filtered,
		"limit":         limit,
		"page":          page,
	}

	c.JSON(http.StatusOK, payload)
}

func OrderDetail(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var order models.Order
	db.Where("id = ?", id).Order("id desc").First(&order)

	var billings []models.OrderBilling
	db.Where("order_id = ? ", id).Order("name asc").Find(&billings)

	var payment models.Payment
	db.Where("id = ?", order.PaymentId).Order("id desc").First(&payment)

	var carts []ProductCartRequest
	db.Raw(`
		SELECT 
			products.id,
			products.image,
			products.name,
			products.price,
			orders_details.qty,
			orders_details.total
		FROM orders_details
		INNER JOIN products_inventories ON products_inventories.id = orders_details.inventory_id
		INNER JOIN orders ON orders.id = orders_details.order_id
		INNER JOIN products ON products.id = products_inventories.product_id
		WHERE orders_details.order_id = ?
	`, id).Scan(&carts)

	discount := (order.TotalDiscount / order.Subtotal) * 100
	taxes := (order.TotalTaxes / order.Subtotal) * 100

	var payload = gin.H{
		"discount": discount,
		"taxes":    taxes,
		"shipment": order.TotalShipment,
		"carts":    carts,
		"order":    order,
		"billings": billings,
		"payment":  payment,
	}

	c.JSON(http.StatusOK, payload)
}

func OrderCancel(c *gin.Context) {

	auth := c.MustGet("claims").(jwt.MapClaims)
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var User models.User
	db.Where("id = ? ", auth["id"]).First(&User)

	db.Exec("DELETE FROM orders_details WHERE order_id = ?", id)
	db.Exec("DELETE FROM orders_carts WHERE order_id = ?", id)
	db.Exec("DELETE FROM orders_billings WHERE order_id = ?", id)
	db.Exec("DELETE FROM orders WHERE id = ?", id)

	Activity := models.Activity{
		UserId:      int64(User.Id),
		Subject:     "Cancel Order",
		Event:       "Canceling Current Order",
		Description: "Your has been canceling current order.",
	}
	db.Create(&Activity)

	c.JSON(http.StatusOK, gin.H{"status": true, "message": "ok"})
}
