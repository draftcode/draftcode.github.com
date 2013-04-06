# -*- coding: utf-8 -*-
from __future__ import division, absolute_import, print_function, unicode_literals

import customize.amazon
import customize.japanese_text_join

AUTHOR = 'draftcode'
DEFAULT_CATEGORY = "Articles"
DELETE_OUTPUT_DIRECTORY = True
PATH = '.'
PAGE_DIR = 'pages'
ARTICLE_DIR = 'articles'
SITENAME = 'draftcode.github.io'
SITEURL = 'http://draftcode.github.io'
STATIC_PATHS = ["images"]
TIMEZONE = 'Asia/Tokyo'
DEFAULT_PAGINATION = 3

ARTICLE_URL = '{date:%Y}/{date:%m}/{date:%d}/{slug}.html'
ARTICLE_SAVE_AS = ARTICLE_URL

FEED_ATOM = "atom.xml"
FEED_DOMAIN = "http://draftcode.github.io/"

DEFAULT_LANG = 'ja'

THEME = 'theme/'

DISQUS_SITENAME = 'draftcode-github-com'
GITHUB_URL = 'http://github.com/draftcode/'
SOCIAL = (('Twitter', 'http://twitter.com/#!/draftcode'),
          ('GitHub', 'http://github.com/draftcode'),
          ('Google+', 'https://plus.google.com/107177890582465029754?rel=author'))
TWITTER_USERNAME = 'draftcode'
