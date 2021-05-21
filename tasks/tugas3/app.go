package main

import (
	"fmt"
)

// A. DEKLARASI TYPE
type (
	tag struct {
		topic string
		banyak int
	}
	tabTopic [99]string
	tabTag 	 [99]tag
)

func main() {
	var TabTopic = tabTopic{}
	var TabTag = tabTag{}

	// C. MENGISI ARRAY DAN FREQUENCY
	input(TabTopic, &TabTag)

	// D. SUB PROGRAM MENCARI TRENDING TOPIC
	trendingTopic(&TabTag)
}


func trendingTopic(tt *tabTag) {
	trending := tt[0]

	for i := 1; i < len(tt); i++ {
		// B. KEBERADAAN TOPIC
		if tt[i].banyak > trending.banyak {
			trending = tt[i]
		}
	}

	fmt.Println(trending.topic)
}

func input(tt tabTopic, t *tabTag) {
	var inputTopic string
	var newTopic  [len(tt)]string
	l:
		for i := 0; i < len(tt); i++ {
				fmt.Scan(&inputTopic)
				if inputTopic == "stop" {
					break l
				} else {
					newTopic[i] = inputTopic
				}
		}

	duplicate_frequency := make(map[string]int)
	for _, item := range newTopic {
		// check if the item/element exist in the duplicate_frequency map
		_, exist := duplicate_frequency[item]
		if exist {
			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
		} else {
			duplicate_frequency[item] = 1 // else start counting from 1
		}
	}
	for k, _ := range duplicate_frequency {
		if k == ""{delete(duplicate_frequency, k)}
	}
	i := 0
	for key,value := range duplicate_frequency{
		tag := tag{topic:  key, banyak: value}
		t[i] = tag
		i++
	}
}