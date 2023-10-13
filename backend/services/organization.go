package services

import (
	"bufio"
	"fmt"
	"os"
)

func GetOrganizations() {

	fmt.Println("Running Get Organizations")

	file, err := os.Open("data/organizations")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
