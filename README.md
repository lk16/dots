## Strategy games framework written in Golang

### Run like this (or similar)
go run -tags gtk_3_18 main.go


### TODO
- [x] split cli_game in controller with different front_ends:
    - [x] cli
    - [x] gtk
    - [ ] webserver
- [ ] game state evaluation:
    - [ ] hashtable
    - [ ] parallellised search
    - [ ] opening book
- [ ] rated tournaments between players
- [ ] read PGNs:
    - [ ] kurnik
    - [ ] flyordie
- [ ] support for other games:
    - [ ] connect four
    - [ ] checkers
    - [ ] trexo
    - [ ] chess