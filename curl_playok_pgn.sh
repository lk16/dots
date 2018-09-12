#!/bin/bash
curl -s "https://www.playok.com/en/game.phtml?gid=rv&pid=$1&txt" > $1.txt
