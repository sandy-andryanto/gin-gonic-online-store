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

type Order struct {
	Id            uint64    `json:"id" gorm:"primary_key"`
	UserId        uint64    `json:"user_id" gorm:"index;not null"`
	User          User      `json:"-" gorm:"foreignKey:user_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	PaymentId     uint64    `json:"payment_id" gorm:"index;not null"`
	Payment       Payment   `json:"-" gorm:"foreignKey:payment_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	InvoiceNumber string    `json:"invoice_number" gorm:"index;size:255;not null"`
	TotalItem     uint16    `json:"total_item" gorm:"index;default:0"`
	Subtotal      float64   `json:"subtotal" gorm:"type:decimal(18,4);default:0;index"`
	TotalDiscount float64   `json:"total_discount" gorm:"type:decimal(18,4);default:0;index"`
	TotalTaxes    float64   `json:"total_taxes" gorm:"type:decimal(18,4);default:0;index"`
	TotalShipment float64   `json:"total_shipment" gorm:"type:decimal(18,4);default:0;index"`
	TotalPaid     float64   `json:"total_paid" gorm:"type:decimal(18,4);default:0;index"`
	Status        uint8     `json:"status" gorm:"index;default:0"`
	CreatedAt     time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Products      []Product `gorm:"many2many:orders_carts"`
	Billings      []OrderBilling
	Details       []OrderDetail
}

func (Order) TableName() string {
	return "orders"
}
