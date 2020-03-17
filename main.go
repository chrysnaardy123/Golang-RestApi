package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "html/template"
    "path"
)
type Article struct {
    Name string `json:"name"`
    Surname string `json:"surname"`
    Fullname string `json:"fullname"`

}

//Simulasi database untuk array nanti
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request){
  	var filepath = path.Join("views", "index.html")
var tmpl, err = template.ParseFiles(filepath)
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}

var data = map[string]interface{}{
    "title": "Golang Resful Api",
}

err = tmpl.Execute(w, data)
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
}
    fmt.Println("Endpoint Hit: homePage")
}


func main() {

	Articles = []Article{
Article{Name: "Sidney",Surname: "Pittman",Fullname: "Jimmy Woodward"},
Article{Name: "Jerome",Surname: "Hall",Fullname: "Ashley Schwartz"},
Article{Name: "Vincent",Surname: "Nichols",Fullname: "Adam Chang"},
Article{Name: "Allen",Surname: "Burnette",Fullname: "Denise Nixon"},
Article{Name: "Sherri",Surname: "Dodson",Fullname: "Dana Finch"},
Article{Name: "Rita",Surname: "Britt",Fullname: "Brooke Coley"},
Article{Name: "Sarah",Surname: "Sanchez",Fullname: "Claire McKenna"},
Article{Name: "Shannon",Surname: "Godwin",Fullname: "Dorothy Fernandez"},
Article{Name: "Ellen",Surname: "Corbett",Fullname: "Nelson Stanley"},
Article{Name: "Frederick",Surname: "Lawson",Fullname: "Wesley McClure"},
Article{Name: "Maria",Surname: "Anthony",Fullname: "Marguerite Shapiro"},
Article{Name: "Julie",Surname: "Song",Fullname: "Ted Aldridge"},
Article{Name: "Joan",Surname: "Burgess",Fullname: "Helen Dickinson"},
Article{Name: "Colleen",Surname: "Murphy",Fullname: "Christopher Kelley"},
Article{Name: "Gary",Surname: "Bowling",Fullname: "Brenda Fletcher"},
Article{Name: "Christopher",Surname: "Harper",Fullname: "Peter Harris"},
Article{Name: "Charlene",Surname: "McKenzie",Fullname: "Clifford Lindsay"},
Article{Name: "Sidney",Surname: "Pittman",Fullname: "Jimmy Woodward"},
Article{Name: "Jerome",Surname: "Hall",Fullname: "Ashley Schwartz"},
Article{Name: "Vincent",Surname: "Nichols",Fullname: "Adam Chang"},
Article{Name: "Allen",Surname: "Burnette",Fullname: "Denise Nixon"},
Article{Name: "Sherri",Surname: "Dodson",Fullname: "Dana Finch"},
Article{Name: "Rita",Surname: "Britt",Fullname: "Brooke Coley"},
Article{Name: "Sarah",Surname: "Sanchez",Fullname: "Claire McKenna"},
Article{Name: "Shannon",Surname: "Godwin",Fullname: "Dorothy Fernandez"},
Article{Name: "Ellen",Surname: "Corbett",Fullname: "Nelson Stanley"},
Article{Name: "Frederick",Surname: "Lawson",Fullname: "Wesley McClure"},
Article{Name: "Maria",Surname: "Anthony",Fullname: "Marguerite Shapiro"},
Article{Name: "Julie",Surname: "Song",Fullname: "Ted Aldridge"},
Article{Name: "Joan",Surname: "Burgess",Fullname: "Helen Dickinson"},
Article{Name: "Colleen",Surname: "Murphy",Fullname: "Christopher Kelley"},
Article{Name: "Gary",Surname: "Bowling",Fullname: "Brenda Fletcher"},
Article{Name: "Christopher",Surname: "Harper",Fullname: "Peter Harris"},
Article{Name: "Charlene",Surname: "McKenzie",Fullname: "Clifford Lindsay"},
Article{Name: "Sidney",Surname: "Pittman",Fullname: "Jimmy Woodward"},
Article{Name: "Jerome",Surname: "Hall",Fullname: "Ashley Schwartz"},
Article{Name: "Vincent",Surname: "Nichols",Fullname: "Adam Chang"},
Article{Name: "Allen",Surname: "Burnette",Fullname: "Denise Nixon"},
Article{Name: "Sherri",Surname: "Dodson",Fullname: "Dana Finch"},
Article{Name: "Rita",Surname: "Britt",Fullname: "Brooke Coley"},
Article{Name: "Sarah",Surname: "Sanchez",Fullname: "Claire McKenna"},
Article{Name: "Shannon",Surname: "Godwin",Fullname: "Dorothy Fernandez"},
Article{Name: "Ellen",Surname: "Corbett",Fullname: "Nelson Stanley"},
Article{Name: "Frederick",Surname: "Lawson",Fullname: "Wesley McClure"},
Article{Name: "Maria",Surname: "Anthony",Fullname: "Marguerite Shapiro"},
Article{Name: "Julie",Surname: "Song",Fullname: "Ted Aldridge"},
Article{Name: "Joan",Surname: "Burgess",Fullname: "Helen Dickinson"},
Article{Name: "Colleen",Surname: "Murphy",Fullname: "Christopher Kelley"},
Article{Name: "Gary",Surname: "Bowling",Fullname: "Brenda Fletcher"},
Article{Name: "Christopher",Surname: "Harper",Fullname: "Peter Harris"},
Article{Name: "Charlene",Surname: "McKenzie",Fullname: "Clifford Lindsay"},
Article{Name: "Sidney",Surname: "Pittman",Fullname: "Jimmy Woodward"},
Article{Name: "Jerome",Surname: "Hall",Fullname: "Ashley Schwartz"},
Article{Name: "Vincent",Surname: "Nichols",Fullname: "Adam Chang"},
Article{Name: "Allen",Surname: "Burnette",Fullname: "Denise Nixon"},
Article{Name: "Sherri",Surname: "Dodson",Fullname: "Dana Finch"},
Article{Name: "Rita",Surname: "Britt",Fullname: "Brooke Coley"},
Article{Name: "Sarah",Surname: "Sanchez",Fullname: "Claire McKenna"},
Article{Name: "Shannon",Surname: "Godwin",Fullname: "Dorothy Fernandez"},
Article{Name: "Ellen",Surname: "Corbett",Fullname: "Nelson Stanley"},
Article{Name: "Frederick",Surname: "Lawson",Fullname: "Wesley McClure"},
Article{Name: "Maria",Surname: "Anthony",Fullname: "Marguerite Shapiro"},
Article{Name: "Julie",Surname: "Song",Fullname: "Ted Aldridge"},
Article{Name: "Joan",Surname: "Burgess",Fullname: "Helen Dickinson"},
Article{Name: "Colleen",Surname: "Murphy",Fullname: "Christopher Kelley"},
Article{Name: "Gary",Surname: "Bowling",Fullname: "Brenda Fletcher"},
Article{Name: "Christopher",Surname: "Harper",Fullname: "Peter Harris"},
Article{Name: "Charlene",Surname: "McKenzie",Fullname: "Clifford Lindsay"},
Article{Name: "Sidney",Surname: "Pittman",Fullname: "Jimmy Woodward"},
Article{Name: "Jerome",Surname: "Hall",Fullname: "Ashley Schwartz"},
Article{Name: "Vincent",Surname: "Nichols",Fullname: "Adam Chang"},
Article{Name: "Allen",Surname: "Burnette",Fullname: "Denise Nixon"},
Article{Name: "Sherri",Surname: "Dodson",Fullname: "Dana Finch"},
Article{Name: "Rita",Surname: "Britt",Fullname: "Brooke Coley"},
Article{Name: "Sarah",Surname: "Sanchez",Fullname: "Claire McKenna"},
Article{Name: "Shannon",Surname: "Godwin",Fullname: "Dorothy Fernandez"},
Article{Name: "Ellen",Surname: "Corbett",Fullname: "Nelson Stanley"},
Article{Name: "Frederick",Surname: "Lawson",Fullname: "Wesley McClure"},
Article{Name: "Maria",Surname: "Anthony",Fullname: "Marguerite Shapiro"},
Article{Name: "Julie",Surname: "Song",Fullname: "Ted Aldridge"},
Article{Name: "Joan",Surname: "Burgess",Fullname: "Helen Dickinson"},
Article{Name: "Colleen",Surname: "Murphy",Fullname: "Christopher Kelley"},
Article{Name: "Gary",Surname: "Bowling",Fullname: "Brenda Fletcher"},
Article{Name: "Christopher",Surname: "Harper",Fullname: "Peter Harris"},
Article{Name: "Charlene",Surname: "McKenzie",Fullname: "Clifford Lindsay"},
Article{Name: "Sidney",Surname: "Pittman",Fullname: "Jimmy Woodward"},
Article{Name: "Jerome",Surname: "Hall",Fullname: "Ashley Schwartz"},
Article{Name: "Vincent",Surname: "Nichols",Fullname: "Adam Chang"},
Article{Name: "Allen",Surname: "Burnette",Fullname: "Denise Nixon"},
Article{Name: "Sherri",Surname: "Dodson",Fullname: "Dana Finch"},
Article{Name: "Rita",Surname: "Britt",Fullname: "Brooke Coley"},
Article{Name: "Sarah",Surname: "Sanchez",Fullname: "Claire McKenna"},
Article{Name: "Shannon",Surname: "Godwin",Fullname: "Dorothy Fernandez"},
Article{Name: "Ellen",Surname: "Corbett",Fullname: "Nelson Stanley"},
Article{Name: "Frederick",Surname: "Lawson",Fullname: "Wesley McClure"},
Article{Name: "Maria",Surname: "Anthony",Fullname: "Marguerite Shapiro"},
Article{Name: "Julie",Surname: "Song",Fullname: "Ted Aldridge"},
Article{Name: "Joan",Surname: "Burgess",Fullname: "Helen Dickinson"},
Article{Name: "Colleen",Surname: "Murphy",Fullname: "Christopher Kelley"},
Article{Name: "Gary",Surname: "Bowling",Fullname: "Brenda Fletcher"},
Article{Name: "Christopher",Surname: "Harper",Fullname: "Peter Harris"},
Article{Name: "Charlene",Surname: "McKenzie",Fullname: "Clifford Lindsay"},
Article{Name: "Sidney",Surname: "Pittman",Fullname: "Jimmy Woodward"},
Article{Name: "Jerome",Surname: "Hall",Fullname: "Ashley Schwartz"},
Article{Name: "Vincent",Surname: "Nichols",Fullname: "Adam Chang"},
Article{Name: "Allen",Surname: "Burnette",Fullname: "Denise Nixon"},
Article{Name: "Sherri",Surname: "Dodson",Fullname: "Dana Finch"},
Article{Name: "Rita",Surname: "Britt",Fullname: "Brooke Coley"},
Article{Name: "Sarah",Surname: "Sanchez",Fullname: "Claire McKenna"},
Article{Name: "Shannon",Surname: "Godwin",Fullname: "Dorothy Fernandez"},
Article{Name: "Ellen",Surname: "Corbett",Fullname: "Nelson Stanley"},
Article{Name: "Frederick",Surname: "Lawson",Fullname: "Wesley McClure"},
Article{Name: "Maria",Surname: "Anthony",Fullname: "Marguerite Shapiro"},
Article{Name: "Julie",Surname: "Song",Fullname: "Ted Aldridge"},
Article{Name: "Joan",Surname: "Burgess",Fullname: "Helen Dickinson"},
Article{Name: "Colleen",Surname: "Murphy",Fullname: "Christopher Kelley"},
Article{Name: "Gary",Surname: "Bowling",Fullname: "Brenda Fletcher"},
Article{Name: "Christopher",Surname: "Harper",Fullname: "Peter Harris"},
Article{Name: "Charlene",Surname: "McKenzie",Fullname: "Clifford Lindsay"},
Article{Name: "Sidney",Surname: "Pittman",Fullname: "Jimmy Woodward"},
Article{Name: "Jerome",Surname: "Hall",Fullname: "Ashley Schwartz"},
Article{Name: "Vincent",Surname: "Nichols",Fullname: "Adam Chang"},
Article{Name: "Allen",Surname: "Burnette",Fullname: "Denise Nixon"},
Article{Name: "Sherri",Surname: "Dodson",Fullname: "Dana Finch"},
Article{Name: "Rita",Surname: "Britt",Fullname: "Brooke Coley"},
Article{Name: "Sarah",Surname: "Sanchez",Fullname: "Claire McKenna"},
Article{Name: "Shannon",Surname: "Godwin",Fullname: "Dorothy Fernandez"},
Article{Name: "Ellen",Surname: "Corbett",Fullname: "Nelson Stanley"},
Article{Name: "Frederick",Surname: "Lawson",Fullname: "Wesley McClure"},
Article{Name: "Maria",Surname: "Anthony",Fullname: "Marguerite Shapiro"},
Article{Name: "Julie",Surname: "Song",Fullname: "Ted Aldridge"},
Article{Name: "Joan",Surname: "Burgess",Fullname: "Helen Dickinson"},
Article{Name: "Colleen",Surname: "Murphy",Fullname: "Christopher Kelley"},
Article{Name: "Gary",Surname: "Bowling",Fullname: "Brenda Fletcher"},
Article{Name: "Christopher",Surname: "Harper",Fullname: "Peter Harris"},
Article{Name: "Charlene",Surname: "McKenzie",Fullname: "Clifford Lindsay"},
Article{Name: "Sidney",Surname: "Pittman",Fullname: "Jimmy Woodward"},
Article{Name: "Jerome",Surname: "Hall",Fullname: "Ashley Schwartz"},
Article{Name: "Vincent",Surname: "Nichols",Fullname: "Adam Chang"},
Article{Name: "Allen",Surname: "Burnette",Fullname: "Denise Nixon"},
Article{Name: "Sherri",Surname: "Dodson",Fullname: "Dana Finch"},
Article{Name: "Rita",Surname: "Britt",Fullname: "Brooke Coley"},
Article{Name: "Sarah",Surname: "Sanchez",Fullname: "Claire McKenna"},
Article{Name: "Shannon",Surname: "Godwin",Fullname: "Dorothy Fernandez"},
Article{Name: "Ellen",Surname: "Corbett",Fullname: "Nelson Stanley"},
Article{Name: "Frederick",Surname: "Lawson",Fullname: "Wesley McClure"},
Article{Name: "Maria",Surname: "Anthony",Fullname: "Marguerite Shapiro"},
Article{Name: "Julie",Surname: "Song",Fullname: "Ted Aldridge"},
Article{Name: "Joan",Surname: "Burgess",Fullname: "Helen Dickinson"},
Article{Name: "Colleen",Surname: "Murphy",Fullname: "Christopher Kelley"},
Article{Name: "Gary",Surname: "Bowling",Fullname: "Brenda Fletcher"},
Article{Name: "Christopher",Surname: "Harper",Fullname: "Peter Harris"},
Article{Name: "Charlene",Surname: "McKenzie",Fullname: "Clifford Lindsay"},
    }
    handleRequests()
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
    //Json new encoder digunakan untuk convert/encode array to a json response
}

func handleRequests(){
	http.HandleFunc("/",homePage)
	//route
	http.HandleFunc("/api",returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000",nil))
}