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

package schema

import "database/sql"

type UserProfileSchema struct {
	Image     sql.NullString `json:"image"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Gender    string         `json:"gender"`
	Country   string         `json:"country"`
	City      string         `json:"city"`
	ZipCode   string         `json:"zip_code"`
	Address   string         `json:"address"`
}

type UserPasswordSchema struct {
	OldPassword     string `json:"old_password"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"password_confirm"`
}
