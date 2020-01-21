package sy9d

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/ledboot/Tilde/initial"
	"github.com/ledboot/Tilde/logger"
	"github.com/ledboot/Tilde/models"
	"strconv"
	"strings"
	"time"
)

func Run() {
	total := getPageNum()
	logger.Infof("total page num : %d", total)
	getItem(2)
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
		colly.Debugger(&debug.LogDebugger{}),
		colly.Async(true),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*sy9d.com",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	c.OnHTML("table#threadlisttableid", func(e *colly.HTMLElement) {
		e.ForEach("tbody", func(i int, element *colly.HTMLElement) {
			if strings.Contains(element.Attr("id"), "normalthread") {
				href := element.ChildAttr("div.tl_ct>a.s", "href")
				title := element.ChildText("div.tl_ct>a")
				publishTime := element.ChildText("div.info>em>span")
				sum := sha1.Sum([]byte(href))
				hash := hex.EncodeToString(sum[:])
				count, _ := initial.GetMongoDB().C("sy9d").Find(fmt.Sprintf("hash=%s", hash)).Count()
				if count == 0 {
					item := &models.Sy9d{
						Hash:        hash,
						Title:       title,
						Url:         href,
						PublishTime: publishTime,
						CreatedAt:   time.Now(),
						UpdatedAt:   time.Now(),
					}
					logger.Info(item)
					err := initial.GetMongoDB().C("sy9d").Insert(item)
					if err != nil {
						logger.Error(err)
					}
				}

			}
		})
	})

	for i := 1; i <= total; i++ {
		c.Visit(fmt.Sprintf("%s%d", pageUrl, i))
	}
	c.Wait()
}
