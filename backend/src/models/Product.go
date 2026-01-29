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

package models

import (
	"database/sql"
	"time"
)

type Product struct {
	Id          uint64         `json:"id" gorm:"primary_key"`
	BrandId     uint64         `json:"brand_id" gorm:"index;not null"`
	Brand       Brand          `json:"-" gorm:"foreignKey:brand_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Image       sql.NullString `json:"image" gorm:"index;size:191;default:null;"`
	Sku         string         `json:"sku" gorm:"index;size:100;not null"`
	Name        string         `json:"name" gorm:"index;size:255;not null"`
	Price       float64        `json:"price" gorm:"type:decimal(18,4);default:0;index"`
	TotalOrder  uint16         `json:"total_order" gorm:"index;default:0"`
	TotalRating uint16         `json:"total_rating" gorm:"index;default:0"`
	Description string         `json:"description"  gorm:"type:text;default null"`
	Details     string         `json:"details"  gorm:"type:text;default null"`
	Status      uint8          `json:"status" gorm:"index;default:0"`
	PublishedAt *time.Time     `json:"published_at" gorm:"index"`
	CreatedAt   time.Time      `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Categories  []Category     `gorm:"many2many:products_categories"`
	Orders      []Order        `gorm:"many2many:orders_carts"`
	Users       []User         `gorm:"many2many:products_wishlists"`
	Images      []ProductImage
	Inventories []ProductInventory
	Reviews     []ProductReview
}

func (Product) TableName() string {
	return "products"
}
