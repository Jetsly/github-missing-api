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

type TrendRepo struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Description string `json:"description"`
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

type TrendDeveloper struct {
	Username string    `json:"username"`
	Name     string    `json:"name"`
	Type     string    `json:"type"`
	Url      string    `json:"url"`
	Avatar   string    `json:"avatar"`
	Repo     TrendRepo `json:"repo"`
}

func removeDefaultAvatarSize(src string) string {
	if len(src) == 0 {
		return src
	}
	// re3, _ := regexp.Compile("?s=.*$")
	// return re3.ReplaceAllString(src, "")
	return src
}

func toInt(s string) int {
	n, _ := strconv.Atoi(strings.TrimSpace(strings.ReplaceAll(s, ",", "")))
	return n
}

func createClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	return &http.Client{Transport: tr}
}

func fetchDom(url string) *goquery.Document {
	client := createClient()
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
	return doc
}

func FetchRepositories(language string, since string) []TrendRepository {
	var url = GITHUB_URL + "/trending/" + language + "?since=" + since
	doc := fetchDom(url)
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

func FetchDevelopers(language string, since string) []TrendDeveloper {
	var url = GITHUB_URL + "/trending/developers/" + language + "?since=" + since
	doc := fetchDom(url)
	var developers []TrendDeveloper
	doc.Find(".Box article.Box-row").Each(func(i int, user *goquery.Selection) {
		var relativeUrl = user.Find(".h3 a").AttrOr("href", "")
		var repo = user.Find(".mt-2 > article")
		//   $reuserpo.find('svg').remove();
		developers = append(developers, TrendDeveloper{
			Username: relativeUrl[1:],
			Name:     strings.TrimSpace(user.Find(".h3 a").Text()),
			Type:     user.Find("img").Parent().AttrOr("data-hovercard-type", ""),
			Url:      fmt.Sprintf("%s%s", GITHUB_URL, relativeUrl),
			Avatar:   removeDefaultAvatarSize(user.Find("img").AttrOr("src", "")),
			Repo: TrendRepo{
				Name:        strings.TrimSpace(repo.Find("a").Text()),
				Description: strings.TrimSpace(repo.Find(".f6.mt-1").Text()),
				Url:         fmt.Sprintf("%s%s", GITHUB_URL, repo.Find("a").AttrOr("href", "")),
			},
		})
	})
	return developers
}
