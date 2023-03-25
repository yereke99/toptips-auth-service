package middleware

import "github.com/gin-gonic/gin"

func AuthorizeAdmin() gin.HandlerFunc

func AuthorizeWorker() gin.HandlerFunc
