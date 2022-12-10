package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type GraphQL struct {
	OperationName string            `json:"operationName"`
	Query         string            `json:"query"`
	Variables     map[string]string `json:"variables"`
}

type Data struct {
	Data interface{} `json:"data"`
}

// graphql 查询
func GetGraphql(data interface{}) ([]byte, error) {
	url := "https://leetcode-cn.com/graphql"
	var bf *bytes.Buffer

	switch t := data.(type) {
	case string:
		bf = bytes.NewBufferString(t)
	case []byte:
		bf = bytes.NewBuffer(t)
	}

	bts, err := NewRequest(http.MethodPost, url, bf)
	return bts, err
}

type TranslationList struct {
	Translations []TranslationItem `json:"allQuestionsBeta"`
}

// 翻译元数据
type TranslationItem struct {
	TitleSlug  string `json:"titleSlug"`  // 问题英文标题
	Difficult  string `json:"difficulty"` // 难度
	QuestionId string `json:"questionId"` // 问题id
	Title      string `json:"title"`      // 问题标题
}

// 查询题目翻译
func GetProblemListsTranlation() (int, []TranslationItem, error) {
	g := GraphQL{OperationName: "allQuestions", Query: "query allQuestions {  allQuestionsBeta {    ...questionSummaryFields    __typename  }}fragment questionSummaryFields on QuestionNode {  title  titleSlug  translatedTitle  questionId  questionFrontendId  status  difficulty  isPaidOnly  categoryTitle  __typename}"}
	body, err := json.Marshal(g)
	if err != nil {
		return 0, nil, err
	}
	resp, err := GetGraphql(body)
	if err != nil {
		return 0, nil, err
	}

	var data Data
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return 0, nil, err
	}

	var t TranslationList
	ret, err := json.Marshal(data.Data)
	if err != nil {
		return 0, nil, err
	}
	err = json.Unmarshal(ret, &t)
	if err != nil {
		return 0, nil, err
	}
	return len(t.Translations), t.Translations, nil
}

type QuestionDetail struct {
	Question Question `json:"question"`
}

type Question struct {
	TranslationItem
	Content           string        `json:"content"`
	TranslatedContent string        `json:"translatedContent"`
	CodeSnippets      []CodeSnippet `json:"codeSnippets"`
}

type CodeSnippet struct {
	Lang     string `json:"lang"`     // 编程语言
	LangSlug string `json:"langSlug"` // 小写编程语言
	Code     string `json:"code"`     // 代码
}

// 查询问题详情
func GetProblemItem(titleSlug string) (*QuestionDetail, error) {
	g := GraphQL{OperationName: "questionData",
		Query:     "query questionData($titleSlug: String!) {  question(titleSlug: $titleSlug) {    questionId    questionFrontendId    boundTopicId    title    titleSlug    content    translatedTitle    translatedContent    isPaidOnly    difficulty    likes    dislikes    isLiked    similarQuestions    contributors {      username      profileUrl      avatarUrl      __typename    }    langToValidPlayground    topicTags {      name      slug      translatedName      __typename    }    companyTagStats    codeSnippets {      lang      langSlug      code      __typename    }    stats    hints    solution {      id      canSeeDetail      __typename    }    status    sampleTestCase    metaData    judgerAvailable    judgeType    mysqlSchemas    enableRunCode    envInfo    book {      id      bookName      pressName      source      shortDescription      fullDescription      bookImgUrl      pressImgUrl      productUrl      __typename    }    isSubscribed    isDailyQuestion    dailyRecordStatus    editorType    ugcQuestionId    __typename  }}",
		Variables: map[string]string{"titleSlug": titleSlug}}
	body, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	resp, err := GetGraphql(body)
	if err != nil {
		return nil, err
	}

	var data Data
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return nil, err
	}

	var t QuestionDetail
	ret, err := json.Marshal(data.Data)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(ret, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
