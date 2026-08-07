package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type album struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}

var albums = []album{
	{ID:"1",Title:"Blue Train",Artist:"John Coltrane",Price:56.99},
	{ID:"2",Title:"Jeru",Artist:"Gerry Mulligan",Price:17.99},
	{ID:"3",Title:"Sarah Vaughan and Clifford Brown",Artist:"Sarah Vaughan",Price:39.99},
} 

func main(){
	//initialized the gin router 
	router :=gin.Default()

	//Associate the Get methpod and /albums path with the getAlbums function
	router.GET("/albums",getAlbums)


	//Associate the POST methpod and /albums path with the postAlbums function
	router.POST("/albums",postAlbums)

	// Associate the GET method and /albums/:id path with the getAlbumByID function
	router.GET("/albums/:id",getAlbumByID)

	//start the server
	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context){
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err !=nil{
		return
	}
    
	//add new album to the slice
	albums = append(albums,newAlbum)

	//send back a 201 created status with the new album details
	c.IndentedJSON(http.StatusCreated,newAlbum)
}

func getAlbumByID(c *gin.Context){
	//get id from the parameter
	id := c.Param("id")

	//loop through the slice and find the album with the matching id
	for _,a :=range albums{
		if a.ID == id {
			c.IndentedJSON(http.StatusOK,a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message":"album not found"})

}