# md2dsas

Github flavored Markdown に近い Markdown から、 DSAS Blog 用の HTML を生成する

* fenced code block から、 `<pre class="prog prettify lang-python">` のような pre タグを作る
* h1 はタイトル用にテキストのまま、 h2 移行はヘッダレベルを +1 する (h2 が h3 になる)
