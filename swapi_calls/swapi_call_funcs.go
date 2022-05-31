package swapi_calls

import (
	"encoding/json"
	"log"

	"github.com/justinrlurie/go-http/safe_http"
)

// SWAPI_url_concat is a public function that returns the swapi root url + the parameter to be returned,
// and also checks if the param input is valid. Strictly for root parameter level only
func SWAPI_url_concat (root_param string) (url string) {

	var swapi_params = []string{"people", "planets", "films", "species", "vehicles", "starships"}
	if Contains(swapi_params, root_param) == false {
		log.Fatal("Invalide SWAPI root parameter")
	}

	base_url := "https://swapi.dev/api/"
	
	return base_url	+ root_param
}

// SWAPI_get_characters_from_film is a public function that writes a list of characters
// from a valid SWAPI film paramater to a People data struct. Currently the function 
// makes an HTTP request for each character from the desired film. 
// TODO : Decrease the number of HTTP requests by requesting entire list of characters once 
// and selecting those that are in the movie. Slightly complicated due to pagination of results.
func SWAPI_get_characters_from_film (film_title string, persons *[]swapi_data_structs.Person) {
	// TODO : This function is a bit long, maybe break it into smaller functions?
	var film_title_params = map[string]string {
		"The Phantom Menace" 		: "4",
		"Attack of the Clones" 		: "5",
		"Revenge of the Sith" 		: "6",
		"A New Hope" 				: "1", 
		"The Empire Strikes Back" 	: "2",
		"Return of the Jedi" 		: "3",
	} // This map is sort of reversed. This is because we want to keep referncing by movie name
	  // for clarity, and we need the map to return a numerical value to construct the url for a single movie.

	if Contains(Map_keys_to_slice_string(film_title_params), film_title) == false {
		log.Fatal("Invalid SWAPI film_title")
	} // Call to 2 functions. Outter function checks if film_title is contained somewhere 
	  // in film_title params (i.e. is it a valid film name?). Inner function deals with the messy
	  // reverse map above by converting the keys of the map to a slice of strings.

	base_url := SWAPI_url_concat("films") + "/" + film_title_params[film_title] + "/?format=json"
	// Generate url for the desired movie.

	response, api_get_from_err := safe_http.Safe_http_get(base_url, 5)
	// HTTP request
	if api_get_from_err != nil {
		log.Println("api_get_from_err")
		log.Fatal(api_get_from_err)
	} // Not sure if we need this extra error check

	var film_struct swapi_data_structs.Film
	json.Unmarshal(response, &film_struct)
	// parse response into struct from swapi_data_structs
	
	characters_slice := film_struct.Characters
	// string slice containing urls for the following for loop

	for i := range characters_slice {
		var temp_person_struct swapi_data_structs.Person
		// temporary person struct to store data for one person
		response, api_get_from_err = safe_http.Safe_http_get(characters_slice[i], 10)
		// This HTTP request gets called on each loop (bad)
		json.Unmarshal(response, &temp_person_struct)
		*persons = append(*persons, temp_person_struct)
		// append to variable outside of loop with pointer. This is the functions "output"
	}
}

func SWAPI_get_homeworld_name_and_pop_from_url (homeworld_url string) (planet_name string, planet_pop string) {
	response, api_get_from_err := safe_http.Safe_http_get(homeworld_url, 10)
	// HTTP request

	if api_get_from_err != nil {
		log.Println("api_get_from_err")
		log.Fatal(api_get_from_err)
	} // Not sure if we need this extra error check

	var planet swapi_data_structs.Planet
	json.Unmarshal(response, &planet)

	return planet.Name, planet.Population
}

// Contains is a public function to test if a string is present in a slice
func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// Map_vals_to_slice is a public function to store the <string> keys
// of a map into a slice.
func Map_keys_to_slice_int(map_in map[string]int) (slice_out []string) {
	k := make([]string, 0, len(map_in))

	for  key := range map_in {
   		k = append(k, key)
	}

	return k
}

func Map_keys_to_slice_string(map_in map[string]string) (slice_out []string) {
	k := make([]string, 0, len(map_in))

	for  key := range map_in {
   		k = append(k, key)
	}

	return k
}