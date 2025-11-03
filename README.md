# 🏆 trophy

## Description:
`trophy` is a terminal application launcher for GNU/Linux.
It's like a terminal rofi (trophy).

![trophy screenshot](https://private-user-images.githubusercontent.com/84289565/509023769-65950271-07d8-4959-bb11-c1f1dc2406d6.png?jwt=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3NjIxODQwMTYsIm5iZiI6MTc2MjE4MzcxNiwicGF0aCI6Ii84NDI4OTU2NS81MDkwMjM3NjktNjU5NTAyNzEtMDdkOC00OTU5LWJiMTEtYzFmMWRjMjQwNmQ2LnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNTExMDMlMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjUxMTAzVDE1MjgzNlomWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPWRmMDJkYjBlYWM1MjFjZGUyNDM4ZWViY2RkNTJlMmRjMDlkYmQ5MDgyNzgwZTQ2ZGU4ZTNjN2Q4ZWU0NDAwYTMmWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0In0.LvY53aWoS-LNiF2BPlb3BQejtcvVXLYdUgxPq2wfWkU)
![trophy screenshot](https://private-user-images.githubusercontent.com/84289565/509023993-d21ed168-531c-472f-a6f2-b1348808baec.png?jwt=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3NjIxODQwMTYsIm5iZiI6MTc2MjE4MzcxNiwicGF0aCI6Ii84NDI4OTU2NS81MDkwMjM5OTMtZDIxZWQxNjgtNTMxYy00NzJmLWE2ZjItYjEzNDg4MDhiYWVjLnBuZz9YLUFtei1BbGdvcml0aG09QVdTNC1ITUFDLVNIQTI1NiZYLUFtei1DcmVkZW50aWFsPUFLSUFWQ09EWUxTQTUzUFFLNFpBJTJGMjAyNTExMDMlMkZ1cy1lYXN0LTElMkZzMyUyRmF3czRfcmVxdWVzdCZYLUFtei1EYXRlPTIwMjUxMTAzVDE1MjgzNlomWC1BbXotRXhwaXJlcz0zMDAmWC1BbXotU2lnbmF0dXJlPWZlMmQ1OWEyN2QwMDExOTkzN2I0NzA5YzlhOTgyYTU4OTIxZWIzYTc3NTU3MzJkMDEyNWY3MDg3MDg0OTY1YjImWC1BbXotU2lnbmVkSGVhZGVycz1ob3N0In0.Onrl6zoUvsM_bUWYrLB8KQqs0kNgZCCGJW5x3YVpDe4)

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
