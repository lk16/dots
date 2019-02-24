let ws;

let state = {
    "white": [27, 36],
    "black": [28, 35],
    "turn": 0
};

update_fields = function() {
    for(let i=0; i<64; i++){
        let y = Math.floor(i/8);
        let x = i%8;

        let image = "";

        if(state.white.includes(i)){
            image = "white.png";
        } else if(state.black.includes(i)){
            image = "black.png";
        } else {
            image = "empty.png";
        }

        $("#board img").eq(i).attr("src", "static/" + image);
    }
};

$(function(){
    for(let y=0; y<8;y++){
        let row = $("<tr></tr>");
        $("#board").append(row);
        for (let x=0; x<8; x++){
            $(row, "tr").append("<td></td>");
        }
    }

    $('#board td').append('<img src="static/empty.png" />');

    update_fields();

    if (ws) {
        return false;
    }
    ws = new WebSocket('ws://localhost:8080/ws');
    ws.onopen = function(evt) {
        console.log("OPEN");
    };
    ws.onclose = function(evt) {
        console.log("CLOSE");
        ws = null;
    };
    ws.onmessage = function(evt) {
        console.log("RESPONSE: " + evt.data);
        state = JSON.parse(evt.data);
        update_fields();
    };
    ws.onerror = function(evt) {
        console.log("ERROR: " + evt.data);
    };
    return false;
});

$(document).on("click", "#board td", function () {
    let y = $(this).parent().index();
    let x = $(this).index();
    let cell_id = 8*y + x;

    let ws_message = {
        'event': 'click',
        'data': {
            'cell': cell_id,
            'state': state
        }
    };

    ws.send(JSON.stringify(ws_message));
    return false;
});