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

type Authentication struct {
	Id         uint64     `json:"id" gorm:"primary_key"`
	UserId     int64      `json:"user_id" gorm:"index;not null"`
	User       User       `gorm:"foreignKey:user_id;references:id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	AuthType   string     `json:"type" gorm:"index;size:100;not null"`
	Credential string     `json:"credential" gorm:"index;size:180;not null"`
	Token      string     `json:"token" gorm:"index;size:100;not null"`
	Status     uint8      `json:"status" gorm:"index;default:0"`
	ExpiredAt  *time.Time `json:"expired_at" gorm:"index"`
	CreatedAt  time.Time  `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time  `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Authentication) TableName() string {
	return "authentications"
}
