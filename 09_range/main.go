package main

import "fmt"

func main() {
	ids := []int{33, 46, 51, 29, 14, 62}

	//Loop through Ids
	for i, id := range ids {
		fmt.Printf("%d ID=%d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID=%d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println(sum)

	//Range with Maps
	emails := map[string]string{"Bob": "bob@gmail.com", "Mike": "mike@gmail.com"}
	for k, v := range emails {
		fmt.Printf("%s : %s\n", k, v)
	}

	for k := range emails {
		fmt.Printf("Name: %s\n", k)
	}
}
