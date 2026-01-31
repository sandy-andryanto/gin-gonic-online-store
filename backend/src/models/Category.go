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
	"database/sql"
	"time"
)

type Category struct {
	Id          uint64         `json:"id" gorm:"primary_key"`
	Image       sql.NullString `json:"image" gorm:"index;size:191;default:null;"`
	Name        string         `json:"name" gorm:"index;size:255;not null"`
	Description string         `json:"description"  gorm:"type:text;default null"`
	Displayed   uint8          `json:"displayed" gorm:"index;default:0"`
	Status      uint8          `json:"status" gorm:"index;default:0"`
	CreatedAt   time.Time      `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Products    []Product      `gorm:"many2many:products_categories"`
}

func (Category) TableName() string {
	return "categories"
}
