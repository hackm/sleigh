# Sleigh

![build](https://travis-ci.org/hackm/sleigh.svg?branch=master) ![license](https://camo.githubusercontent.com/9e700fcd5dd47fa817872997918e8f741b9c4403/687474703a2f2f622e7265706c2e63612f76312f4c6963656e73652d4d49542d7265642e706e67)

`Sleigh` is simple and light weight file sync tool.
You can sync file some os platform(Windows, Mac, Linux,,) in closed internet.

## Install

How to run command.

### Download build file

Getting latest release build file.
[release page](https://github.com/hackm/sleigh/releases)

### App installer

Install by brew

```
brew install sleigh
```

### Build from go get

If you can use go-lang build enviroment, getting newest version by self build.

```
go get github.com/hackm/sleigh
cd $GOPATH/src/github.com/hackm/sleigh
go build
./sleigh
```

## Usage

```
$cd /path/to/share_folder
$sleigh -r hackm -p pass
```

## Options

| name | short | content |
|:----:|:----:|:-------:|
| --room | -r | room name ex, "hackm". **required** |
| --password | -p | room password. |
| --listen | -l | listen port number. default:  |

## License

MIT