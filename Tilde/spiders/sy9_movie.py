# -*- coding: utf-8 -*-
from scrapy.http import Request
from scrapy.spiders import CrawlSpider

from Tilde.items import Sy9MovieItem, Sy9MovieItemDetail


class Sy9MovieSpider(CrawlSpider):
    name = 'sy9-movie'
    allowed_domains = ['sy9d.com']

    start_urls = [
        "http://www.sy9d.com/forum-179-1.html",
        # "http://www.sy9d.com/forum.php?mod=forumdisplay&fid=179",
    ]
    # rules = [
    #     Rule(LinkExtractor(allow=r'/forum-179-\d+.html'), callback='parse_movie_list',
    #          follow=True),
    # ]

    custom_settings = {
        'ITEM_PIPELINES': {
            'demo.pipelines.Sy9Pipelines': 300,
            'demo.pipelines.Sy9ImgPipelines': 2,
        }
    }

    def parse(self, response):
        tbody = response.xpath('//tbody[re:test(@id,"^normalthread_.*")]')
        for movie in tbody:
            title = movie.css("div.tl_ct > a::text").extract_first()
            url = movie.css("div.tl_ct > a::attr(href)").extract_first()
            item = Sy9MovieItem(title=title, url=url)
            yield item

            request = Request(url=url, callback=self.parse_movie_detail)
            request.meta['url'] = url
            yield request

        # next_page_url = response.css("a.nxt::attr(href)").extract_first()
        # if next_page_url is not None:
        #     yield Request(next_page_url)

    def parse_movie_detail(self, response):
        domain = "http://www.sy9d.com/"
        url = response.meta['url']
        imgUrlTmp = response.css('ignore_js_op > img::attr(zoomfile)').extract()
        imgUrl = ""
        description = ""
        file_urls = []
        for index in range(len(imgUrlTmp)):
            if imgUrlTmp[index]:
                if imgUrl:
                    imgUrl = imgUrl + ";" + domain + imgUrlTmp[index]
                else:
                    imgUrl = domain + imgUrlTmp[index]
                file_urls.append(domain + imgUrlTmp[index])

        descriptionTmp = response.css('font::text').extract()
        for index in range(len(descriptionTmp)):
            s = self.clean_spaces(descriptionTmp[index])
            if s:
                description = description + s;

        if imgUrl:
            detailItem = Sy9MovieItemDetail(url=url, imgUrl=imgUrl, description=description, file_urls=file_urls)
            yield detailItem

    def clean_spaces(self, str):
        str = str.replace('\r\n', '')
        str = str.replace('\t', ' ')
        str = str.replace('\f', ' ')
        str = str.replace('\r', '')
        str = str.replace('\n', '')
        return str
