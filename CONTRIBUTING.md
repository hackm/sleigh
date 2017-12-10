# Contribution Guide

Welcomes your contribution. To make the process as seamless as possible, we recommend you read this contribution guide.

## Development Workflow

Start by forking this repository, make changes in a branch and then send a pull request.

### Setup Repository
Fork https://github.com/hackm/sleigh/fork source repository to your own personal repository.

```sh
$ mkdir -p $GOPATH/src/github.com/sleigh
$ cd $GOPATH/src/github.com/sleigh
$ git clone <paste saved URL for personal forked repo>
$ cd sleigh
```

### Set up git remote as ``upstream``
```sh
$ cd $GOPATH/src/ghttps://github.com/hackm/sleigh
$ git remote add upstream https://github.com/hackm/sleigh
$ git fetch upstream
$ git merge upstream/master
...
```

### Create a feature branch
Before changing the code, make sure you create a separate branch for these changes

```
$ git checkout -b my-new-feature
```

### Test
After your code changes, make sure

- To run `make`
- To squash your commits into a single commit. `git rebase -i`. It's okay to force update your pull request.
- To run `go test -race ./...` and `go build` completes.

### Commit changes
After verification, commit your changes.

```
$ git commit -am 'Add some feature'
```

### Push to the branch
Push your locally committed changes to the remote origin
```
$ git push origin new-feature
```

### Create a Pull Request
Pull requests can be created via GitHub. Refer to [this document](https://help.github.com/articles/creating-a-pull-request/) for detailed steps on how to create a pull request. After a Pull Request gets peer reviewed and approved, it will be merged.
