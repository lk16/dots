
# Dots
Othello game with AI

### Install
```go get -tags gtk_3_18 github.com/lk16/dots```

### Test
```go test -tags gtk_3_18 dots/... ```

### TODO
- [ ] move heuristic to bot_heuristic
- [ ] create package and interface for tree search algorithms
- [ ] unit test tree search using ffo test set
- [ ] rename board package to othello
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