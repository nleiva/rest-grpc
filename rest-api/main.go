// Go SwaggerUI example
//
// This documentation describes a Person API.
//
//     Schemes: http
//     Version: 0.0.1
//     License: BSD-3-Clause https://opensource.org/licenses/BSD-3-Clause
//     Contact: Nicolas Leiva <nleiva@cisco.com> https://nleiva.com
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Person is a person :-)
// swagger:response personResp
type Person struct {
	// in:body
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

// swagger:response peopleResp
type peopleResp struct {
	// in:body
	People []Person `json:"people,omitempty"`
}

// swagger:response personReq
type personReq struct {
	// in:body
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

var people []Person

// GetPersonEndpoint returns the person for the provided id.
// swagger:operation GET /people/{id} person GetPersonEndpoint
//
// Returns the person for the provided id.
// ---
// parameters:
// - name: id
//   in: path
//   description: ID of the person.
//   required: true
//   type: string
// responses:
//   '200':
//     "$ref": "#/responses/personResp"
func GetPersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

// GetPeopleEndpoint returns the list of current people.
// swagger:route GET /people people GetPeopleEndpoint
//
// Returns the list of current people.
// ---
// responses:
//   '200':
//     "$ref": "#/responses/peopleResp"
func GetPeopleEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// CreatePersonEndpoint creates a new person.
// swagger:operation POST /people/{id} person CreatePersonEndpoint
//
// Creates a new person.
// ---
// parameters:
// - name: id
//   in: path
//   description: ID of the person.
//   required: true
//   type: string
// - name: info
//   in: body
//   description: Name and Email of the person.
//   required: true
//   schema:
//     "$ref": "#/responses/personReq"
// responses:
//   '200':
//     "$ref": "#/responses/peopleResp"
func CreatePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// DeletePersonEndpoint creates a new person.
// swagger:operation DELETE /people/{id} person CreatePersonEndpoint
//
// Creates a new person.
// ---
// parameters:
// - name: id
//   in: path
//   description: ID of the person.
//   required: true
//   type: string
// responses:
//   '200':
//     "$ref": "#/responses/peopleResp"
func DeletePersonEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Name: "John", Email: "john@cisco.com"})
	people = append(people, Person{ID: "2", Name: "Jane", Email: "jane@cisco.com"})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./dist/")))
	log.Fatal(http.ListenAndServe(":8000", router))
}

/*
For the person api, we'd need these:

/people (GET) -> All persons in the database
/people/{id} (GET) -> Get a single person
/people/{id} (POST) -> Create a new person
/people/{id} (DELETE) -> Delete a person

http://[::1]:8000/people
http://[::1]:8000/people/1

*/
