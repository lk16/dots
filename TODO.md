### TODO
- [ ] docker
    - [ ] use environment variables to configure frontend end bot

- [ ] web front end
    - [ ] use css ids instead of class where appropriate
    - [ ] fix bug: last few moves don't show analysis in browser
    - [ ] fix bug: web client should ignore received bot_move_reply after undo
    - [ ] reconnect websocket periodically if server goes down
    - [ ] redo index.html with templates
    - [ ] move web folder into best-practices location
    - [ ] javascript code is a mess


- [ ] playok bot
    - [ ] clean and merge

- [ ] cleaning
    - [ ] clean board tests:
        - [ ] redo genTestBoards()
    - [ ] move bitset out of othello package or drop completely?
    - [ ] cleaner error handling
        - [ ] use `github.com/pkg/errors` everywhere
        - [ ] get rid of `fmt.Errorf()`
    - [ ] list features of this project on top of README

- [ ] bot
    - [ ] introduce SearchWinner()
    - [ ] allow faster killing of analysis go-routines
    - [ ] unit test tree search using ffo test set
    - [ ] optimize analysis algorithm
        - [ ] parallel search with hash table in separate goroutine
        - [ ] create benchmarks like https://campoy.cat/blog/justforfunc-28-benchmarks/
    - [ ] openings book
        - [ ] set up db with models + config file with https://github.com/jinzhu/gorm
        - [ ] train/use opening book from games against bot
        - [ ] train opening book stand alone
        - [ ] PGN
            - [ ] parse kurnik/flyordie PGNs
            - [ ] evaluate parsed PGNs
            - [ ] train opening book from PGNs
