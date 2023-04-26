package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// algorithms
type algorithms struct {
	Name     string          `json:"category_slug"`
	User     string          `json:"user_name"`
	ACEasy   int             `json:"ac_easy"`
	ACMedium int             `json:"ac_medium"`
	ACHard   int             `json:"ac_hard"`
	AC       int             `json:"num_solved"`
	Problems []problemStatus `json:"stat_status_pairs"`
}

type problemStatus struct {
	Status     string `json:"status"`
	State      `json:"stat"`
	IsFavor    bool `json:"is_favor"`
	IsPaid     bool `json:"paid_only"`
	Difficulty `json:"difficulty"`
}

// State
type State struct {
	ACs       int    `json:"total_acs"`
	Title     string `json:"question__title"`
	IsNew     bool   `json:"is_new_question"`
	Submitted int    `json:"total_submitted"`
	ID        int    `json:"frontend_question_id"`
	TitleSlug string `json:"question__title_slug"`
}

// Difficulty
type Difficulty struct {
	Level int `json:"level"`
}

func main() {
	url := fmt.Sprintf("https://leetcode.com/api/problems/Algorithms/")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	res := new(algorithms)
	err = json.Unmarshal(body, &res)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Problems)

	// var data map[string]interface{}
	// err = json.NewDecoder(resp.Body).Decode(&data)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// problems := data["stat_status_pairs"].([]interface{})
	// for _, problem := range problems {
	// 	problemData := problem.(map[string]interface{})
	// 	title := problemData["stat"].(map[string]interface{})["question__title"].(string)
	// 	fmt.Println(title)
	// }
}
