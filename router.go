package main

import (
	"github.com/RaymondCode/simple-demo/controller"
	"github.com/RaymondCode/simple-demo/utils"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// test api
	apiRouter.GET("/test/", controller.Test)

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)
	// user api

	userApi := apiRouter.Group("/user")
	// session middleware
	//userApi.Use(sessions.Sessions("mysession", global.DY_SESSION_STORE))
	userApi.POST("/register/", controller.Register)

	userApi.POST("/login/", controller.Login)

	////session login
	//userApi.GET("", controller.UserInfo)
	//
	//apiRouter.POST("/publish/action/", controller.Publish)
	//apiRouter.GET("/publish/list/", controller.PublishList)

	//jwt logic
	userApi.GET("", utils.JWTAuthMiddleware(), controller.UserInfo)

	apiRouter.POST("/publish/action/", utils.JWTAuthMiddleware(), controller.Publish)
	apiRouter.GET("/publish/list/", utils.JWTAuthMiddleware(), controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", utils.JWTAuthMiddleware(), controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", utils.JWTAuthMiddleware(), controller.FavoriteList)
	apiRouter.POST("/comment/action/", utils.JWTAuthMiddleware(), controller.CommentAction)
	apiRouter.GET("/comment/list/", utils.JWTAuthMiddleware(), controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", utils.JWTAuthMiddleware(), controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", utils.JWTAuthMiddleware(), controller.FollowList)
	apiRouter.GET("/relation/follower/list/", utils.JWTAuthMiddleware(), controller.FollowerList)
}
