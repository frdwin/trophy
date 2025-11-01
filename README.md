# trophy

## Description:
`trophy` is a terminal application launcher for GNU/Linux.
It's like a terminal rofi (trophy).

## Dependencies:
- A terminal;
- A fuzzy finder application, like fzf.

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

## Next steps:
- Fix bug when opening steam games shortcuts
- Add a configuration file
