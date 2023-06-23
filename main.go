package main

import (
	"encoding/json"
	"example/entities"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Load data within the context
	router.Use(myMiddleware())

	//fetch users
	router.GET("/users", getUsers)
	//fetch posts
	router.GET("/posts", getPosts)

	router.GET("/usersAndPost/:id", getUsersAndPostByID)

	router.Run("localhost:8082")
}

func myMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
		c.Set("users", getUsersData())
		c.Set("posts", getPostsData())
        c.Next()
    }
}

func getUsersData()([]entities.User){
	file1, _ := ioutil.ReadFile("templates/users.json")

	x := entities.Users{}

	err := json.Unmarshal([]byte(file1), &x.List)
	if err!=nil{
		return nil
	}
	return x.List
}

func getPostsData()([]entities.Post){
	file1, _ := ioutil.ReadFile("templates/posts.json")

	x := entities.Posts{}

	err := json.Unmarshal([]byte(file1), &x.List)
	if err!=nil{
		return nil
	}
	return x.List
}

func getUsers(c *gin.Context){
	db, ok := c.Keys["users"].([]entities.User)
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Error in fetching User Data"})
        return
	}
	c.IndentedJSON(http.StatusOK, db)
}

func getPosts(c *gin.Context){
	db, ok := c.Keys["posts"].([]entities.Post)
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Error in fetching Posts Data"})
        return
	}
	c.IndentedJSON(http.StatusOK, db)
}

func getUsersAndPostByID(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err!=nil{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Request Parameters did not parse successfully."})
		return
	}

	x:=entities.Resp{}
	found:=false

	users, ok := c.Keys["users"].([]entities.User)
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Error in fetching User Data"})
        return
	}

	for _, v := range users {
		if v.ID == id {
			found=true
			x.ID = v.ID
			x.Name = v.Name
			x.Lat = v.Address.Geo.Lat
			x.Long = v.Address.Geo.Lng
			x.CompName = v.Company.Name
			break
		}
	}

	if !found{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"The user with id "+c.Param("id")+" does not exist."})
        return
	}

	posts, ok := c.Keys["posts"].([]entities.Post)
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"Error in fetching Posts Data"})
        return
	}

	found = false
	for _, v := range posts {
		if v.ID == id {
			found=true
			x.Title = v.Title
			x.Body = v.Body
			break
		}
	}

	if !found{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message":"The post with id "+c.Param("id")+" does not exist."})
        return
	}

	c.IndentedJSON(http.StatusOK, x)
}