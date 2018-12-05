# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://doc.scrapy.org/en/latest/topics/item-pipeline.html
import codecs
import json
from os.path import basename, dirname, join
from urllib.parse import urlparse

from scrapy import Request
from scrapy.exceptions import DropItem
from scrapy.pipelines.images import ImagesPipeline

from Tilde.items import Sy9MovieItem, Sy9MovieItemDetail


class TildePipeline(object):
    def process_item(self, item, spider):
        return item


class Sy9Pipelines(object):
    def __init__(self):
        self.item_file = codecs.open("sy9-item.json", "w", "utf-8-sig")
        self.detail_file = codecs.open("sy9-detail.json", "w", "utf-8-sig")

    def process_item(self, item, spider):
        if isinstance(item, Sy9MovieItem):
            self.process_sy9_item(item, spider)
        else:
            self.process_sy9_detail(item, spider)

    def process_sy9_item(self, item, spider):
        if item['title']:
            line = json.dumps(dict(item), ensure_ascii=False) + "\n"
            self.file.write(line)
            return item
        else:
            raise DropItem("Miss title in %s" % item)

    def process_sy9_detail(self, item, spider):
        if item['url']:
            line = json.dumps(dict(item), ensure_ascii=False) + "\n"
            self.file.write(line)
            return item
        else:
            raise DropItem("Miss title in %s" % item)


class Sy9ImgPipelines(ImagesPipeline):

    def get_media_requests(self, item, info):
        for image_url in item['file_urls']:
            yield Request(image_url)

    def file_path(self, request, response=None, info=None):
        path = urlparse(request.url).path
        return join(basename(dirname(path)), basename(path))

    def process_item(self, item, spider):
        if isinstance(item, Sy9MovieItemDetail):
            self.get_media_requests(item, [])

    def item_completed(self, results, item, info):
        image_paths = [x['path'] for ok, x in results if ok]
        if not image_paths:
            raise DropItem("Item contains no images")
        item['image_paths'] = image_paths
        return item
