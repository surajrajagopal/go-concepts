package mapstring

import "fmt"

func AddMap(emp map[string]int) {
	emp["satya"] = 100
	emp["abani"] = 300
	fmt.Println("Add element to map :", emp)
}

func GetMapValue(emp map[string]int) {
	fmt.Printf("Get element from map : %d\n", emp["satya"])
}

func UpdateMapValues(emp map[string]int) {
	emp["abani"] = 400
	fmt.Println("Update element to map :", emp)
}

func DeleteMapKey(emp map[string]int) {
	delete(emp, "satya")
	fmt.Println("delete key from map :", emp)
}

func IterateMap(emp map[string]int) {
	for k, v := range emp {
		fmt.Printf("key %s, value %d\n", k, v)
	}
}

func KeyExists(emp map[string]int) {
	for k, _ := range emp {
		if value, ok := emp[k]; ok {
			fmt.Printf("key is %s and it's value %d\n", k, value)
		} else {
			fmt.Print("key not found")
		}
	}
}

func MapOfStruct() error {
	type empDetails struct {
		name  string
		empID int
	}
	emp1 := empDetails{
		name:  "satya",
		empID: 10,
	}
	emp2 := empDetails{
		name:  "satyam",
		empID: 100,
	}

	emp := map[string]*empDetails{
		"1": &emp1,
		"2": &emp2,
	}
	fmt.Println("Map of Struct")
	for k, v := range emp {
		fmt.Println(k, v)
	}
	return nil
}
