package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"sort"
)

type Party struct {
	NumOfVotes	int
	WholePart	int
	FractPart	float64
	NumOfSeats	int
}

func number(s string) (int, bool) {
	num, err := strconv.Atoi(s)

	if err != nil {
		return 0, false
	}
	return num, true
}

func GetParties(reader *bufio.Reader) (map[string]Party, []string) {
	var name string
	var names []string

	parties := make(map[string]Party)

	for ; ; {
		name = ""
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		words := strings.Fields(text)
		len_words := len(words)
		for i := 0; i < len_words - 1; i++ {
			if name == "" {
				name += words[i]
			} else {
				name += " " + words[i]
			}
		}
		votes, _ := number(words[len_words - 1])
		_, ok := parties[name]
		if ok == false {
			parties[name] = Party{ 0, 0, 0.0, 0 }
			names = append(names, name)
		}
		party := parties[name]
		party.NumOfVotes += votes
		parties[name] = party
	}
	return parties, names
}

func SeatsForParties(parties map[string]Party, del float64) int {
	countSeats := 0
	for key, value := range parties {
		f := float64(value.NumOfVotes) / float64(del)
		value.WholePart = int(f)
		value.FractPart = f - float64(value.WholePart)
		value.NumOfSeats = int(f)
		parties[key] = value
		countSeats += parties[key].NumOfSeats
	}
	return countSeats
}

func EventOneParty(parties map[string]Party, key string, countSeats *int) {
	p := parties[key]
	p.NumOfSeats += 1
	parties[key] = p
	*countSeats += 1
}

func EventManyParties(parties map[string]Party, keys []string, countSeats *int) {
	var m_keys []int

	m := make(map[int]string)

	for _, value := range keys {
		m[parties[value].NumOfVotes] = value
		m_keys = append(m_keys, parties[value].NumOfVotes)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(m_keys)))
	for _, value := range m_keys {
		if *countSeats == 450 {
			break
		}
		p := parties[m[value]]
		p.NumOfSeats += 1
		parties[m[value]] = p
		*countSeats += 1
	}
}

func AddSeats(parties map[string]Party, countSeats *int) {
	var keys []float64

	m := make(map[float64][]string)

	for key, value := range parties {
		m[value.FractPart] = append(m[value.FractPart], key)
		keys = append(keys, value.FractPart)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(keys)))
	for _, value := range keys {
		if *countSeats == 450 {
			break
		}
		if len(m[value]) == 1 {
			EventOneParty(parties, m[value][0], countSeats)
		} else {
			EventManyParties(parties, m[value], countSeats)
		}
	}
}

func main() {
	var allVotes int

	Reader := bufio.NewReaderSize(os.Stdin, 100000)

	parties, names := GetParties(Reader)

	for _, p := range parties {
		allVotes += p.NumOfVotes
	}

	countSeats := SeatsForParties(parties, float64(allVotes) / 450.0)

	if countSeats < 450 {
		AddSeats(parties, &countSeats)
	}

	for _, value := range names {
		fmt.Println(value, parties[value].NumOfSeats)
	}
}
