package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This service lists all Nigerian companies, their logos, address, type and categories"})
	})

	router.GET("/companies", func(c *gin.Context) {
		companies, err := getCompanies()
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}
		c.JSON(http.StatusOK, companies)
	})

	router.POST("/companies", func(c *gin.Context) {
		fmt.Println("Adding new company...")
		var newCompany Company
		if err := c.ShouldBindJSON(&newCompany); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result := addNew(newCompany.Name, newCompany.Sector, newCompany.Category, newCompany.CEO, newCompany.Revenue, newCompany.IsStartup)
		c.JSON(http.StatusCreated, result)
	})

	return router
}

func main() {
	router := setupRouter()
	router.Run()
}
