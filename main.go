package main

import (
	"StudentRecord/domain/aggregates/userAggregates"
	_ "StudentRecord/infrastructure/DataBase"
	"StudentRecord/infrastructure/Repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()
	r.POST("/user", func(context *gin.Context) {
		var param UserParam
		err := context.BindJSON(&param)
		if err != nil {
			return
		}
		fmt.Println(param)
		var userRep userAggregates.IUserRepository = &Repositories.UserRepository{}
		userRep.Add(userAggregates.Create(param.UserName, param.Password, param.NickName))
	})
	log.Fatal(r.Run(":6010"))
}

type UserParam struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	NickName string `json:"nickName"`
}
