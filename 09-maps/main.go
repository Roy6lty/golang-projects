package main

import "fmt"

func main() {
	fmt.Println("Maps is golang")

	langauges := make(map[string]string) // a Dictionary of strings
	//make(map[keys]values)

	langauges["JS"] = "javascript"
	langauges["RB"] = "Ruby"
	langauges["PY"] = "Python"

	fmt.Println("List of all languages", langauges)
	fmt.Println("JS shorts for:", langauges["JS"])

	delete(langauges, "RB") // delete items from dictionary
	fmt.Println("List of all languages", langauges)

	for key, value := range langauges {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}
