package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var places map[string]Vertex

func main() {
	// default declaration
	places = make(map[string]Vertex)
	fmt.Println(places)
	places["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	places["Kremlin"] = Vertex{
		55.7520, 37.6175,
	}
	fmt.Println(places)
	// map literals
	var morePlaces = map[string]Vertex{
		"Google": Vertex{
			37.42204, -122.08408,
		},
		"Bell Labs": {40.68433, -74.39967}, // notice that the type is omitted
	}
	fmt.Println(morePlaces)

	// inserting into a map
	morePlaces["Admiralty Building"] = Vertex{59.9375, 30.3086}
	fmt.Println(morePlaces)

	// updating a value in a map
	morePlaces["Google"] = Vertex{37, -122,}
	fmt.Println(morePlaces)

	// deleting a value from a map
	delete(morePlaces, "Google")
	fmt.Println(morePlaces)

	// checking a key for presence in map
	v, ok := morePlaces["Google"]
	fmt.Println("The value:", v, "Present?", ok)
	v, ok = morePlaces["Admiralty Building"]
	fmt.Println("The value:", v, "Present?", ok)
}