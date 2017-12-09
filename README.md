# sleigh
TBD
# sleigh

## About

`Sleigh` is simple and light weight file sync tool. 
You can sync file some os platform(Windows, Mac, Linux,,) in closed internet.

## Run

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

## Options

| name | short | content |
|:----:|:----:|:-------:|
| --room | -r | room name ex, "hackm". **required** |
| --password | -p | room password. |
| --listen | -l | listen port number. default:  |

## License