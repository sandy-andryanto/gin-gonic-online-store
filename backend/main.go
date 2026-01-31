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

package main

import (
	config "backend/src/config"
	seed "backend/src/data"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}
	config.Config()
	seed.RunSeed()
	db := config.SetupDB()
	db.LogMode(true)
	r := config.SetupRoutes(db)
	r.Run("0.0.0.0:" + os.Getenv("APP_PORT"))
}
