let ws;

let state = {
    "white": [27, 36],
    "black": [28, 35],
    "turn": 0
};

update_fields = function() {

    let valid_moves = get_valid_moves(state);

    for(let i=0; i<64; i++){
        let image = "";

        if(state.white.includes(i)){
            image = "white.svg";
        } else if(state.black.includes(i)){
            image = "black.svg";
        } else if(valid_moves.includes(i)){
            if(state.turn === 0){
                image = "move_black.svg";
            }  else {
                image = "move_white.svg";
            }
        } else {
            image = "empty.svg";
        }

        $("#board img").eq(i).attr("src", "static/" + image);
    }
};

get_flippable_discs = function(state, move) {
    let me = state.white;
    let opp = state.black;

    if(state.turn === 0){
        me = state.black;
        opp = state.white;
    }

    if(me.includes(move) || opp.includes(move)){
        return [];
    }

    let move_x = move % 8;
    let move_y = Math.floor(move / 8);


    let flippable = [];

    for(let dy=-1; dy<=1; dy++){
        for(let dx=-1;dx<=1;dx++){
            if(dx===0 && dy===0){
                continue;
            }

            let dir_flippable = [];

            let d = 1;
            let p, px, py;

            while(true){
                px = move_x + (d * dx);
                py = move_y + (d * dy);
                p = (8 * py) + px;

                if(px<0 || px>=8 || py<0 || py>=8){
                    break;
                }

                if(!opp.includes(p)){
                    break;
                }

                dir_flippable.push(p);
                d++;
            }

            if(d===1) {
                continue;
            }

            px = move_x + (d * dx);
            py = move_y + (d * dy);
            p = (8 * py) + px;

            if(px<0 || px>=8 || py<0 || py>=8){
                continue;
            }

            if(me.includes(p)){
                flippable.push(...dir_flippable);
            }
        }
    }

    return flippable;
};

get_valid_moves = function(state){
    let valid_moves = [];
    for(let move=0; move<64; move++){
        let flippable = get_flippable_discs(state, move);
        if(flippable.length > 0){
            valid_moves.push(move);
        }
    }
    return valid_moves;
};

request_bot_move = function(){
    let message = {
        'event': 'bot_move',
        'bot_move': {
            'state': state
        }
    };

    ws.send(JSON.stringify(message))
};

$(function(){
    for(let y=0; y<8;y++){
        let row = $("<tr></tr>");
        $("#board").append(row);
        for (let x=0; x<8; x++){
            $(row, "tr").append("<td></td>");
        }
    }

    $('#board td').append('<img src="static/empty.svg" />');

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
        let message = JSON.parse(evt.data);
        switch(message.event){
            case "bot_move_reply":
                state = message.bot_move_reply.state;
                update_fields();
                if(get_valid_moves(state).length === 0){
                    state.turn = 1-state.turn;
                    if(get_valid_moves(state).length !== 0){
                        setTimeout(request_bot_move(), 250);
                    }
                }
        }
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

    if(state.turn === 0 && $("select[name='black_player']").find(":selected").val() !== 'human'){
        return false;
    }

    if(state.turn === 1 && $("select[name='white_player']").find(":selected").val() !== 'human'){
        return false;
    }

    let flipped = get_flippable_discs(state, cell_id);

    if(flipped.length === 0){
        return false;
    }

    if(state.turn === 0){
        state.black.push(cell_id, ...flipped);
        state.white = state.white.filter(x => !flipped.includes(x));
    } else {
        state.white.push(cell_id, ...flipped);
        state.black = state.black.filter(x => !flipped.includes(x));
    }
    state.turn = 1-state.turn;

    if(get_valid_moves(state).length === 0){
        state.turn = 1-state.turn;
    }

    update_fields();

    if((state.turn === 0 && $("select[name='black_player']").find(":selected").val() !== 'human') ||
        (state.turn === 1 && $("select[name='white_player']").find(":selected").val() !== 'human')){

        request_bot_move();
    }

    return false;
});

$(document).on("change", "select", function() {
    let selected = $(this).val();
    let name = $(this).attr('name');

    if(selected === "human") {
        return false;
    }
    console.log(name, state.turn);

    if((name === 'black_player' && state.turn === 0) || (name === "white_player" && state.turn === 1)){
        request_bot_move();
    }
});