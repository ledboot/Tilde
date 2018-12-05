from scrapy import cmdline

name = 'mebookList'
cmd = 'scrapy crawl {0}'.format(name)
cmdline.execute(cmd.split())