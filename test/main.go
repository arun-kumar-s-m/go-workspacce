package main
import (
    "context"
	"fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/go-redis/redis/v8"
)
func main(){
    rdb := redis.NewClusterClient(&redis.ClusterOptions{
        Addrs: [] string {"localhost:6375"},
    })
    m := map[string]string{
        "3FjaMee": "https://moz.com/blog/optimize-featured-snippets",
        "3FjaMee22": "http://localhost:8082/",
    }
    _,err :=rdb.HSet(context.Background(),"token",m).Result()
    if err != nil{
        panic(err)
    }
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
		c.Status(200)
    })
    r.GET("/hello/:id", func(c *gin.Context) {
        // res,err := rdb.HGetAll(c,"token").Result()
        // res := rdb.HGet(c,"token","3FjaMee")
        // if err != nil{
        //     c.JSON(http.StatusConflict,gin.H{
        //         "message": "failed",
        //     })
        // }
//url path spro.io/tokem-dasdfsad
//        res,err := rdb.HGet(c,"tokem-dasdfsad","https://surveysparrow.com").Result()
        idVal := c.Param("id")
        fmt.Println(idVal)
        res,err := rdb.HGet(c,"token",idVal).Result()
		fmt.Println(res)
        if err != nil{
            c.JSON(http.StatusConflict,gin.H{
                "message": "failed",
            })
        }

        // c.JSON(http.StatusMovedPermanently, res)
		c.Status(301)
		c.Writer.Header().Set("location", res)
		// c.Header("location", "http://www.google.com/")
		// c.JSON(http.StatusMovedPermanently, res)
    })
    r.Run()
}