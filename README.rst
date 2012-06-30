=================================
How to build draftcode.github.com
=================================
:Author: draftcode
:Date:   2012-06-30 19:46:24
:Slug:   297051a6-4fb5-4c4a-9df5-efab17b69625

clone してきます。

::

    git clone git://github.com/draftcode/draftcode.github.com.git
    git remote add dropbox ~/Dropbox/repositories/draftcode.github.com.git
    git checkout source

`Pelican`_ が必要です。

::

    git clone https://github.com/ametaireau/pelican.git
    cd pelican
    python setup.py develop

`ghp-import`_ が必要です。

::

    pip install ghp-import

``.git/hooks/post-commit`` に次のようなコードを仕込んでおきます:

::

    pelican -s conf.py && ghp-import output && git push origin gh-pages:master source:source
    git push dropbox gh-pages:gh-pages source:source

.. _`Pelican`: https://github.com/ametaireau/pelican
.. _`ghp-import`: https://github.com/davisp/ghp-import

