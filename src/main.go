package main

import (
	"fmt"
	"net/http"
	"log"
	"encoding/json"
	"github.com/go-redis/redis"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

// var ctx = context.Background()

func RedisClient(w http.ResponseWriter,r *http.Request) {
    rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6375",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    // err := rdb.Set("3FjaMee", "https://moz.com/blog/optimize-featured-snippets", 0).Err()
    // if err != nil {
    //     panic(err)
    // }

	// mapp := map[string] string {"short_url_a" : "long_url_a"}
	// err := rdb.HSet(context.Background(),"token",mapp)
	//  if err != nil {
    //     panic(err)
    // }
	// err := rdb.Set("3FjaMee", "https://moz.com/blog/optimize-featured-snippets", 0).Err()
    // if err != nil {
    //     panic(err)
    // }

	// sh,err := redis.HGet(context.Background(),"token","short_url_a").Result()
	// if err != nil {
    //     panic(err)
    // }
	// fmt.Fprint(w,"VALUE IS ",sh)
	
    val, err := rdb.Get( "3FjaMee").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("ARUN GETTING VALUE FOR KEY : key :::: ", val)
	fmt.Fprint(w,"VALUE IS ",val)
	

    // val2, err := rdb.Get("key2").Result()
    // if err == redis.Nil {
    //     fmt.Println("key2 does not exist")
    // } else if err != nil {
    //     panic(err)
    // } else {
    //     fmt.Println("key2", val2)
    // }
    // Output: key value
    // key2 does not exist
}

type Articles []Article

func allArticles(w http.ResponseWriter,r *http.Request){
	articles := Articles{
		Article{Title:"Title 1",Desc:"Description 1",Content:"Content 1"},
	}
	fmt.Println("Endpoint Hit : All articles ")
	json.NewEncoder(w).Encode(articles)
}
func homePage(w http.ResponseWriter,r *http.Request){
	fmt.Fprint(w,"Homepage Endpoint Hit")
}

func handleRequests(){
	http.HandleFunc("/",homePage)
	http.HandleFunc("/articles",allArticles)
	http.HandleFunc("/redis",RedisClient)
	log.Fatal(http.ListenAndServe(":8081",nil))
}

func main(){
	handleRequests()
	fmt.Println("Go Redis Tutorial")
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6375",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}