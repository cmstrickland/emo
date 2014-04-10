# emo - simple emoji printer

## Bring emoji to your UNIX shell

This is a tiny binary you can use to generate unicode emoji glyps easily at the unix command line using descriptive names

```bash
😃 $ ./emo --list | grep horse
🎠  - carousel-horse
🐎  - horse
🐴  - horse-face
🏇  - horse-racing
😃 $ HORSEFACE=$(./emo -n horse-face)
😃 $ echo $HORSEFACE 
🐴
😃 $ 
```
