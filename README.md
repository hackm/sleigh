# sleigh

`sleigh` is simple, quick, secure and lightweight file sync tool on the Local Area Network. On any platform (Windows, macOS, Linux), it's make you comfortable file sync.

[![Build Status](https://travis-ci.org/hackm/sleigh.svg?branch=master)](https://travis-ci.org/hackm/sleigh)
[![Coverage Status](https://coveralls.io/repos/github/hackm/sleigh/badge.svg?branch=master)](https://coveralls.io/github/hackm/sleigh?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/hackm/sleigh)](https://goreportcard.com/report/github.com/hackm/sleigh)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/hackm/sleigh/blob/master/LICENSE)

```
                    _...,
              o_.-"`    `\
       .--.  _ `'-._.-'""-;     _
     .'    \`_\_  {_.-'""-}  _ / \              888        d8b        888
   _/     .-'  '. {c-._o_.){\|`  |              888        Y8P        888
  (@`-._ /       \{    ^  } \\ _/               888                   888
    `~\  '-._      /'.     }  \}  .-.   .d8888b 888 .d88b. 888 .d88b. 88888b.
     |>:<   '-.__/   '._,} \_/  / ())   88K     888d8P  Y8b888d88P"88b888 "88b
     |     >:<   `'---. ____'-.|(`"`    "Y8888b.88888888888888888  888888  888
     \            >:<  \\_\\_\ | ;           X88888Y8b.    888Y88b 888888  888
      \                 \\-{}-\/  \      88888P'888 "Y8888 888 "Y88888888  888
       \                 '._\\'   /)                               888
        '.                       /(                           Y8b d88P
          `-._ _____ _ _____ __.'\ \
            / \     / \     / \   \ \
         _.'/^\'._.'/^\'._.'/^\'.__) \
     ,=='  `---`   '---'   '---'      )
     `"""""""""""""""""""""""""""""""`
```

## Demo

![demo](https://github.com/hackm/sleigh/blob/master/images/sleigh.gif)

## Features

- Not require configuration
- Easy Install
- Real-Time synchronization
- Quick synchronization
- Only limited on the LAN
- Windows, macOS and Linux support
- Exciting synchronization views
- HTTPS and authentication

## Installation (WIP)

Get the latest [release](https://github.com/hackm/sleigh/releases).

On macOS, sleigh can be installed via [Homebrew](https://brew.sh/)
```
$ brew install sleigh
```

## Build

This project uses the [dep](https://github.com/golang/dep). Go check it out if you don't have them locally installed.  
If you can use go build environment, getting the latest version by yourself own build.

```
$ go get github.com/hackm/sleigh
$ cd $GOPATH/src/github.com/hackm/sleigh
$ go build && go install
$ sleigh
```

## Usage

```
$ cd /path/to/share_folder
$ sleigh
```

## Options

| name | short | content |
|:----:|:----:|:-------:|
| --room | -r | room name e.g. "hackm". |
| --listen | -l | listening UDP port. default: 8986 |

## Motivation

- We want good tool for share files easily and fast in poor network.
- It's may be beneficial for crowded place like hackathon venue.
- We often face to take time to sync on Dropbox in their places.
- AirDrop? I'm Windows User...
- Resilio Sync? It's huge...
- For us &gt; For someone

## TODO

- [x] Multicast device connection
- [x] File create sync
- [x] File modified sync
- [ ] File deleted sync
- [ ] File rename sync
- [ ] Change using port
- [ ] Set room name and password

## Contribute

Please follow [Contributor's Guide](CONTRIBUTING.md)

## License

[MIT © HackM](LICENSE)
