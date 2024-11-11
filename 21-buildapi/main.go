package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Model for courses -file
type Course struct {
	CourseId string `json:"course_id"`
	CourseName string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`

}

type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}

// fake db
//an slice(array) of Course structs
var courses[]Course


//middleware, helper - file 
func (c *Course ) isEmpty() bool{
	return c.CourseId == "" && c.CourseName == ""

}

func main(){
	fmt.Println("API -  LearnCodeOnline")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "ReactJs", CoursePrice:229, Author:  &Author{Fullname: "Hitesh", Website: "pillar.com"} })
	courses = append(courses, Course{CourseId: "2", CourseName: "Golang", CoursePrice:199, Author:  &Author{Fullname: "Ayobami", Website: "go.dev.com"} })

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses/{id}",updateOneCourse ).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")
 
	//listen to port
	log.Fatal(http.ListenAndServe(":4000", r))

	
}

//controller - file

//serve home route 

func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1> Welcome to API by LearnCodeOnline<h1>")) //write method writes an array

 }


//serve getAllCourses route 
func getAllCourses(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Get all courses...")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grabs id from request
	params := mux.Vars(r)

	//loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}	
	json.NewEncoder(w).Encode("No course found with given id")
	
}

func updateOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create on course")
	w.Header().Set("Content-Type", "application/json")

	//first -grab id from request (query parameter)
	//returns id as key value pairs
	params := mux.Vars(r)

	//loop, id, remove, add wit my ID

	 for index, course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index], courses[index+1:]... )
			var course Course
			_ =json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
		//Todo: send response when id is not found

	 }
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//loop, id, remove (index, index2)


	for index, course := range courses {
		if course.CourseId  == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}

} 