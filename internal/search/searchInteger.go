package search

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func SearchIntegersFromString() []string {
	fmt.Println("searching integer ...")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re.FindAllString(text, -1))
	result := re.FindAllString(text, -1)
	return result
}
