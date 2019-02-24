let ws;

$(function(){
    for(let y=0; y<8;y++){
        let row = $("<tr></tr>");
        $("#board").append(row);
        for (let x=0; x<8; x++){
            $(row, "tr").append("<td></td>");
        }
    }

    $('#board td').append('<img src="static/empty.png" />');

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
    console.log(cell_id);
    ws.send(JSON.stringify({'event': 'click', 'data': {'cell': cell_id}}));
    return false;
});