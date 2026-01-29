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

package schema

type UserLoginSchema struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegisterSchema struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"password_confirm"`
}

type UserForgotSchema struct {
	Email string `json:"email"`
}

type UserResetSchema struct {
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"password_confirm"`
}
