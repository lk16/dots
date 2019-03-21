
# Dots
Othello game with AI

### Install
```
go get github.com/lk16/dots
go install github.com/lk16/dots/cmd/dots
```

### Run
```
dots
```

### Test
```
go test ./...
```

### Linter

```
golangci-lint run
```

### TODO
- [ ] web front end
    - [ ] use css ids instead of class where appropriate
    - [ ] fix bug: last few moves don't show analysis in browser
    - [ ] fix bug: web client should ignore received bot_move_reply after undo 
    - [ ] reconnect websocket periodically if server goes down
    - [ ] redo index.html with templates


- [ ] cleaning
    - [ ] return (*Board, error) from othello.RandomBoard()
    - [ ] create BotHeuristic.write() that does error checking
    - [ ] start using https://github.com/pkg/errors with .Cause()
    - [ ] move GameState from frontend to othello package
    - [ ] move web.newState() and web.getBoard() to othello package
    - [ ] go linting
    - [ ] clean board tests
    - [ ] javascript code is a mess

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
