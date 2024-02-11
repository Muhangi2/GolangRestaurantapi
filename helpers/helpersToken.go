package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"golang-Restaurantbooking/database"
)

// Struct
type SignedDetails struct {
	Email     string
	FirstName string
	LastName  string
	UID       string
	jwt.RegisteredClaims
}

var SECRET_KEY = []byte(os.Getenv("SECRET_KEY"))

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GenerateAllTokens(email string, firstName string, lastName string, uid string) (signedToken string, signedRefreshToken string, err error) {

	claims := &SignedDetails{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
		UID:       uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * 24)), // Corrected usage of time.Now().Local() and jwt.NewNumericDate
		},
	}
	refreshClaims := &SignedDetails{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * 24)),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString(SECRET_KEY)
	refreshTokens, err := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaims).SignedString(SECRET_KEY)

	if err != nil {
		log.panic(err)
		return
	}
	return token, refreshTokens, err

}

func UpdateAllTokens(signedToken string, signedRefreshToken string, userId string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var updateObj primitive.D
	updateObj = append(updateObj, bson.E{"token", signedToken})
	updateObj = append(updateObj, bson.E{"refreshtoken", signedRefreshToken})

	updated_at := time.Now()
	updateObj = append(updateObj, bson.E{"updated_at", updated_at})

	upsert := true
	filter := bson.M{"user_id": userId}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}
	_, err := userCollection.UpdateOne(ctx, filter, bson.D{{"$set", updateObj}}, &opt)
	defer cancel()
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}

func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {

	token, err := jwt.ParseWithClaims(signedToken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})
	//token valid
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}
	//the token is expired
	// Extract the time value from jwt.NumericDate and convert it to Unix time
	expiresAtUnix := claims.ExpiresAt.Time().Unix()

	// Compare the Unix time with the current local Unix time
	if expiresAtUnix < time.Now().Unix() {
		msg = "the token is invalid"
		return
	}

	return claims, msg
}
