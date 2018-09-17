# -*- coding: utf-8 -*-

# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://doc.scrapy.org/en/latest/topics/item-pipeline.html
import codecs
import json

from scrapy.exceptions import DropItem

from Tilde.items import Sy9MovieItem


class TildePipeline(object):
    def process_item(self, item, spider):
        return item

class Sy9Pipelines(object):
    def __init__(self):
        self.item_file = codecs.open("sy9-item.json","w","utf-8-sig")
        self.detail_file = codecs.open("sy9-detail.json","w","utf-8-sig")

    def process_item(self,item,spider):
        if isinstance(item,Sy9MovieItem):
            self.process_sy9_item(item,spider)
        else:
            self.process_sy9_detail(item,spider)


    def process_sy9_item(self,item,spider):
        if item['title']:
            line = json.dumps(dict(item),ensure_ascii=False)+"\n"
            self.file.write(line)
            return item
        else:
            raise DropItem("Miss title in %s" % item)

    def process_sy9_detail(self,item,spider):
        if item['url']:
            line = json.dumps(dict(item),ensure_ascii=False)+"\n"
            self.file.write(line)
            return item
        else:
            raise DropItem("Miss title in %s" % item)