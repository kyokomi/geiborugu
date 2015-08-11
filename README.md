# geiborugu

geiborugu（[ゲイ・ボルグ](https://ja.wikipedia.org/wiki/%E3%82%B2%E3%82%A4%E3%83%BB%E3%83%9C%E3%83%AB%E3%82%B0)） is simple slack post message cli tool

## Install

```
go get github.com/kyokomi/geiborugu
```

## Usage

```
cat hogehoge.txt | geiborugu --token <slack_token> --channel <channel_name> --name <bot_name> --icon <bot_icon_url>
```

`--dry-run`: not post Message.

```
cat hogehoge.txt | geiborugu --token "x-xxxxxxx-xxxxx-xxxxx" --channel "#general" --name "bot" --dry-run
```

## License

[MIT](https://github.com/kyokomi/geiborugu/blob/master/LICENSE)

