package main

import (
	"encoding/json"
	"fmt"
	"os"
	"pizzaria/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

var pizzas = []models.Pizza{}

func main() {

	loadPizzas()
	router := gin.Default()
	router.GET("/pizzas", getPizzas)
	router.GET("/pizzas/:id", getPizzasById)
	router.POST("/pizzas", postPizza)
	router.Run(":8811") // listen and serve on 0.0.0.0:8080
}

func getPizzas(c *gin.Context) {

	c.JSON(200, pizzas)
}

func getPizzasById(c *gin.Context) {
	id := c.Param("id")
	fmt.Println("ID:", id)

	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}
	for _, pizza := range pizzas {
		fmt.Println("Checking pizza with ID:", pizza.Id)
		fmt.Println("IF", idInt, "==", pizza.Id, ": ", idInt == pizza.Id)
		if idInt == pizza.Id {
			c.JSON(200, pizza)
			return
		}
	}
	c.JSON(404, gin.H{"error": "Pizza not found"})
}

func postPizza(c *gin.Context) {
	var pizza models.Pizza
	if err := c.ShouldBindJSON(&pizza); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	pizza.Id = len(pizzas) + 1 // Simulating ID assignment
	pizzas = append(pizzas, pizza)
	savePizza() // Save the updated pizza list to file
	fmt.Println("Pizza added:", pizza)
	// Return the created pizza with a 201 status code
	c.JSON(201, pizza)
}

func loadPizzas() {
	fmt.Println("######## Loagind pizzas ########")
	// Simulating loading pizzas from a JSON file
	file, err := os.Open("../dados/pizzas.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	jsDecoder := json.NewDecoder(file)
	if err := jsDecoder.Decode(&pizzas); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
}

func savePizza() {
	fmt.Println("######## Save pizza ########")
	// Simulating loading pizzas from a JSON file
	file, err := os.Create("dados/pizzas.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	jsEncoder := json.NewEncoder(file)
	if err := jsEncoder.Encode(pizzas); err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println("Pizzas saved successfully")
}
