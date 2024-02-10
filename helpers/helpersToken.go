package helpers

import (
	"golang-Restaurantbooking/database"

	"go.mongodb.org/mongo-driver/mongo"
)

var usercollection *mongo.Collection = database.OpenCollection(database.Client, "user")

func GenerateAllTokens() {

}
func UpdateAllTokens() {

}
func ValidateToken() {

}
