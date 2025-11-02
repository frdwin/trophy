# 🏆 trophy

## Description:
`trophy` is a terminal application launcher for GNU/Linux.
It's like a terminal rofi (trophy).

## Dependencies:
- A terminal, like alacritty, ghostty, etc;
- A fuzzy finder application, like fzf, sk, etc.

## Usage:
```bash
$ trophy -h
Usage of trophy:
  -f string
        The fuzzy finder application of your choice. (default "/usr/bin/sk")
  -t string
        The terminal command of your choice to open terminal apps. (default "/usr/bin/ghostty -e")

$ trophy -f "/usr/bin/fzf" -t "/usr/bin/alacritty -e"
```

## Configuration
If you prefer to use a config file, instead of command line arguments, just write a json
named config.json inside $HOME/.config/trophy, like this example:

```json
{
  "fuzzy": "/usr/bin/fzf",
  "terminal": "/usr/bin/alacritty -e"
}
```

## Instalation:
```bash
$ git clone https://github.com/frdwin/trophy && cd trophy
$ go build -o trophy cmd/*.go
$ sudo mv trophy /usr/bin
```
