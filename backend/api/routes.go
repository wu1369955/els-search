package api

import (
	"backend/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// SetupRoutes 设置 API 路由
func SetupRoutes(router *gin.Engine, searchService service.SearchService) {
	// 搜索接口
	router.GET("/api/search", func(c *gin.Context) {
		// 获取查询参数
		index := c.Query("index")
		query := c.Query("q")
		fields := strings.Split(c.Query("fields"), ",")
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
		size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

		// 验证参数
		if index == "" || query == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "index and q are required"})
			return
		}

		// 执行搜索
		result, err := searchService.Search(c.Request.Context(), index, query, fields, page, size)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// 返回结果
		c.JSON(http.StatusOK, result)
	})

	// 健康检查接口
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})
}
