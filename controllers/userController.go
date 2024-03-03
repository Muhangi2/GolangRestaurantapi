package controllers

import (
	"context"
	"fmt"
	"golang-Restaurantbooking/database"
	"golang-Restaurantbooking/helpers"
	"golang-Restaurantbooking/models"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		recordPerPage, err := strconv.Atoi(c.Query("recordperpage"))
		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}

		startIndex, err := strconv.Atoi(c.Query("startIndex"))
		if err != nil || startIndex < 0 {
			startIndex = (page - 1) * recordPerPage
		}

		matchStage := bson.D{{"$match", bson.D{{}}}}
		projectStage := bson.D{{"$project", bson.D{
			{"_id", 0},
			{"total_count", 1},
			{"user_items", bson.D{
				{"$slice", []interface{}{"$data", startIndex, recordPerPage}},
			}},
		}}}

		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, projectStage,
		})
		if err != nil {
			// Log the error for debugging
			log.Println("Error occurred while aggregating user items:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while listing user items"})
			return
		}

		var allUsers []bson.M
		if err := result.All(ctx, &allUsers); err != nil {
			// Log the error for debugging
			log.Println("Error occurred while decoding user items:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred while listing user items"})
			return
		}

		c.JSON(http.StatusOK, allUsers)
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		userId := c.Param("User_id")
		var user models.User
		err := userCollection.FindOne(ctx, bson.M{"user_id": userId})
		defer cancel()
		if err != nil {
			c.JSON(500, gin.H{"error": "error occured while fetching the user"})
		}
		c.JSON(200, user)
	}
}
func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		// Check for email existence
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking the email"})
			return
		}
		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "email already exists"})
			return
		}

		// Check for phone existence
		count, err = userCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while checking the phone"})
			return
		}

		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "phone already exists"})
			return
		}

		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at = user.Created_at
		user.ID = primitive.NewObjectID()
		user.User_id = user.ID.Hex()

		// Hash the password
		hashedPassword := HashPassword(*user.Password)
		user.Password = &hashedPassword

		// Generate the token
		token, refreshToken, err := helpers.GenerateAllTokens(*user.Email, *user.FirstName, *user.SecondName, user.User_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating tokens"})
			return
		}
		user.Token = &token
		user.Refresh_token = &refreshToken

		// Insert user into database
		result, resultError := userCollection.InsertOne(ctx, user)
		if resultError != nil {
			msg := fmt.Sprintf("user not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		//the response
		response := gin.H{
			"message":         "user created successfully",
			"user_id":         user.User_id,
			"result":          result,
			"token":           token,
			"refresh_token":   refreshToken,
			"hashed_password": hashedPassword,
		}

		c.JSON(http.StatusOK, response)
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found, login seems incorrect"})
			return
		}

		// Verify the password
		passwordValid := VerifyPassword(*user.Password, *foundUser.Password)
		if !passwordValid {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "incorrect password"})
			return
		}

		// Generate tokens
		token, refreshToken, err := helpers.GenerateAllTokens(*foundUser.Password, *foundUser.FirstName, *foundUser.SecondName, foundUser.User_id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating tokens"})
			return
		}

		// Update tokens
		helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id)

		// Respond with user details (excluding sensitive fields)
		response := gin.H{
			"user_id":       foundUser.User_id,
			"token":         token,
			"refresh_token": refreshToken,
		}

		c.JSON(http.StatusOK, response)
	}
}

func HashPassword(password string) string {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err)
	}
	return string(hashedBytes)
}

func VerifyPassword(userPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(userPassword))
	return err == nil
}
