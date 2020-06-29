package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	/*os.Stdin.Close()
	//original.txt
	in := bufio.NewScanner(os.Stdin)

	var lines int

	for in.Scan() {
		lines++
	}

	fmt.Printf("There are %d line(s)\n", lines)

	if err := in.Err(); err != nil {
		fmt.Println("ERROR: ", err)
	}

	//fmt.Println("Scanned Text: ", in.Text())
	//fmt.Println("Scanned Bytes: ", in.Bytes())*/

	//original2.txt

	/*args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Please type a search word.")
		return
	}
	query := args[0]
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)

	words := make(map[string]bool)
	for in.Scan() {
		word := strings.ToLower(in.Text())
		if len(word) > 2 {
			words[word] = true
		}
	}

	if words[query] {
		fmt.Printf("The input contains %q.\n", query)
		return
	}
	fmt.Printf("The input does not contain %q.\n", query)

	*/
	var (
		sum     map[string]int
		domains []string
		total   int
		lines   int
	)

	sum = make(map[string]int)
	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		lines++

		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Printf("wrong input: %v (line #%d)\n", fields, lines)
			return
		}
		domain := fields[0]
		visits, err := strconv.Atoi(fields[1])
		if visits < 0 || err != nil {
			fmt.Printf("wrong input: %q (line #%d)\n", fields[1], lines)
			return
		}

		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}
		total += visits
		sum[domain] += visits
	}
	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	sort.Strings(domains)
	for _, domain := range domains {
		visits := sum[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}
	fmt.Printf("\n%-30s %10d\n", "TOTAL", total)

	if err := in.Err(); err != nil {
		fmt.Println("> Err:", err)
	}

}
