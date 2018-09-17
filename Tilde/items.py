# -*- coding: utf-8 -*-

# Define here the models for your scraped items
#
# See documentation in:
# https://doc.scrapy.org/en/latest/topics/items.html

import scrapy


class TildeItem(scrapy.Item):
    # define the fields for your item here like:
    # name = scrapy.Field()
    pass


class Sy9MovieItem(scrapy.Item):
    title = scrapy.Field()
    url = scrapy.Field()


class Sy9MovieItemDetail(scrapy.Item):
    url = scrapy.Field()
    imgUrl = scrapy.Field()
    description = scrapy.Field()