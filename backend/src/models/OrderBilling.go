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

type OrderBilling struct {
	Id          uint64    `json:"id" gorm:"primary_key"`
	OrderId     uint64    `json:"order_id" gorm:"index;not null"`
	Order       Order     `gorm:"foreignKey:order_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Name        string    `json:"name" gorm:"index;size:255;not null"`
	Description string    `json:"description"  gorm:"type:text;default null"`
	Status      uint8     `json:"status" gorm:"index;default:0"`
	CreatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (OrderBilling) TableName() string {
	return "orders_billings"
}
