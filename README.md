
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
- [ ] create package and interface for tree search algorithms
- [ ] create benchmarks like https://campoy.cat/blog/justforfunc-28-benchmarks/
- [ ] unit test tree search using ffo test set
- [ ] linting
- [ ] remove most single char identifiers
- [ ] xot openings
- [ ] parallel search with hash table in separate goroutine
- [ ] better cli arguments
- [ ] better cli help strings
- [ ] game state evaluation:
- [ ] train opening book from games against bot
- [ ] parse kurnik/flyordie PGNs
- [ ] evaluate parsed PGNs
- [ ] train opening book from PGNs
- [ ] train opening book stand alone
- [ ] rated tournaments between players and bots
- [ ] support front_end webserver

#### Ideas
- create terminal UI like https://github.com/rouzwawi/reversi-go/blob/master/cmd/reversi/main.go