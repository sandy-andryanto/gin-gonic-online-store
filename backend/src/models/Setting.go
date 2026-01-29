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

type Setting struct {
	Id        uint64    `json:"id" gorm:"primary_key"`
	KeyName   string    `json:"key_name" gorm:"index;size:255;not null"`
	KeyValue  string    `json:"key_value"  gorm:"type:longtext;default:null;"`
	Status    uint8     `json:"status" gorm:"index;default:0"`
	CreatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Setting) TableName() string {
	return "settings"
}
