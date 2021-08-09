package main

import (
	"./Config"
	"./Repo"
	"./Route"
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
グループ周りのルーティングセットアップ
*/
func setupGroupRouter(api *gin.RouterGroup, repo *Repo.Repository) {

	groupRouter := api.Group("/group")

	groupRouter.GET("/:uuid", Route.GetGroupHandle(repo))
	groupRouter.PUT("/:uuid", Route.UpdateGroupHandle(repo))

	groupRouter.DELETE("/:uuid", Route.DeleteGroupHandle(repo))
	groupRouter.POST("", Route.AddGroupHandle(repo))
}

/**
ユーザー周りのルーティングセットアップ
*/
func setupUserRouter(api *gin.RouterGroup, repo *Repo.Repository) {
	userRouter := api.Group("/user")

	userRouter.GET("", Route.GetAllUsersProfileInGroupHandle(repo))
	userRouter.GET("/:uuid", Route.GetUserProfileHandle(repo))

	userRouter.PUT("/:uuid", Route.UpdateUserProfileHandle(repo))

	userRouter.DELETE("/:uuid", Route.DeleteUserHandle(repo))
	userRouter.POST("", Route.AddUserHandle(repo))
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func SetupServer() error {

	repository, repoErr := Repo.SetupDatabase(Config.Env)

	if repoErr != nil {
		return repoErr
	}

	engine := gin.Default()

	//Version 1
	v1 := engine.Group("/api/v1")

	v1.Use(corsMiddleware())

	/**
	グループ周りのAPI
	*/
	setupGroupRouter(v1, repository)
	setupUserRouter(v1, repository)

	return engine.Run(fmt.Sprintf(":%d", Config.Env.App_Port))
}
