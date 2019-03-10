
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
    - [ ] fix analyzer bugs
        - [ ] prevent outdated analyze_move_reply message
        - [ ] make analyze_move imply analyze_stop for running analysis
    - [ ] add xot
    - [ ] add undo option
    - [ ] show heuristic for white player in white

- [ ] cleaning
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