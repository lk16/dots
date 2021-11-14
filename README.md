
# Dots
Othello game with AI

[Live demo](https://heylu.uk/at/dots/)

## Features
* Nice frontend
* Strong AI
* [Xot openings](http://berg.earthlingz.de/xot/aboutxot.php?lang=en)
* Best move analysis
* Undo option


![alt text](assets/screenshot.png "dots screenshot")

### Requirements:
* Docker
* Docker-compose

### How to use:
* Run `docker-compose up`
* Go to [localhost:8080](http://localhost:8080) in your browser

### How to run tests:
* Run `./test.sh`

---

# Dots Eval
Command line tool to evaluate one position. Outputs JSON response

Not available as endpoint because:
- Denial of Service risk
- This is a quick hack

# How to use:
```sh
./dots-eval -me 0x301c080400 -opp 0x80a02020000 -depth 7 -exact 14
```

### Upcoming features:
* See [TODO](TODO.md) file
