package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	old := os.Args[1]
	file1, err := os.Open(old)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file1.Close()
	scanner := bufio.NewScanner(file1)
	previous := make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "|")
		previous[items[1]], _ = strconv.Atoi(items[0])
	}
	new := os.Args[2]
	file2, err := os.Open(new)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file2.Close()
	scanner = bufio.NewScanner(file2)
	var concatenatedText string
	position := 1
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "|")
		stars, _ := strconv.Atoi(items[0])
		lastStars, ok := previous[items[1]]
		if !ok {
			lastStars = 0
		}
		diff := stars - lastStars
		repo := items[1]
		license := items[2]
		if license == "null" {
			license = ""
		} else {
			license = "- " + license
		}
		desc := ""
		if len(items) > 3 {
			desc = items[3]
		}
		if diff > 0 {
			concatenatedText += fmt.Sprintf(
				"|%d|[**%s**](https://github.com/%s)<br>%s %s|%d<br><sup>(+%d)</sup>|\n",
				position, repo, repo, desc, license, stars, diff)
		} else {
			concatenatedText += fmt.Sprintf(
				"|%d|[**%s**](https://github.com/%s)<br>%s %s|%d|\n",
				position, repo, repo, desc, license, stars)
		}
		position += 1
		// if position > 100 {
		// 	break
		// }
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	fmt.Println("| |Repository|Stars|")
	fmt.Println("|:---:|:---|:---:|")
	fmt.Println(concatenatedText)
	now := time.Now()
	day := now.Day()
	month := now.Month()
	year := now.Year()
	date := fmt.Sprintf("%d/%s/%d", day, month, year)
	fmt.Printf("<sub>This list is compiled automatically using Go, Github Actions and the Github API and was last updated on %s.</sub>\n", date)
}
