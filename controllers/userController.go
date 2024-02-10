package controllers

import (
	"bytes"
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
		recordperpage, err := strconv.Atoi(c.Query("recordperpage"))
		if err != nil || recordperpage < 1 {
			recordperpage = 10
		}
		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}
		startIndex := (page - 1) * recordperpage
		startIndex, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{"$match", bson.D{{}}}}
		projectStage := bson.D{{"$project", bson.D{
			{"_id", 0},
			{"total_count", 1},
			{"user_items", bson.D{
				{"$slice", []interface{}{"$data", startIndex, recordperpage}},
			}},
		}}}

		result, err := userCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, projectStage,
		})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while listing to user items"})
		}
		var allUsers []bson.M
		if err = result.All(ctx, &allUsers); err != nil {
			log.Fatal(err)
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
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		//change to struct form
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		//validate the struct
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		}
		//check for email
		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking the email"})
			return
		}
		password := HashPassword(*user.Password)
		user.Password = &password
		
		 count,err:userCollection.CountDocuments(ctx,bson.M{"phone":user.Phone})
		 defer cancel()
		 if err1=nil{
			log.Panic(err)
			c.JSON(http.StatusInternalServerError,gin.H{"error":"error occures while checking the phone"})
	   	return 
		}
		 if count > 0 {
          c.JSON(http.StatusInternalServerError,gin.H{"error":"phone or eamil already exists."})
		 }
		user.Created_at=time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
		user.Updated_at=time.Parse(time.RFC3339,time.Now().Format(time.RFC3339))
		user.ID=primitive.NewObjectID()
		user.User_id=user.ID.Hex()

		//generate the token
		token,refreshToken,_:=helpers.GenerateAllTokens(*user.Email,*user.FirstName,*user.SecondName,user.User_id)
		user.Token=&token
		user.Refresh_token=&Refresh_token
    
		result,resultError :=userCollection.InsertOne(ctx,user)
		if resultError !=nil{
			msg:=fmt.Sprintf("user not created")
			c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
			return
		}
		defer cancel();
		c.JSON(200,result)
	
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
   var ctx,cancel=context.WithTimeout(context.Background(),100*time.Second)
   var user models.User
   Var foundUser models.User

   if err!=c.BindJSON(&user);
   err!=nil{
	c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
   }
   err:=userCollection.FindOne(ctx,bson.M{"email":user.Email}).Decode(&foundUser)
   defer cancel()
   if err!=nil{
	c.JSON(http.StatusInternalServerError,gin.H{"error":"user not found login seems not correct"})
	return
   }
   //lets verify the password
   passwordValid,msg:=VerifyPassword(*user.Password,*foundUser.Password)
   defer  cancel()
 if passwordValid!=true{
	c.JSON(http.StatusInternalServerError,gin.H{"error":msg})
	return
 }
//if all goes well generate the token
token,refreshToken,_:=helper.GenerateAllTokens(*founderUser.Password,*founderUser.FirstName,*founderUser.SecondName,founderUser.User_id)
//update tokens
helper.UpdateAllTokens(token,refreshToken,foundUser.user_id)
//status
c.JSON(http.StatusOK,foundUser)
	}
}

func HashPassword(password string) string {
  bytes,err:=bcrypt.GenerateFromPassword([]byte(password),14)
  if err!=nil{
	log.Panic(err)
  }
	return string(bytes)

}

func VerifyPassword(userPassword string, providePassword string) (bool, string) {
   err:=bcrypt.CompareHashAndPassword([]byte(providePassword),[]byte(userPassword))
   check:=true
   msg:=""

   if err!=nil{
	msg=fmt.Sprintf("password is incorrect")
	check:=false
   }
   return check,msg
}
