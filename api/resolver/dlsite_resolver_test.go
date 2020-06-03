package resolver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEveryoneGame(t *testing.T) {
	url := "https://www.dlsite.com/soft/work/=/product_id/VJ013471.html"
	metadata := Metadata{
		Title:       "キラキラモンスターズ Season1",
		Thumbnail:   "https://img.dlsite.jp/modpub/images2/work/professional/VJ014000/VJ013471_img_main.jpg",
		IsR18:       false,
		IsNovelGame: true,
	}

	dlsiteResolver := NewDLsiteResolver()
	resolvedMetadata, err := dlsiteResolver.Resolve(url)

	if assert.NoError(t, err) {
		assert.Equal(t, metadata, resolvedMetadata)
	}
}

func TestR18Game(t *testing.T) {
	url := "https://www.dlsite.com/pro/work/=/product_id/VJ012243.html"
	metadata := Metadata{
		Title:       "ラズベリーキューブ",
		Thumbnail:   "https://img.dlsite.jp/modpub/images2/work/professional/VJ013000/VJ012243_img_main.jpg",
		IsR18:       true,
		IsNovelGame: true,
	}

	dlsiteResolver := NewDLsiteResolver()
	resolvedMetadata, err := dlsiteResolver.Resolve(url)

	if assert.NoError(t, err) {
		assert.Equal(t, metadata, resolvedMetadata)
	}
}

func TestASMR(t *testing.T) {
	url := "https://www.dlsite.com/maniax/work/=/product_id/RJ241712.html"
	metadata := Metadata{
		Title:       "ふたりがけ催眠ドライイキサポート編",
		Thumbnail:   "https://img.dlsite.jp/modpub/images2/work/doujin/RJ242000/RJ241712_img_main.jpg",
		IsR18:       true,
		IsNovelGame: false,
	}

	dlsiteResolver := NewDLsiteResolver()
	resolvedMetadata, err := dlsiteResolver.Resolve(url)

	if assert.NoError(t, err) {
		assert.Equal(t, metadata, resolvedMetadata)
	}
}
