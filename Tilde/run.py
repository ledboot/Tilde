from scrapy import cmdline

name = 'sy9-movie'
cmd = 'scrapy crawl {0}'.format(name)
cmdline.execute(cmd.split())