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

package config

import (
	models "backend/src/models"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func SetupDB() *gorm.DB {
	godotenv.Load(".env")
	USER := os.Getenv("DB_USERNAME")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DBNAME := os.Getenv("DB_DATABASE")
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, PORT, DBNAME)
	db, err := gorm.Open(os.Getenv("DB_CONNECTION"), URL)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Config() {
	db := SetupDB()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.Activity{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.Authentication{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.Brand{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.Category{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.Colour{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.NewsLetter{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.Order{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.OrderBilling{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.OrderDetail{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.Payment{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.Product{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.ProductImage{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.ProductInventory{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.ProductReview{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.Setting{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.Size{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(models.User{})
}