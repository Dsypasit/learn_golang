package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	_ "github.com/gocolly/colly"
)

type Player struct {
	Name        string
	Pos         string
	Age         string
	PlayingTime struct {
		MatchPlay int
		Starts    int
		Min       int
	}
	Performance struct {
		Goals      int
		Assists    int
		YellowCard int
		RedCard    int
	}
}

// func removeLower

func main() {
	url := "https://fbref.com/en/comps/9/Premier-League-Stats"
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		r, _ := regexp.Compile(`\/\w+\/squads\/\w+\/.*`)
		// r, _ := regexp.Compile(`https:\/\/fbref.com\/\w+\/squads\/\w+\/.*`)
		fmt.Println(r.MatchString(h.Attr("href")), h.Attr("href"))
		if r.MatchString(h.Attr("href")) {
			fmt.Printf("go web more : %s\n", h.Attr("href"))
			h.Request.Visit(h.Attr("href"))
		}
	})

	c.OnHTML("table", func(e *colly.HTMLElement) {
		selection := e.DOM
		// if strings.Contains()
		table_name := selection.AttrOr("id", "none")
		players := []Player{}
		if strings.Contains(table_name, "stats_standard") {
			selection.Find("tbody>tr").Each(func(i int, s *goquery.Selection) {
				Name := s.Find("[data-stat=player]").AttrOr("csk", "none")
				Pos := s.Find("[data-stat=position]").Text()
				Age := s.Find("[data-stat=age]").Text()
				MatchPlay, _ := strconv.Atoi(s.Find("[data-stat=games]").Text())
				Starts, _ := strconv.Atoi(s.Find("[data-stat=games_starts]").Text())
				Min, _ := strconv.Atoi(s.Find("[data-stat=minutes]").Text())
				Goals, _ := strconv.Atoi(s.Find("[data-stat=goals]").Text())
				Asists, _ := strconv.Atoi(s.Find("[data-stat=assitsts]").Text())
				YellowCard, _ := strconv.Atoi(s.Find("[data-stat=cards_yellow]").Text())
				RedCard, _ := strconv.Atoi(s.Find("[data-stat=cards_red]").Text())
				player := Player{
					Name: Name,
					Pos:  Pos,
					Age:  Age,
					PlayingTime: struct {
						MatchPlay int
						Starts    int
						Min       int
					}{
						MatchPlay: MatchPlay,
						Starts:    Starts,
						Min:       Min,
					},
					Performance: struct {
						Goals      int
						Assists    int
						YellowCard int
						RedCard    int
					}{
						Goals:      Goals,
						Assists:    Asists,
						YellowCard: YellowCard,
						RedCard:    RedCard,
					},
				}
				players = append(players, player)
			})
			for _, p := range players[:5] {
				fmt.Println(p)
			}
		}
	})
	c.Visit(url)
}
