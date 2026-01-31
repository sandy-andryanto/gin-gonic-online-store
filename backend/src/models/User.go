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

type User struct {
	Id              uint64         `json:"id" gorm:"primary_key"`
	Email           string         `json:"email" gorm:"index;size:191;not null"`
	Phone           string         `json:"phone" gorm:"index;size:191;default:null"`
	Password        string         `json:"password" gorm:"index;size:255;not null"`
	Salt            string         `json:"salt" gorm:"index;size:255;"`
	Image           sql.NullString `json:"image" gorm:"index;size:191;default:null;"`
	FirstName       sql.NullString `json:"first_name" gorm:"index;size:191;default:null;"`
	LastName        sql.NullString `json:"last_name" gorm:"index;size:191;default:null;"`
	Gender          sql.NullString `json:"gender" gorm:"index;size:2;default:null;"`
	Country         sql.NullString `json:"country" gorm:"index;size:191;default:null;"`
	City            sql.NullString `json:"city" gorm:"index;size:191;default:null;"`
	ZipCode         sql.NullString `json:"zip_code" gorm:"index;size:64;default:null;"`
	Address         sql.NullString `json:"address"  gorm:"type:text;default:null;"`
	Status          uint8          `json:"status" gorm:"index;default:0"`
	CreatedAt       time.Time      `gorm:"index;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"index;default:CURRENT_TIMESTAMP" json:"updated_at"`
	Products        []Product      `gorm:"many2many:products_wishlists"`
	Activities      []Activity
	Authentications []Authentication
	Reviews         []ProductReview
}

func (User) TableName() string {
	return "users"
}
