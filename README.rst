=====================
`Minify CSS`_ via Go_
=====================

`Minify CSS`_ via Golang_.

The steps:

  1. Concatenate all CSS file.
  2. Remove C/C++ comments [2]_ [3]_
  3. Remove all leading and trailing white space of each line.

Development Environment:

  - `Ubuntu 15.10`_
  - `Go 1.6`_

Install
+++++++

.. code-block:: bash

  $ go get -u github.com/siongui/mincss

Usage
+++++

.. code-block:: go

  package main

  import (
  	"github.com/siongui/mincss"
  	"io/ioutil"
  )

  func main() {
  	cssPathes := []string{
  		"tipitaka/app/css/tipitaka-latn.css",
  		"tipitaka/app/css/app.css",
  	}
  	minifiedCss := mincss.MinifyCSS(cssPathes)
  	println(minifiedCss)
  	if err := ioutil.WriteFile("tipitaka/app/css/app.min.css", []byte(minifiedCss), 0644); err != nil {
  		panic(err)
  	}
  }

UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] `regex match c comments <https://www.google.com/search?q=regex+match+c+comments>`_

.. [2] `Finding Comments in Source Code Using Regular Expressions – Stephen Ostermiller <http://blog.ostermiller.org/find-comment>`_

.. [3] `Comments - cppreference.com <http://en.cppreference.com/w/cpp/comment>`_

.. [4] `regexp - The Go Programming Language <https://golang.org/pkg/regexp/>`_

.. [5] `ioutil - The Go Programming Language <https://golang.org/pkg/io/ioutil/>`_

.. [6] `bufio - The Go Programming Language <https://golang.org/pkg/bufio/>`_

.. [7] `bytes - The Go Programming Language <https://golang.org/pkg/bytes/>`_

.. [8] `strings - The Go Programming Language <https://golang.org/pkg/strings/>`_

.. _Minify CSS: https://www.google.com/search?q=Minify+CSS
.. _Go: https://golang.org/
.. _Golang: https://golang.org/
.. _Ubuntu 15.10: http://releases.ubuntu.com/15.10/
.. _Go 1.6: https://golang.org/dl/
.. _UNLICENSE: http://unlicense.org/
