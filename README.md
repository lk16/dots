
# Dots
Othello game with AI

### Install
```go get -tags gtk_3_18 github.com/lk16/dots```

### TODO
- [ ] support xot openings
- [ ] support several front_ends:
    - [x] command line
    - [x] gtk
    - [ ] webserver
- [ ] game state evaluation:
    - [x] mtdf
    - [ ] hashtable
    - [ ] parallel search
    - [ ] learning opening book
        - [ ] from from games against bot
        - [ ] from PGNs
        - [ ] stand alone
- [ ] rated tournaments between players and bots
- [ ] PGNs:
    - [ ] parse
        - [ ] kurnik
        - [ ] flyordie
    - [ ] evaluate
