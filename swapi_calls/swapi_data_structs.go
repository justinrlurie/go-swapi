// This file contains all of the struct types to store json data
// Use the following link to quickly conver from JSON to formatted go structs:
// https://mholt.github.io/json-to-go/

package swapi_data_structs

import "time"

// ROOT REQUESTS STRUCTS:

type People struct { 		// https://swapi.dev/api/people/?format=json
	Count    int         	`json:"count"`
	Next     string      	`json:"next"`
	Previous interface{} 	`json:"previous"`
	Results  []Person	 	`json:"results"`
}

type Planets struct { 		// https://swapi.dev/api/planets/?format=json
	Count    int         	`json:"count"`
	Next     string      	`json:"next"`
	Previous interface{} 	`json:"previous"`
	Results  []Planet	 	`json:"results"`
}

type Films struct {			// https://swapi.dev/api/films/?format=json
	Count    int         	`json:"count"`
	Next     interface{} 	`json:"next"`
	Previous interface{} 	`json:"previous"`
	Results  []Film		 	`json:"results"`
}

type Species struct {		// https://swapi.dev/api/species/?format=json
	Count    int         	`json:"count"`
	Next     interface{} 	`json:"next"`
	Previous interface{} 	`json:"previous"`
	Results  []Specie	 	`json:"results"`
}

type Vehicles struct {		// https://swapi.dev/api/vehicles/?format=json
	Count    int         	`json:"count"`
	Next     interface{} 	`json:"next"`
	Previous interface{} 	`json:"previous"`
	Results  []Vehicle	 	`json:"results"`
}

type Starships struct {		// https://swapi.dev/api/starships/?format=json
	Count    int         	`json:"count"`
	Next     interface{} 	`json:"next"`
	Previous interface{} 	`json:"previous"`
	Results  []Starship	 	`json:"results"`
}

// END ROOT REQUEST STRUCTS

// MEMBER REQUEST STRUCTS:

type Person struct {
	Name      				string        	`json:"name"`
	Height    				string        	`json:"height"`
	Mass      				string        	`json:"mass"`
	HairColor 				string        	`json:"hair_color"`
	SkinColor 				string        	`json:"skin_color"`
	EyeColor  				string        	`json:"eye_color"`
	BirthYear 				string        	`json:"birth_year"`
	Gender    				string        	`json:"gender"`
	Homeworld 				string        	`json:"homeworld"`
	Films     				[]string      	`json:"films"`
	Species   				[]interface{} 	`json:"species"`
	Vehicles  				[]string      	`json:"vehicles"`
	Starships 				[]string      	`json:"starships"`
	Created   				time.Time     	`json:"created"`
	Edited    				time.Time     	`json:"edited"`
	URL       				string        	`json:"url"`
}

type Planet struct {
	Name           			string    		`json:"name"`
	RotationPeriod 			string    		`json:"rotation_period"`
	OrbitalPeriod  			string    		`json:"orbital_period"`
	Diameter       			string    		`json:"diameter"`
	Climate        			string    		`json:"climate"`
	Gravity        			string    		`json:"gravity"`
	Terrain        			string    		`json:"terrain"`
	SurfaceWater   			string    		`json:"surface_water"`
	Population     			string    		`json:"population"`
	Residents      			[]Person  		`json:"residents"`
	Films          			[]string  		`json:"films"`
	Created        			time.Time 		`json:"created"`
	Edited         			time.Time 		`json:"edited"`
	URL            			string    		`json:"url"`
} 

type Film struct {
	Title        			string    		`json:"title"`
	EpisodeID    			int       		`json:"episode_id"`
	OpeningCrawl 			string    		`json:"opening_crawl"`
	Director     			string    		`json:"director"`
	Producer     			string    		`json:"producer"`
	ReleaseDate  			string    		`json:"release_date"`
	Characters   			[]string  		`json:"characters"`
	Planets      			[]string  		`json:"planets"`
	Starships    			[]string  		`json:"starships"`
	Vehicles     			[]string  		`json:"vehicles"`
	Species      			[]string  		`json:"species"`
	Created      			time.Time 		`json:"created"`
	Edited       			time.Time 		`json:"edited"`
	URL          			string    		`json:"url"`
}

type Specie struct {
	Name            		string    		`json:"name"`
	Classification  		string    		`json:"classification"`
	Designation     		string    		`json:"designation"`
	AverageHeight   		string    		`json:"average_height"`
	SkinColors      		string    		`json:"skin_colors"`
	HairColors      		string    		`json:"hair_colors"`
	EyeColors       		string    		`json:"eye_colors"`
	AverageLifespan 		string    		`json:"average_lifespan"`
	Homeworld       		string    		`json:"homeworld"`
	Language        		string    		`json:"language"`
	People          		[]string  		`json:"people"`
	Films           		[]string  		`json:"films"`
	Created         		time.Time 		`json:"created"`
	Edited          		time.Time 		`json:"edited"`
	URL             		string    		`json:"url"`
}

type Vehicle struct {
	Name                 	string        	`json:"name"`
	Model                	string        	`json:"model"`
	Manufacturer         	string        	`json:"manufacturer"`
	CostInCredits        	string        	`json:"cost_in_credits"`
	Length               	string        	`json:"length"`
	MaxAtmospheringSpeed 	string        	`json:"max_atmosphering_speed"`
	Crew                 	string        	`json:"crew"`
	Passengers           	string        	`json:"passengers"`
	CargoCapacity        	string        	`json:"cargo_capacity"`
	Consumables          	string        	`json:"consumables"`
	VehicleClass         	string        	`json:"vehicle_class"`
	Pilots               	[]interface{} 	`json:"pilots"`
	Films                	[]string      	`json:"films"`
	Created              	time.Time     	`json:"created"`
	Edited               	time.Time     	`json:"edited"`
	URL                  	string        	`json:"url"`
}

type Starship struct {
	Name                 	string        	`json:"name"`
	Model                	string        	`json:"model"`
	Manufacturer         	string        	`json:"manufacturer"`
	CostInCredits        	string        	`json:"cost_in_credits"`
	Length               	string        	`json:"length"`
	MaxAtmospheringSpeed 	string        	`json:"max_atmosphering_speed"`
	Crew                 	string        	`json:"crew"`
	Passengers           	string        	`json:"passengers"`
	CargoCapacity        	string        	`json:"cargo_capacity"`
	Consumables          	string        	`json:"consumables"`
	HyperdriveRating     	string        	`json:"hyperdrive_rating"`
	Mglt                 	string        	`json:"MGLT"`
	StarshipClass        	string        	`json:"starship_class"`
	Pilots               	[]interface{} 	`json:"pilots"`
	Films                	[]string      	`json:"films"`
	Created              	time.Time     	`json:"created"`
	Edited               	time.Time     	`json:"edited"`
	URL                  	string        	`json:"url"`
}

//END MEMBER REQUEST STRUCTS