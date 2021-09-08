Hat
===

Hat is cat for human.

Use
---

    $ hat go.mod LICENSE
    ------------------------------------------------------------------------
    FILE: go.mod
    ------------------------------------------------------------------------
    module github.com/bbriano/hat

    go 1.16

    require golang.org/x/term v0.0.0-20210615171337-6886f2dfbf5b // indirect
    ------------------------------------------------------------------------
    FILE: LICENSE
    ------------------------------------------------------------------------
    Copyright (c) 2021 Briano Goestiawan

    Permission to use, copy, modify, and/or distribute this
    software for any purpose with or without fee is hereby
    granted, provided that the above copyright notice and this
    permission notice appear in all copies.

    THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS
    ALL WARRANTIES WITH REGARD TO THIS SOFTWARE INCLUDING ALL
    IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS. IN NO
    EVENT SHALL THE AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT,
    INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES WHATSOEVER
    RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
    ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION,
    ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE
    OF THIS SOFTWARE.

About
-----

Based on Classic Hat (shell script): https://gist.github.com/bbriano/0c5d4f9dc0f06e22fd066e3abf243ba1

Benchmark
---------

Tested on 201 files with a sum of 286,278 characters (all times are in ms).

| Program     | #1   | #2   | #3   | Avg  |
|-------------|------|------|------|------|
| Classic Hat | 3816 | 4087 | 3892 | 3932 |
| Hat         | 110  | 115  | 120  | 115  |

34x speed improvement over Classic Hat.
