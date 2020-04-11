
# Playok.com login with curl

This file contains `curl` commands how to login to [playok.com](https://playok.com/) and connect to its websockets.

This is here for informative purposes and shows no new information of the internals of this site than was already public at the time of writing.

```sh
PLAYOK_USERNAME=''
PLAYOK_PASSWORD=''

USER_AGENT=''

# login
# this gets the 'ku' cookie
curl -i 'https://www.playok.com/en/login.phtml' \
-H "User-Agent: $USER_AGENT" \
--data "cc=0&username=$PLAYOK_USERNAME&pw=$PLAYOK_PASSWORD"

COOKIE_KU=''
COOKIE_KSESSION=''

# go to reversi page
# this gets the 'kt' cookie
curl -I 'https://www.playok.com/en/reversi/' \
-H "User-Agent: $USER_AGENT" \
-H "Cookie: ku=$COOKIE_KU; ksession=$COOKIE_KSESSION; kbeta=rv"

COOKIE_KT=''

# generate WS_KEY randomly
WS_KEY=$(cat /dev/urandom | tr -dc a-z | head -c 16 | base64)
echo $WS_KEY


# connect to websocket
# this uses the 'kt' and 'ku' cookies
curl -i 'https://x.playok.com:17003/ws/' \
-H "User-Agent: $USER_AGENT" \
-H 'Pragma: no-cache' \
-H 'Origin: https://www.playok.com' \
-H "Sec-WebSocket-Key: $WS_KEY" \
-H 'Upgrade: websocket' \
-H "Cookie: ku=$COOKIE_KU; kt=$COOKIE_KT;" \
-H 'Connection: Upgrade' \
-H 'Sec-WebSocket-Version: 13'
```