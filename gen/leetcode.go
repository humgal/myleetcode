package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// 生成文件
func GenerateFiles(rootPath, titleSlug string) error {
	detail, err := GetProblemItem(titleSlug)
	if err != nil {
		return err
	}

	// 文件夹名
	dirName := trap(detail.Question.QuestionId) + "." + detail.Question.Title
	// 文件夹路径
	dirPath := path.Join(rootPath, dirName)
	err = os.MkdirAll(dirPath, 0777)
	if err != nil {
		return err
	}

	var goContent string
	var cContent string
	var cppContent string
	var rustContent string
	var jsContent string
	for _, snippet := range detail.Question.CodeSnippets {
		if snippet.Lang == "Go" {
			goContent = snippet.Code
			goFile := detail.Question.TitleSlug + ".go"
			// 写入go文件
			err = ioutil.WriteFile(path.Join(dirPath, goFile), []byte("package code\r\n\r\n"+goContent), os.ModePerm)
			if err != nil {
				return err
			}
		} else if snippet.Lang == "C" {
			cContent = snippet.Code
			cFile := detail.Question.TitleSlug + ".c"

			// 写入go文件
			err = ioutil.WriteFile(path.Join(dirPath, cFile), []byte(cContent), os.ModePerm)
			if err != nil {
				return err
			}
		} else if snippet.Lang == "C++" {
			cppContent = snippet.Code
			cppFile := detail.Question.TitleSlug + ".cpp"

			// 写入go文件
			err = ioutil.WriteFile(path.Join(dirPath, cppFile), []byte(cppContent), os.ModePerm)
			if err != nil {
				return err
			}
		} else if snippet.Lang == "Rust" {
			rustContent = snippet.Code
			rustFile := detail.Question.TitleSlug + ".rs"

			// 写入go文件
			err = ioutil.WriteFile(path.Join(dirPath, rustFile), []byte(rustContent), os.ModePerm)
			if err != nil {
				return err
			}
		} else if snippet.Lang == "JavaScript" {
			jsContent = snippet.Code
			jsFile := detail.Question.TitleSlug + ".js"

			// 写入go文件
			err = ioutil.WriteFile(path.Join(dirPath, jsFile), []byte(jsContent), os.ModePerm)
			if err != nil {
				return err
			}
		}
	}

	// 写入md文件
	readmecnFile := "readme_cn.md"
	content_cn := "#### 题目\r\n" + detail.Question.TranslatedContent + "\r\n\r\n #### 题解"
	err = ioutil.WriteFile(path.Join(dirPath, readmecnFile), []byte(content_cn), os.ModePerm)
	if err != nil {
		return err
	}
	readmeFile := "readme.md"
	content := "#### 题目\r\n" + detail.Question.Content + "\r\n\r\n #### 题解"
	err = ioutil.WriteFile(path.Join(dirPath, readmeFile), []byte(content), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// 获取所有需要生成的文件
func GenerateAllFiles(rootPath string) {
	nums, items, err := GetProblemListsTranlation()
	if err != nil {
		return
	}
	fmt.Printf("will generate total:%d questions\n", nums)
	for _, item := range items {
		err = GenerateFiles(rootPath, item.TitleSlug)
		if err != nil {
			fmt.Printf("ERROR:generate file [%s,%s] error:%v\n", item.QuestionId, item.TitleSlug, err)
		} else {
			fmt.Printf("generate question:[%s,%s] success\n", item.QuestionId, item.TitleSlug)
		}
	}
}

func trap(s string) string {
	var result string
	ret := 5 - len(s)
	if ret > 0 {
		for i := 0; i < ret; i++ {
			result += "0"
		}
	}
	result += s
	return result
}
