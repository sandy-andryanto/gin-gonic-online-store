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

package models

import (
	"time"
)

type ProductImage struct {
	Id        uint64    `json:"id" gorm:"primary_key"`
	ProductId uint64    `json:"product_id" gorm:"index;not null"`
	Product   Product   `gorm:"foreignKey:product_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Path      string    `json:"path" gorm:"index;size:255;not null"`
	Sort      uint16    `json:"sort" gorm:"index;default:0"`
	Status    uint8     `json:"status" gorm:"index;default:0"`
	CreatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (ProductImage) TableName() string {
	return "products_images"
}
