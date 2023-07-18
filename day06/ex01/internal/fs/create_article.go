package fs

import (
	"day06/ex01/internal/database"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateArticle(name string) error {
	if !isCorrectArticle(name) {
		return errors.New("article already exists")
	}

	mdFile := fmt.Sprintf("./ui/md/%s", name)
	if _, err := os.Create(mdFile); err != nil {
		log.Println(err)
	}

	htmlName := strings.Split(name, ".md")[0] + ".html"
	htmlFile := fmt.Sprintf("./ui/html/%s", htmlName)
	if _, err := os.Create(htmlFile); err != nil {
		log.Println(err)
	}

	return nil
}

func isCorrectArticle(link string) bool {
	articles, err := database.GetArticles()
	if err != nil {
		log.Println(err)
	}
	for current, _ := range articles {
		if current == link {
			return false
		}
	}
	return true
}
