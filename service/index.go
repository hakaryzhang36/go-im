package service

import "github.com/gin-gonic/gin"

// @BasePath

// GetIndex
// @Summary 首页api
// @Schemes
// @Description 访问首页的接口
// @Tags index
// @Produce json
// @Success 200 {string} welcome!
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome!",
	})
}
