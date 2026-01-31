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

type ReviewSchema struct {
	Review string `json:"review"`
	Rating int32  `json:"rating"`
}

type CreateCartSchema struct {
	SizeId   uint64 `json:"size_id"`
	ColourId uint64 `json:"colour_id"`
	Qty      uint32 `json:"qty"`
}

type CheckoutSchema struct {
	PaymentId uint64 `json:"payment_id"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Country   string `json:"country"`
	City      string `json:"city"`
	ZipCode   string `json:"zip_code"`
	Address   string `json:"address"`
	Notes     string `json:"notes"`
}
