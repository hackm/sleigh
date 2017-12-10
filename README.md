# sleigh

`Sleigh` is simple and light weight file sync tool.  
You can sync file some os platform(Windows, Mac, Linux,,) in the Local Area Network.

[![Build Status](https://travis-ci.org/hackm/sleigh.svg?branch=master)](https://travis-ci.org/hackm/sleigh)
[![Coverage Status](https://coveralls.io/repos/github/hackm/sleigh/badge.svg?branch=master)](https://coveralls.io/github/hackm/sleigh?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/hackm/sleigh)](https://goreportcard.com/report/github.com/hackm/sleigh)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/hackm/sleigh/blob/master/LICENSE)

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

## Contribute

Please follow [Contributor's Guide](CONTRIBUTING.md)

## License

[MIT © HackM](LICENSE)
