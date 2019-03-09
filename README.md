
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
    - [ ] proper SVG generating endpoint
    - [ ] fix analyzer bugs
    - [ ] add xot
    - [ ] add undo option

- [ ] cleaning
    - [ ] use project lay-out like https://github.com/golang-standards/project-layout
    - [ ] linting

- [ ] bot
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