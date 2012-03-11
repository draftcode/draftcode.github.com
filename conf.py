#!/usr/bin/python2.7
# vim: fileencoding=utf-8
from __future__ import unicode_literals

AUTHOR = 'draftcode'
SITENAME = 'draftcode.github.com'
SITEURL = 'http://draftcode.github.com'
AUTHOR_URL = SITEURL
PATH = 'entry/'
THEME = 'theme/'

ARTICLE_URL = '{date:%Y}/{date:%m}/{date:%d}/{slug}.html'
ARTICLE_SAVE_AS = '{date:%Y}/{date:%m}/{date:%d}/{slug}.html'
GITHUB_URL = 'http://github.com/draftcode/'
DISQUS_SITENAME = 'draftcode-github-com'
PDF_GENERATOR = False
WITH_PAGINATION = True
DEFAULT_LANG = 'ja'
TIMEZONE = 'Asia/Tokyo'

FEED = "atom.xml"
# FEED_RSS = None
# CATEGORY_FEED = None
# TRANSLATION_FEED = None

SOCIAL = (('twitter', 'http://twitter.com/#!/draftcode'),
          ('github', 'http://github.com/draftcode'),)
TWITTER_USERNAME = 'draftcode'

# static paths will be copied under the same name
STATIC_PATHS = ["img",]

# A list of files to copy from the source to the destination
FILES_TO_COPY = (('extra/robots.txt', 'robots.txt'),)

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

