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

import "time"

type Activity struct {
	Id          uint64    `json:"id" gorm:"primary_key"`
	UserId      int64     `json:"user_id" gorm:"index;not null"`
	User        User      `json:"-" gorm:"foreignKey:user_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Subject     string    `json:"subject" gorm:"index;size:255;not null"`
	Event       string    `json:"event" gorm:"index;size:255;not null"`
	Description string    `json:"description"  gorm:"type:text;not null"`
	Status      uint8     `json:"status" gorm:"index;default:1"`
	CreatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Activity) TableName() string {
	return "activities"
}
