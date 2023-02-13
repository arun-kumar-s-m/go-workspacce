package main
import (
    "context"
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
    }
    _,err :=rdb.HSet(context.Background(),"token",m).Result()
    if err != nil{
        panic(err)
    }
    r := gin.Default()
    r.GET("/hello", func(c *gin.Context) {
        // res,err := rdb.HGetAll(c,"token").Result()
        // res := rdb.HGet(c,"token","3FjaMee")
        // if err != nil{
        //     c.JSON(http.StatusConflict,gin.H{
        //         "message": "failed",
        //     })
        // }

        res,err := rdb.HGet(c,"token","3FjaMee").Result()
        if err != nil{
            c.JSON(http.StatusConflict,gin.H{
                "message": "failed",
            })
        }

        c.JSON(http.StatusOK, res)
    })
    r.Run()
}