/**
 * This file is part of the Sandy Andryanto Blog Applicatione.
 *
 * @author     Sandy Andryanto <sandy.andryanto.blade@gmail.com>
 * @copyright  2024
 *
 * For the full copyright and license information,
 * please view the LICENSE.md file that was distributed
 * with this source code.
 */

package controllers

import (
	helpers "backend/src/helpers"
	models "backend/src/models"
	schema "backend/src/schema"
	services "backend/src/services"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func AuthLogin(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var user models.User

	var input schema.UserLoginSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(strings.TrimSpace(input.Email)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The email field is required.!"})
		return
	}

	if len(strings.TrimSpace(input.Password)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The password field is required.!"})
		return
	}

	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user with e-mail address " + input.Email + " not found!"})
		return
	}

	decrypt := helpers.Decrypt(user.Password, user.Salt)

	if user.Status == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "You need to confirm your account. We have sent you an activation code, please check your email.!"})
		return
	}

	if input.Password != decrypt {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "incorrect password!"})
		return
	}

	Activity := models.Activity{
		UserId:      int64(user.Id),
		Subject:     "User Login",
		Event:       "Sign In",
		Description: "Sign in to application",
	}
	db.Create(&Activity)

	c.JSON(http.StatusOK, gin.H{"token": services.JWTAuthService().GenerateToken(int(user.Id), user.Email, true)})
}

func AuthRegister(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var user models.User

	var input schema.UserRegisterSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(strings.TrimSpace(input.Name)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The name field is required.!"})
		return
	}

	if len(strings.TrimSpace(input.Email)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The email field is required.!"})
		return
	}

	if len(strings.TrimSpace(input.Password)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The passwword field is required.!"})
		return
	}

	if len(strings.TrimSpace(input.ConfirmPassword)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The password_confirm field is required.!"})
		return
	}

	if len(strings.TrimSpace(input.Password)) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you have to enter at least 8 digit!"})
		return
	}

	if strings.TrimSpace(input.Password) != strings.TrimSpace(input.ConfirmPassword) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "These passwords don't match!"})
		return
	}

	if err := db.Where("email = ?", input.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The email already exists"})
		return
	}

	bytes := make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	key := hex.EncodeToString(bytes) //encode key in bytes to string and keep as secret, put in a vault
	encrypted := helpers.Encrypt(input.Password, key)
	token := (uuid.New()).String()

	FirstName := ""
	LastName := ""
	Names := strings.Split(input.Name, " ")

	if len(Names) > 0 {
		FirstName = Names[0]
		LastName = strings.Join(Names[1:], " ")
	} else {
		FirstName = Names[0]
	}

	User := models.User{
		FirstName: helpers.NewNullString(FirstName),
		LastName:  helpers.NewNullString(LastName),
		Email:     input.Email,
		Password:  encrypted,
		Status:    0,
		Salt:      key,
	}
	db.Create(&User)

	Verification := models.Authentication{
		UserId:     int64(User.Id),
		AuthType:   "email-confirm",
		Credential: input.Email,
		Token:      token,
		Status:     0,
		ExpiredAt: func(t time.Time) *time.Time {
			t = t.Add(30 * time.Minute)
			return &t
		}(time.Now()),
	}
	db.Create(&Verification)

	Activity := models.Activity{
		UserId:      int64(User.Id),
		Subject:     "User Register",
		Event:       "Sign Up",
		Description: "Register new user account",
	}
	db.Create(&Activity)

	c.JSON(http.StatusOK, gin.H{"message": "Your account has been created. Please check your email for the confirmation message we just sent you.", "token": token})
}

func AuthConfirm(c *gin.Context) {

	var verification models.Authentication
	var user models.User

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("token = ? AND status = ? ", c.Param("token"), 0).First(&verification).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "We can't find a user with that  token is invalid.!"})
		return
	}

	if err := db.Where("email = ?", verification.Credential).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found!"})
		return
	}

	updateVerification := models.Authentication{
		ExpiredAt: func(t time.Time) *time.Time { return &t }(time.Now()),
		Status:    2,
	}

	updateUser := models.User{
		Status: 1,
	}

	if err := db.Model(&verification).Updates(updateVerification).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update verification"})
		return
	}

	if err := db.Model(&user).Updates(updateUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	Activity := models.Activity{
		UserId:      int64(user.Id),
		Subject:     "User Verification",
		Event:       "Email Confirmation",
		Description: "Confirm new member registration account",
	}
	db.Create(&Activity)

	c.JSON(http.StatusOK, gin.H{"message": "Your registration is complete. Now you can login."})
}

func AuthEmailForgot(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var user models.User

	var input schema.UserForgotSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(strings.TrimSpace(input.Email)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The email field is required.!"})
		return
	}

	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "We can't find a user with that e-mail address."})
		return
	}

	token := (uuid.New()).String()

	ResetPassword := models.Authentication{
		UserId:     int64(user.Id),
		AuthType:   "reset-password",
		Credential: input.Email,
		Token:      token,
		Status:     0,
		ExpiredAt: func(t time.Time) *time.Time {
			t = t.Add(30 * time.Minute)
			return &t
		}(time.Now()),
	}
	db.Create(&ResetPassword)

	Activity := models.Activity{
		UserId:      int64(user.Id),
		Subject:     "Request Forgot Password",
		Event:       "Forgot Password",
		Description: "Request reset password link",
	}
	db.Create(&Activity)

	c.JSON(http.StatusOK, gin.H{"message": "We have e-mailed your password reset link!", "token": token})
}

func AuthEmailReset(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var user models.User
	var resetPassword models.Authentication

	var input schema.UserResetSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(strings.TrimSpace(input.Email)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The email field is required.!"})
		return
	}

	if len(strings.TrimSpace(input.Password)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The password field is required.!"})
		return
	}

	if len(strings.TrimSpace(input.ConfirmPassword)) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The password_confirm field is required.!"})
		return
	}

	if len(strings.TrimSpace(input.Password)) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you have to enter at least 8 digit!"})
		return
	}

	if strings.TrimSpace(input.Password) != strings.TrimSpace(input.ConfirmPassword) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "These passwords don't match!"})
		return
	}

	if err := db.Where("credential = ? AND token = ? AND status = 0", input.Email, c.Param("token")).First(&resetPassword).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "This password and email reset token is invalid."})
		return
	}

	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user with e-mail address " + input.Email + " not found!"})
		return
	}

	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	key := hex.EncodeToString(bytes)
	encrypted := helpers.Encrypt(input.Password, key)

	updateUser := models.User{
		Status:   1,
		Password: encrypted,
		Salt:     key,
	}

	updatePasswordReset := models.Authentication{
		ExpiredAt: func(t time.Time) *time.Time { return &t }(time.Now()),
		Status:    2,
	}

	if err := db.Model(&resetPassword).Updates(updatePasswordReset).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password reset"})
		return
	}

	if err := db.Model(&user).Updates(updateUser).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	Activity := models.Activity{
		UserId:      int64(user.Id),
		Subject:     "User Recovery",
		Event:       "Reset Password",
		Description: "Reset account password",
	}
	db.Create(&Activity)

	c.JSON(http.StatusOK, gin.H{"message": "Your password has been reset!"})
}
