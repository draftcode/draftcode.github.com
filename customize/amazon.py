# -*- coding: utf-8 -*-
from __future__ import division, absolute_import, print_function, unicode_literals

import bottlenose
import docutils.nodes
import docutils.parsers.rst
import docutils.parsers.rst.directives
import lxml.objectify

import amazon_secret

def asin_reference_role(role, rawtext, text, lineno, inliner,
                        options=None, content=None):
    amazon = bottlenose.Amazon(str(amazon_secret.AMAZON_ACCESS_KEY_ID),
                               str(amazon_secret.AMAZON_SECRET_KEY),
                               str(amazon_secret.AMAZON_ASSOC_TAG),
                               Region="JP")
    response = amazon.ItemLookup(ItemId=str(text))
    root = lxml.objectify.fromstring(response)
    url = root.Items.Item.DetailPageURL
    title = root.Items.Item.ItemAttributes.Title

    if not options: options = {}
    node = docutils.nodes.reference(rawtext, title, refuri=url, **options)
    return [node], []

docutils.parsers.rst.roles.register_canonical_role('asin', asin_reference_role)
docutils.parsers.rst.roles.register_canonical_role('isbn', asin_reference_role)

class AmazonDirective(docutils.parsers.rst.Directive):
    required_arguments = 1
    optional_arguments = 0
    final_argument_whitespace = True
    has_content = False

    def run(self):
        amazon = bottlenose.Amazon(str(amazon_secret.AMAZON_ACCESS_KEY_ID),
                                   str(amazon_secret.AMAZON_SECRET_KEY),
                                   str(amazon_secret.AMAZON_ASSOC_TAG),
                                   Region="JP")
        response = amazon.ItemLookup(ItemId=str(self.arguments[0]))
        root = lxml.objectify.fromstring(response)
        url = root.Items.Item.DetailPageURL
        title = root.Items.Item.ItemAttributes.Title

        node = docutils.nodes.reference(self.block_text, title, refuri=url,
                                        **self.options)
        # print(node)
        # return [node]
        return []

docutils.parsers.rst.directives.register_directive('amazon', AmazonDirective)
