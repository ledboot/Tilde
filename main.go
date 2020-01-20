package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gocolly/colly"
	"github.com/ledboot/tilde/initial"
	"github.com/ledboot/tilde/logger"
	"github.com/ledboot/tilde/models"
	_ "github.com/ledboot/tilde/routers"
	"strconv"
	"strings"
	"time"
)

func main() {

	beego.Run()
	//total := getPageNum()
	//logger.Infof("total page num : %d", total)
	//getItem(total)
}

func getPageNum() int {
	total := 0
	var err error
	c := colly.NewCollector()
	c.OnHTML("span#fd_page_top>div.pg>a.last", func(element *colly.HTMLElement) {
		total, err = strconv.Atoi(strings.Split(element.Text, " ")[1])
	})

	c.Visit("http://www.sy9d.com/forum.php?mod=forumdisplay&fid=179")
	c.Wait()
	return total
}

func getItem(total int) {
	pageUrl := "http://www.sy9d.com/forum.php?mod=forumdisplay&fid=179&page="
	c := colly.NewCollector(
		colly.MaxDepth(1),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		Delay:       5 * time.Second,
		RandomDelay: 15 * time.Second,
	})

	c.OnHTML("table#threadlisttableid", func(e *colly.HTMLElement) {
		e.ForEach("tbody", func(i int, element *colly.HTMLElement) {
			if strings.Contains(element.Attr("id"), "normalthread") {
				href := element.ChildAttr("div.tl_ct>a.s", "href")
				title := element.ChildText("div.tl_ct>a")
				sum := sha1.Sum([]byte(href))
				hash := hex.EncodeToString(sum[:])
				count, _ := initial.GetMongoDB().C("sy9d").Find(fmt.Sprintf("hash=%s", hash)).Count()
				if count == 0 {
					item := &models.Sy9d{
						Hash:  hash,
						Title: title,
						Url:   href,
					}
					err := initial.GetMongoDB().C("sy9d").Insert(item)
					if err != nil {
						logger.Error(err)
					}
				}

			}
		})
	})

	c.Visit(fmt.Sprintf("%s%d", pageUrl, 1))

	//for i := 1; i <= total; i++ {
	//	c.Visit(fmt.Sprintf("%s%d", pageUrl, i))
	//}
	c.Wait()
}
