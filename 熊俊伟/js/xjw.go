package main
import (
	"github.com/gin-gonic/gin"
	"net/http"
)
type User struct{
	Name string `form:"username" json:"username" binding:"required"`
}
func main(){
	MAP:= make(map[string]User)
	r := gin.Default()
	r.POST("add",func(c *gin.Context) { //增加
		var user User
		c.ShouldBind(&user)
		MAP[user.Name] = user
		c.JSON(200, gin.H{
			"add": user,
		})
	})
	r.GET("/getall",func(c *gin.Context){ //获取
		if len(MAP)!=0 {
			c.JSON(200, gin.H{
				"getall": MAP,
			})
		}else{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "暂未输入信息",
			})
			}
		})
	r.DELETE("/delete/:username",func(c *gin.Context) { //删
		username := c.Param("username")
		delete(MAP,username)
		c.JSON(200, gin.H{
			"delete": "删除成功",
		})
	})
	r.Run(":8080")
}