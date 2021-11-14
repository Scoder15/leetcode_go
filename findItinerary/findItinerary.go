package main

import (
	"fmt"
	"sort"
)

func main() {

	//tickets := [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}
	tickets := [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}

	fmt.Println(findItinerary(tickets))

}

func findItinerary(tickets [][]string) []string {
	res := make([]string, 0)
	track := make([]string, 0)
	target := make(map[string][]string, 0)
	for i := 0; i < len(tickets); i++ {
		target[tickets[i][0]] = append(target[tickets[i][0]], tickets[i][1])
	}
	//fmt.Println("target:", target)
	track = append(track, "JFK")
	backtrack(&res, &track, tickets, "JFK", &target)
	return res

}

func backtrack(res *[]string, track *[]string, tickets [][]string, start string, target *map[string][]string) bool {
	if len(*track) == len(tickets)+1 {
		temp := make([]string, len(*track))
		copy(temp, *track)
		*res = temp
		return true
	}
	//val := target[start]
	sort.Strings((*target)[start])
	//reverse(val)
	//fmt.Println("val:", (*target)[start])
	for _ = range (*target)[start] {
		endPoint := (*target)[start][0]
		//fmt.Println("endPoint:", endPoint)
		(*target)[start] = (*target)[start][1:]
		*track = append(*track, endPoint)
		if backtrack(res, track, tickets, endPoint, target) {
			return true
		}
		*track = (*track)[:len(*track)-1]
		(*target)[start] = append((*target)[start], endPoint)
	}
	return false
}

func reverse(s []string) {

	n := len(s)

	for i := 0; i <= n/2; i++ {
		// temp := s[i]
		// s[i] = s[n-1-i]
		// s[n-1-i] = temp
		s[i], s[n-1-i] = s[n-1-i], s[i]
	}
}

func maxProfit(prices []int) int {
	return 0
}
