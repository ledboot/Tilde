import functools
from scrapy.exceptions import DropItem


def check_spider_pipeline(process_item_method):
    @functools.wraps(process_item_method)
    def wrapper(self, item, spider):
        msg = '%%s %s pipeline step' % (self.__class__.__name__,)
        if self.__class__ in spider.pipeline:
            spider.logger.debug(msg % 'executing')
            return process_item_method(self, item, spider)
        else:
            spider.logger.debug(msg % 'skipping')
            raise DropItem("Missing pipeline property")

    return wrapper
