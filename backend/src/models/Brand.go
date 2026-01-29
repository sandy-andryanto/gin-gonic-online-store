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

type Brand struct {
	Id          uint64         `json:"id" gorm:"primary_key"`
	Image       sql.NullString `json:"image" gorm:"index;size:191;default:null;"`
	Name        string         `json:"name" gorm:"index;size:255;not null"`
	Description string         `json:"description"  gorm:"type:text;default null"`
	Status      uint8          `json:"status" gorm:"index;default:0"`
	CreatedAt   time.Time      `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Products    []Product
}

func (Brand) TableName() string {
	return "brands"
}
