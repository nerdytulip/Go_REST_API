package http
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r*http.Request) {
	articles := Articles{
		Article{Title:"Test Title", Desc:"Test Description", Content:"Test Content: Hello World" },
	}

	fmt.Println("EndPoint Hit: All Articles EndPoint")

	json.NewEncoder(w).Encode(articles)
}


func homePage (w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "HomePage endpoint hit")
}

func handleRequest(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles",allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main(){
	handleRequest()
}
