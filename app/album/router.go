package album

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest_api/app"
	"strconv"
)

// @Summary Получить список альбомов
// @Description Получить список всех альбомов
// @Tags Albums
// @Accept json
// @Produce json
// @Success 200 {array} Album
// @Router /albums [get]
func getAlbums(c *gin.Context) {
	db, err := app.GetDBEngine()
	if err != nil {
		panic(err)
	}
	repo := NewAlbumRepository(db)
	albums, err := repo.GetAlbums()
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusOK, albums)

}

// @Summary Создать альбом
// @Description Создать новый альбом
// @Tags Albums
// @Accept json
// @Produce json
// @Param input body CreateAlbumInput true "Данные альбома"
// @Success 201 {object} Album
// @Router /albums [post]
func createAlbums(c *gin.Context) {
	var input CreateAlbumInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while parse json"})
		return
	}
	if *input.Price <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "price cannot be less than or equal to 0"})
		return
	}
	db, err := app.GetDBEngine()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while get db engine"})
		return
	}
	repo := NewAlbumRepository(db)
	newAlbum, err := repo.CreateAlbum(input.Title, input.Artist, *input.Price)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while create new album"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// @Summary Удалить альбом по ID
// @Description Удаляет альбом по указанному ID
// @Tags Albums
// @Accept json
// @Produce json
// @Param id query int true "ID альбома для удаления"
// @Success 202 {object} int
// @Router /albums [delete]
func deleteAlbums(c *gin.Context) {
	strId := c.Query("id")
	id, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	db, err := app.GetDBEngine()
	if err != nil {
		panic(err)
	}
	repo := NewAlbumRepository(db)
	count, err := repo.DeleteAlbum(uint(id))
	if err != nil {
		panic(err)
	}
	c.IndentedJSON(http.StatusAccepted, count)
}

func InitAlbumRouter(engine *gin.Engine) {
	//db := app.GetDBEngine()
	//repo := NewAlbumRepository(db)

	albumRouter := engine.Group("/albums")
	{
		albumRouter.GET("", getAlbums)
		albumRouter.POST("", createAlbums)
		albumRouter.DELETE("", deleteAlbums)
	}

}
