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
	"time"
)

type ProductReview struct {
	Id        uint64    `json:"id" gorm:"primary_key"`
	ProductId uint64    `json:"product_id" gorm:"index;not null"`
	Product   Product   `gorm:"foreignKey:product_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	UserId    uint64    `json:"user_id" gorm:"index;not null"`
	User      User      `gorm:"foreignKey:user_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Rating    uint16    `json:"rating" gorm:"index;default:0"`
	Review    string    `json:"review"  gorm:"type:text;default null"`
	Status    uint8     `json:"status" gorm:"index;default:0"`
	CreatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (ProductReview) TableName() string {
	return "products_reviews"
}
