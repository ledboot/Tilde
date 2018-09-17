# -*- coding: utf-8 -*-
import scrapy
from scrapy.linkextractors import LinkExtractor
from scrapy.spiders import CrawlSpider, Rule

from Tilde.items import Sy9MovieItem, Sy9MovieItemDetail


class Sy9MovieSpider(CrawlSpider):
    name = 'sy9-movie'
    allowed_domains = ['sy9d.com']
    start_urls = [
        "http://www.sy9d.com/forum-179-1.html",
    ]

    custom_settings = {
        'ITEM_PIPELINES': {
            'demo.pipelines.Sy9Pipelines': 300,
        }
    }

    rules = (
        Rule(LinkExtractor(allow=r'/forum-179-\d.html'), callback='parse_item', follow=True),
    )

    def parse_item(self, response):
        tbody = response.xpath('//tbody[re:test(@id,"^normalthread_.*")]')
        for movie in tbody:
            title = movie.css("div.tl_ct > a::text").extract_first()
            url = movie.css("div.tl_ct > a::attr(href)").extract_first()
            item = Sy9MovieItem(title=title, url=url)
            yield item

        request = scrapy.Request(url=url, callback=self.parse_movie_detail)
        request.meta['url'] = url
        yield request

    def parse_movie_detail(self, response):
        domain = "http://www.sy9d.com/"
        url = response.meta['url']
        imgUrl = response.css('ignore_js_op > img::attr(zoomfile)').extract()
        for index in range(len(imgUrl)):
            imgUrl[index] = domain + imgUrl[index]

        description = response.css('font::text').extract()
        delList = []
        for index in range(len(description)):
            if description[index] == '\r\n' or (not description[index]):
                delList.append(index)

        for i in delList:
            del description[i]

        detailItem = Sy9MovieItemDetail(url=url, imgUrl=imgUrl, description=description)

        yield detailItem