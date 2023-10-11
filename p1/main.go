package main

import (
	// "encoding/json"
    "fmt"
    "math/rand"
    // "github.com/go-gota/gota/dataframe"
    // "github.com/go-gota/gota/series"
    // "io/ioutil"
    // "log"
    // "net/http"
    // "os"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}


type language struct {
	NAME string 
	Popularity int64
}

func add_mini(x int, y int) int{
	return x + y
}

func random_price_generator(curr_price float64) float64 {
    var new_price float64
    new_price = curr_price + rand.Float64()
    fmt.Println(rand.Float64())
    return new_price
}
// // A Response struct to map the Entire Response
// type Response struct {
//     Name    string    `json:"name"`
//     Pokemon []Pokemon `json:"pokemon_entries"`
// }

// // A Pokemon Struct to map every pokemon to.
// type Pokemon struct {
//     EntryNo int            `json:"entry_number"`
//     Species PokemonSpecies `json:"pokemon_species"`
// }

// // A struct to map our Pokemon's Species which includes it's name
// type PokemonSpecies struct {
//     Name string `json:"name"`
// }

func main() {
	// m := add_mini(5, 6)
	// fmt.Println(m)

	// var newAlbum album
	// // var myslice []int

	// newAlbum = album{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99}
 //    albums = append(albums, newAlbum)
 //    // fmt.Println(albums)
 //    fmt.Println(rand.Intn(100))

 //    for num := 1; num <= 100; num++ {
 //    	if num == 10 {
 //    		fmt.Println(num)
 //    	}

 //    }

    // nums := []int{2, 3, 4}
    // for i, num := range nums {
    // 	fmt.Println(i, " ", num)
    // }

    var python language
    python = language{NAME: "python", Popularity: 100}
    fmt.Println(python)
    random_price_generator(5)


    // response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
    // if err != nil {
    //     fmt.Print(err.Error())
    //     os.Exit(1)
    // }

    // responseData, err := ioutil.ReadAll(response.Body)
    // if err != nil {
    //     log.Fatal(err)
    // }

    // var responseObject Response
    // json.Unmarshal(responseData, &responseObject)

    // fmt.Println(responseObject.Name)
    // fmt.Println(len(responseObject.Pokemon))

    // for i := 0; i < len(responseObject.Pokemon); i++ {
    //     fmt.Println(responseObject.Pokemon[i].Species.Name)
    // }

}
