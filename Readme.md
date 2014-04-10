# emo - simple emoji printer

## Bring emoji to your UNIX shell

This is a tiny binary you can use to generate unicode emoji glyphs easily at the unix command line using descriptive names

```bash
$ ./emo --list | grep horse
🎠  - carousel-horse
🐎  - horse
🐴  - horse-face
🏇  - horse-racing
$ HORSEFACE=$(./emo -n horse-face)
$ echo $HORSEFACE
🐴
$ 
```
## Building

Requires golang 1.2. Should just build with 'go build'

## Using

The code should be portable, however you need a shell interpreter that has a decent unicode locale, and is using a font that has nice glyphs for the Unicode-6 emoji code points. Works well OOTB on my Mac.
