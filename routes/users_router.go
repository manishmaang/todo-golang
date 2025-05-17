package routes

import (
	"todo-app/controllers"
	"todo-app/middlewares"

	"github.com/gin-gonic/gin"
)

// "fmt"
// "github.com/gin-gonic/gin"
// "github.com/golang-jwt/jwt/v5"
// "net/http"
// "strings"

func User_routes (r * gin.Engine){
  r.GET("/", middlewares.Authenticate_User(), controllers.Get_users) // since authenticate is a function which returns gin handler func there we explicitly need to call it
  
}

