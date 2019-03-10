
# Dots
Othello game with AI

### Install
```
go get -tags gtk_3_18 github.com/lk16/dots
```

### Test
```
go test -tags gtk_3_18 dots/...
```

### Linter

```
gometalinter --enable-all --disable=goimports --disable=gofmt dots/...
```

### TODO
- [ ] web front end
    - [ ] add undo option
    - [ ] use css ids instead of class where appropriate
    - [ ] make every wsMessage have two fields: event and data
    - [ ] generalize ws message handling
    - [ ] fix bug: last few moves don't show analysis in browser
    - [ ] reconnect websocket periodically if server goes down


- [ ] cleaning
    - [ ] move web.newState() and web.getBoard() to othello package
    - [ ] remove gtk front-end
    - [ ] use import github.com/lk16/dots everywhere
    - [ ] use project lay-out like https://github.com/golang-standards/project-layout
    - [ ] start using https://github.com/pkg/errors with .Cause()
    - [ ] linting

- [ ] bot
    - [ ] allow faster killing of analysis go-routines
    - [ ] fix pending bugs from treesearch
        - [ ] unit test tree search using ffo test set
    - [ ] optimize analysis algorithm
        - [ ] parallel search with hash table in separate goroutine
        - [ ] create benchmarks like https://campoy.cat/blog/justforfunc-28-benchmarks/
    - [ ] move treesearch package back into BotHeuristic
    - [ ] openings book
        - [ ] set up db with models + config file with https://github.com/jinzhu/gorm
        - [ ] train/use opening book from games against bot
        - [ ] train opening book stand alone
        - [ ] PGN
            - [ ] parse kurnik/flyordie PGNs
            - [ ] evaluate parsed PGNs
            - [ ] train opening book from PGNs

#### Ideas
- create terminal UI like https://github.com/rouzwawi/reversi-go/blob/master/cmd/reversi/main.go