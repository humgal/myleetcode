package main

type Title struct {
	Question_id           int    `json:"question_id"`
	Question__title       string `json:"question__title"`
	Question__title_slug  string `json:"question__title_slug"`
	Question__hide        bool   `json:"question__hide"`
	Total_acs             int    `json:"total_acs"`
	Total_submitted       int    `json:"total_submitted"`
	Total_column_articles int    `json:"total_column_articles"`
	Frontend_question_id  string `json:"frontend_question_id"`
	Is_new_question       bool   `json:"is_new_question"`
}

type Problems struct {
	Stat      Title  `json:"stat"`
	Status    string `json:"status"`
	Paid_only bool   `json:"paid_only"`
	Is_favor  bool   `json:"is_favor"`
	Frequency int    `json:"frequency"`
	Progress  int    `json:"progress"`
}

func main() {
	GenerateAllFiles("../leetcode")
}
