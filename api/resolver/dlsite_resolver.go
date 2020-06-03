package resolver

import (
	"errors"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// DLsiteResolver is a metadata resolver for DLsite
type DLsiteResolver struct{}

// NewDLsiteResolver returns new DLsiteResolver
func NewDLsiteResolver() *DLsiteResolver {
	return &DLsiteResolver{}
}

// Resolve products metadata of DLsite
func (r *DLsiteResolver) Resolve(url string) (Metadata, error) {
	metadata := Metadata{}

	// データを取得
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return metadata, errors.New("データの取得に失敗しました。")
	}

	// タイトルを取得
	metadata.Title = doc.Find("#work_name a").Text()
	if len(metadata.Title) <= 0 {
		return metadata, errors.New("タイトルの取得に失敗しました。")
	}

	// サムネイルを取得
	var isExists bool
	metadata.Thumbnail, isExists = doc.Find(".slider_item img").Attr("src")
	if !isExists || len(metadata.Thumbnail) <= 0 {
		return metadata, errors.New("サムネイルの取得に失敗しました。")
	}
	metadata.Thumbnail = "https:" + metadata.Thumbnail

	// 年齢制限を取得
	logo, isExists := doc.Find(".logo img").Attr("alt")
	if !isExists || !strings.HasPrefix(logo, "DLsite") {
		return metadata, errors.New("年齢制限の取得に失敗しました。")
	}
	metadata.IsR18 = false
	if logo == "DLsite R18" {
		metadata.IsR18 = true
	}

	// ジャンルを取得
	doc.Find(".work_genre a").Each(func(_ int, s *goquery.Selection) {
		if "アドベンチャー" == s.Find("span").Text() {
			metadata.IsNovelGame = true
		}
	})

	return metadata, nil
}
