# scrapbox

Utils that don't deserve separate repo

## httpwd

Create password credentials for htaccess and htpasswd files. Only bcrypt.

## md2html

Convert markdown from stdint to html.Uses [goldmark](https://github.com/yuin/goldmark)

## psub

Standalone implementation of Fish shell [psub](https://fishshell.com/docs/current/cmds/psub.html)

```sh
diff $(sort file1 | psub) $(sort file2 | psub)
# more or less equals
diff <(sort file1) <(sort file2)
```

## Walk

Walk directory tree outputting paths to stdout.

## nanoid

Generate nanoid

```sh
nanoid -c 16 # 0WDx0Y8QReDJpFzo
```

## tg

Like [jo](https://github.com/jpmens/jo) but for HTML

```sh
$ tg p This is paragraph of $(tg -a class=foo b text)
<p>
This is paragraph of
<b class="foo" >
text
<b/>
<p/>
```
