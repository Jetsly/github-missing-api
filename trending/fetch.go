package trending

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const GITHUB_URL string = "https://github.com"

type TrendUser struct {
	Href     string `json:"href"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
}

type TrendRepository struct {
	Author             string      `json:"author"`
	Name               string      `json:"name"`
	Avatar             string      `json:"avatar"`
	Url                string      `json:"url"`
	Description        string      `json:"description"`
	Language           string      `json:"language"`
	LanguageColor      string      `json:"languageColor"`
	Stars              int         `json:"stars"`
	Forks              int         `json:"forks"`
	CurrentPeriodStars int         `json:"currentPeriodStars"`
	TrendUser          []TrendUser `json:"builtBy"`
}

func removeDefaultAvatarSize(src string) string {
	return src
	// return strings.re(/\?s=.*$/, '');
}

func toInt(s string) int {
	n, _ := strconv.Atoi(strings.TrimSpace(strings.ReplaceAll(s, ",", "")))
	return n
}

func FetchRepositories(language string, since string) []TrendRepository {
	var url = GITHUB_URL + "/trending/" + language + "?since=" + since
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var repositories []TrendRepository
	doc.Find(".Box article.Box-row").Each(func(i int, repo *goquery.Selection) {
		var title = strings.Split(strings.TrimSpace(repo.Find(".h3").Text()), "/")
		var username, repoName = strings.TrimSpace(title[0]), strings.TrimSpace(title[1])
		var relativeUrl = repo.Find(".h3 a").AttrOr("href", "")
		var colorNode = repo.Find(".repo-language-color")
		var langNode = repo.Find("[itemprop=programmingLanguage]")
		var lang string = ""
		var langColor string = ""
		if colorNode.Length() > 0 {
			langColor = strings.ReplaceAll(strings.TrimSpace(colorNode.AttrOr("style", "")), "background-color:", "")
		}
		if langNode.Length() > 0 {
			lang = strings.TrimSpace(langNode.Text())
		}
		var builtBy []TrendUser
		repo.Find(`span:contains("Built by")`).Find(`[data-hovercard-type="user"]`).Each(func(_ int, user *goquery.Selection) {
			builtBy = append(builtBy, TrendUser{
				Username: user.ChildrenFiltered("img").AttrOr("alt", ""),
				Href:     fmt.Sprintf("%s%s", GITHUB_URL, user.AttrOr("href", "")),
				Avatar:   removeDefaultAvatarSize(user.ChildrenFiltered("img").AttrOr("src", "")),
			})
		})

		repositories = append(repositories, TrendRepository{
			Author:             username,
			Name:               repoName,
			Avatar:             fmt.Sprintf("%s/%s.png", GITHUB_URL, username),
			Url:                fmt.Sprintf("%s%s", GITHUB_URL, relativeUrl),
			Description:        strings.TrimSpace(repo.Find("p.my-1").Text()),
			Language:           lang,
			LanguageColor:      langColor,
			Stars:              toInt(repo.Find(`[href="` + relativeUrl + `/stargazers"]`).Text()),
			Forks:              toInt(repo.Find(`[href="` + relativeUrl + `/network/members"]`).Text()),
			CurrentPeriodStars: toInt(strings.Split(repo.Find(".float-sm-right").Text(), " ")[0]),
			TrendUser:          builtBy,
		})

	})
	return repositories
}
