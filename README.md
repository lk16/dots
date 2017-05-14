
# Dots
Strategy games framework written in Golang

### Run like this (or similar)
```go run -tags gtk_3_18 main.go```


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
    - [ ] opening book
        - [ ] learn from from games against bot
        - [ ] learn from PGNs
        - [ ] learn stand alone
- [ ] rated tournaments between players and bots
- [ ] PGNs:
    - [ ] parse
        - [ ] kurnik
        - [ ] flyordie
    - [ ] evaluate
- [ ] support several games:
    - [x] othello
    - [ ] connect four
    - [ ] checkers
    - [ ] trexo
    - [ ] chess