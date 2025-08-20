package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Cat struct {
    ID       string  `json:"id"`
    Name     string  `json:"name"`
    Breed    string  `json:"breed"`
    Age     int     `json:"age"`
    Weight      float64 `json:"weight"`
}

var cats = []Cat{
    {ID: "001", Name: "Som Som", Breed: "Scottish Fold", Age: 2, Weight: 5.30},
    {ID: "002", Name: "Sam Lee", Breed: "Sphynx", Age: 1, Weight: 2.15},
	{ID: "003", Name: "Kao Jao", Breed: "American Shorthair", Age: 3, Weight: 8.13},
	{ID: "004", Name: "Kanom Jean", Breed: "Persian", Age: 2, Weight: 4.25},
	{ID: "005", Name: "Pan Pee", Breed: "Main Coon", Age: 1, Weight: 1.95},
}

func getCats(c *gin.Context) {
	ageQuery := c.Query("age")

	if ageQuery != "" {
		filter := []Cat{}
		for _, cat := range cats {
			if fmt.Sprint(cat.Age) == ageQuery {
				filter = append(filter, cat)
			}
		}
		c.JSON(http.StatusOK, filter)
		return
	}
	c.JSON(http.StatusOK, cats)
}

func main()  {
	r := gin.Default()

	r.GET("/health", func (c *gin.Context)  {
		c.JSON(200, gin.H{"message" : "healthy"})
	})

	api := r.Group("/api/v1") 
	{
		api.GET("/cats", getCats)
	}

	r.Run(":8080")
}