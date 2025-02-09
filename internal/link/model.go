package link

import (
	"go_dev/internal/statistic"
	"math/rand"

	"gorm.io/gorm"
)

type Link struct {
	gorm.Model
	URL   string                `json:"url"`
	Hash  string                `json:"hash" gorm:"uniqueIndex"`
	Stats []statistic.Statistic `json:"stats" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func NewLink(url string) *Link {
	link := &Link{
		URL: url,
	}
	link.GenerateHash()
	return link
}

func (link *Link) GenerateHash() {
	link.Hash = RunStringRuns(10)
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RunStringRuns(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
