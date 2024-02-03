package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Subjects struct{
	Title  string `json:"Title"`
	Author string `json:"Author"`
}

type subjectList []Subjects

type Response struct {
	Persons []Person `json:"students"`
}

type Person struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type SingleStudent []Student

type Student struct {
	Id        string `json:"Id"`
	FirstName string `json:"First_name"`
	LastName  string `json:"Last_name"`
	Age 	  string `json:"Age"`
	YearOfStudy string `json:"Year_of_study"` // Заменил Hobby на YearOfStudy
}

func main(){
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/health_check", HealthCheck).Methods("GET")
	router.HandleFunc("/students", Students).Methods("GET")
	router.HandleFunc("/students/{FirstName}", getOneStudent).Methods("GET")
	router.HandleFunc("/", homeLink)

	//start and listen to requests
	http.ListenAndServe(":8080", router)
}

var subjects = subjectList{
	{
		Title:  "Mathematics",
		Author: "John Smith",
	},
	{
		Title:  "History",
		Author: "Jane Doe",
	},
	{
		Title:  "Science",
		Author: "Albert Einstein",
	},
}

func homeLink(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, Welcome to my page!. I like studying different subjects. Here is the list of subjects for learning.")
	json.NewEncoder(w).Encode(subjects)
}

func HealthCheck(w http.ResponseWriter, r *http.Request){
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Description: This app  used to give  information about me and student who study with me. Welcome to my app;)\nAuthor: Moldash Dariya")
}

func getOneStudent(w http.ResponseWriter, r *http.Request){
	studentName := mux.Vars(r)["FirstName"]

	for _, singleStudent := range students {
		if singleStudent.FirstName == studentName {
			json.NewEncoder(w).Encode(singleStudent)
		}
	}
}

func Students(w http.ResponseWriter, r *http.Request){
	log.Println("entering students end point")
	var response Response 
	studentsList := prepareResponse()

	response.Persons = studentsList

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return 
	}

	w.Write(jsonResponse)
}

var students = SingleStudent{
	{
		Id: "1",       
		FirstName: "Amina", 
		LastName: "Mamyrbekova",
		Age: "18",
		YearOfStudy: "2", 
	},
	{
		Id: "2",       
		FirstName: "Dina", 
		LastName: "Kim",
		Age: "19",
		YearOfStudy: "3",
	},
	{
		Id: "3",       
		FirstName: "Mira", 
		LastName: "Kanav",
		Age: "19",
		YearOfStudy: "1",
	},
	{
		Id: "4",       
		FirstName: "Sun", 
		LastName: "Zhemin",
		Age: "18",
		YearOfStudy: "4",
	},
}

func prepareResponse() []Person{
	var studentsList []Person

	var person Person
	person.Id = 1
	person.FirstName = "Amina"
	person.LastName = "Mamyrbekova"
	studentsList = append(studentsList, person)

	person.Id = 2
	person.FirstName = "Dina"
	person.LastName = "Kim"
	studentsList = append(studentsList, person)

	person.Id = 3
	person.FirstName = "Mira"
	person.LastName = "Kanav"
	studentsList = append(studentsList, person)

	person.Id = 4
	person.FirstName = "Sun"
	person.LastName = "Zhemin"
	studentsList = append(studentsList, person)
	return studentsList
}
