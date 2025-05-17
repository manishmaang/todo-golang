package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"todo-app/config"
	"todo-app/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Get_users(c *gin.Context) {
	pageStr := c.Query("page") // returns "" if not present
	page := 1                  // default value

	fmt.Println("printing the pagestr and page here", pageStr, " now page : ", page)

	if pageStr != "" {
		page_number, err := strconv.Atoi(pageStr)
		if err == nil {
			page = page_number
		}
	}

	db := config.GetDB()
	user_collection := db.Collection("users")

	limit := 10
	asked_limit := c.Query("limit")
	if asked_limit != "" {
		new_limit, err := strconv.Atoi(asked_limit)
		if err == nil && new_limit < limit {
			limit = new_limit
		}
	}

	skip := limit * (page - 1)

	findOptions := options.Find()
	findOptions.SetLimit(int64(limit))
	findOptions.SetSkip(int64(skip))

	cursor, err := user_collection.Find(context.TODO(), bson.M{}, findOptions) //all queries using driver pass an bson document so no use of lean as data is already an obj
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
		return
	}

	defer cursor.Close(context.TODO())

	var users []models.Users

	for cursor.Next(context.TODO()) {
		var user models.Users
		if err := cursor.Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding user"})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

// func GetUsers(c *gin.Context) {
//     db := config.GetDB()
//     collection := db.Collection("users")

//     // Parse page query param (use your existing code)
//     pageStr := c.Query("page")
//     page := 1
//     if pageStr != "" {
//         if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
//             page = p
//         }
//     }

//     limit := int64(10) // number of users per page
//     skip := int64((page - 1) * int(limit))

//     findOptions := options.Find()
//     findOptions.SetLimit(limit)
//     findOptions.SetSkip(skip)

//     cursor, err := collection.Find(context.TODO(), bson.M{}, findOptions)
//     if err != nil {
//         c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching users"})
//         return
//     }
//     defer cursor.Close(context.TODO())

//     var users []models.Users
//     for cursor.Next(context.TODO()) {
//         var user models.Users
//         if err := cursor.Decode(&user); err != nil {
//             c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding user"})
//             return
//         }
//         users = append(users, user)
//     }

//     c.JSON(http.StatusOK, users)
// }
