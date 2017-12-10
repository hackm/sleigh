# sleigh

`Sleigh` is simple and light weight file sync tool.  
You can sync file some os platform(Windows, Mac, Linux,,) in the Local Area Network.

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

## Motivation

- We want good tool for share files easily and fast in poor network. 
- It's may be beneficial for crowded place like hackathon venue.
- We often face to take time to sync on Dropbox in there places.
- AirDrop? I'm Windows User...
- Resilio Sync? It's huge...
- For us > For someone 

## TODO

- [x] Multicast device connection
- [x] File create sync
- [x] File modified sync
- [x] File deleted sync
- [ ] File rename sync
- [ ] Change using port
- [ ] Set room name and password

## Contribute

Please follow [Contributor's Guide](CONTRIBUTING.md)

## License

[MIT © HackM](LICENSE)
