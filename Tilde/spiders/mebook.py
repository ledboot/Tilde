import re

import scrapy
from scrapy.spiders import CrawlSpider

from Tilde.items import MeBookItem


class mebookSpider(CrawlSpider):
    name = "mebookList"
    allowed_domains = ['mebook.cc/']

    start_urls = [
        "http://mebook.cc/page/1",
    ]
    totalPage = 0

    def parse(self, response):
        for book in response.css('ul.list').xpath('./li'):
            url = book.css('div.content').xpath('./h2/a/@href').extract_first()
            title = book.css('div.content').xpath('./h2/a/@title').extract_first()
            cover = book.css('div.img').xpath('./a/img/@src').extract_first()
            publishDate = book.css('div.info::text').extract_first()
            item = MeBookItem(title=title, url=url, cover=cover, publishDate=publishDate)
            yield item

        pageNav = response.css('div.pagenavi')
        currentPage = pageNav.css('span.current::text').extract_first()
        pageNumber = pageNav.css('span.page-numbers::text').extract_first()
        findPage = re.findall(r"共 (.+?) 页", pageNumber)
        if findPage.__len__() == 1:
            self.totalPage = int(findPage[0])
            currentPage = int(currentPage)
            if currentPage <= self.totalPage:
                currentPage += 1
                next_url = response.urljoin(str(currentPage))
                yield scrapy.Request(next_url, callback=self.parse)
