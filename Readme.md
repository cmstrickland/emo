# emo - simple emoji printer

## Bring emoji to your UNIX shell

This is a simple static binary utility program you can use to generate unicode emoji glyphs easily at the unix command line using descriptive names

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

## Output modes

emo indexes all the emoji it knows about using a normalized name, based on the descriptive name for the code points used by the Unicode charts. These take the form of whitespace separated upper case English strings. emo normalizes these by transforming them to lowercase hyphen separated words.

e.g. 'U+1f600 - GRINNING FACE' becomes 'grinning-face'

pass one of these normal names to the program as an argument, and it will print the unicode glyph. In order for this to work well you need to be using a unicode locale, and be using a typeface in your terminal that has glyphs for the relevant code points. To date, this is only tested on my Mac (using "en_GB.UTF-8" and Menlo).

The program accepts command line switches in the usual style. Here are some of the more interesting ones.

The flag parsing is shonky. Typically only a single flag makes sense 

### help switches

* --list list all the known emoji and their names
* --help lists a summary of the command flags, and their default values

### mode switches, only one of these is valid

* -O output an octal escape sequence for the emoji in bytes, suitable for using as PS1
* -o output an octal escape sequence for the string 0 padded, suitable for using with echo
* -x output a hexadecimal escape sequence for the byte string
* -U output a unicode code point in U+NNNN notation
* -e interpolate backslash escapes in the argument string like echo. Prefix emoji name with \e

### output modifier switches

* -n suppress newline ( like echo )

## Examples
```bash
bash $ ./emo --list | grep grin
😸 - grinning-cat-face-with-smiling-eyes
😀 - grinning-face
😁 - grinning-face-with-smiling-eyes
bash $ ./emo grinning-face-with-smiling-eyes
😁
bash $ ./emo -O grinning-face 
\360\237\230\200
bash $ PS1="\360\237\230\200 >"
😀 >./emo -o grinning-face
\0360\0237\0230\0200
😀 >echo -e "\0360\0237\0230\0200" 
😀
😀 >./emo -U grinning-cat-face-with-smiling-eyes
U+1F638
😀 >./emo -e 'Hello, world! \egrinning-cat-face-with-smiling-eyes !!'
Hello, world! 😸 !!
😀 >
```
