#!/usr/bin/python2.7
# vim: fileencoding=utf-8
from __future__ import unicode_literals

AUTHOR = 'draftcode'
DEFAULT_CATEGORY = "Articles"
DELETE_OUTPUT_DIRECTORY = True
PATH = '.'
PAGE_DIR = 'pages'
ARTICLE_DIR = 'articles'
SITENAME = 'draftcode.github.com'
SITEURL = 'http://draftcode.github.com'
STATIC_PATHS = ["images"]
TIMEZONE = 'Asia/Tokyo'
DEFAULT_PAGINATION = 3

ARTICLE_URL = '{date:%Y}/{date:%m}/{date:%d}/{slug}.html'
ARTICLE_SAVE_AS = ARTICLE_URL

FEED = "atom.xml"
FEED_DOMAIN = "http://draftcode.github.com/"

DEFAULT_LANG = 'ja'

THEME = 'theme/'

DISQUS_SITENAME = 'draftcode-github-com'
GITHUB_URL = 'http://github.com/draftcode/'
SOCIAL = (('Twitter', 'http://twitter.com/#!/draftcode'),
          ('GitHub', 'http://github.com/draftcode'),)
TWITTER_USERNAME = 'draftcode'

##############################################################################

import unicodedata
import docutils.nodes
import docutils.transforms

class JapaneseTextJoin(docutils.transforms.Transform):
    default_priority = 800

    def apply(self):
        for text in self.document.traverse(docutils.nodes.Text):
            if (not isinstance(text.parent, docutils.nodes.literal_block) and
                not isinstance(text.parent, docutils.nodes.raw)):
                lines = []
                prev_category = ''
                for line in text.astext().splitlines():
                    line = line.strip()
                    if len(lines) > 0 and len(line) > 0:
                        if prev_category == 'Lo' or unicodedata.category(line[0]) == 'Lo':
                            lines[-1] += line
                        else:
                            lines.append(line)
                        prev_category = unicodedata.category(line[-1])
                    else:
                        lines.append(line)
                joined_text = '\n'.join(lines)
                text.parent.replace(text, docutils.nodes.Text(joined_text, text.rawsource))

import docutils.parsers.rst
origparser = docutils.parsers.rst.Parser
class Parser(origparser):
    def get_transforms(self):
        return origparser.get_transforms(self) + [JapaneseTextJoin]
docutils.parsers.rst.Parser = Parser

