function update_fields(board) {

    let valid_moves = get_valid_moves(board);

    for(let i=0; i<64; i++){
        let image = '';

        if(board.white.includes(i)){
            image = './svg/field?disc=white';
        } else if(board.black.includes(i)){
            image = './svg/field?disc=black';
        } else if(valid_moves.includes(i)){
            if(board.turn === 0){
                image = './svg/field?move=black';
            }  else {
                image = './svg/field?move=white';
            }
        } else {
            image = './svg/field';
        }

        $('#board img').eq(i).attr('src', image);
    }

    $('.white_disc_count').attr('src', './svg/field?disc=white&text=' + board.white.length);
    $('.black_disc_count').attr('src', './svg/field?disc=black&text=' + board.black.length);
}

function get_flippable_discs(board, move) {
    let me = board.white;
    let opp = board.black;

    if(board.turn === 0){
        me = board.black;
        opp = board.white;
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
}

function get_valid_moves(board){
    let valid_moves = [];
    for(let move=0; move<64; move++){
        let flippable = get_flippable_discs(board, move);
        if(flippable.length > 0){
            valid_moves.push(move);
        }
    }
    return valid_moves;
}

function get_player_to_move(state) {
    if(state.board.turn === 0){
      return state.players.black;
    }
    return state.players.white;
}

function request_ws_move(){

    let message;

    switch(get_player_to_move(state)){
        case 'human':
            break;
        case 'bot':
            message = {
                'event': 'bot_move_request',
                'data': {
                    'state': state.board
                }
            };
            ws.send(JSON.stringify(message));
            break;
        case 'analyzer':
            message = {
                'event': 'analyze_move_request',
                'data': {
                    'state': state.board
                }
            };
            ws.send(JSON.stringify(message));
            break;
    }
}

function request_analysis_stop(){

    if(get_player_to_move(state) !== "analyzer"){
        return;
    }

    let message = {
        'event': 'analyze_stop_request'
    };

    ws.send(JSON.stringify(message));
}

function arraysEqual(a, b) {

    if (a === b) {
        return true;
    }

    if (a == null || b == null || a.length !== b.length) {
        return false;
    }

    for (let i = 0; i < a.length; ++i) {
        if(!(
            a.includes(b[i]) &&
            b.includes(a[i])
        )){
            return false;
        }
    }
    return true;
}

let ws;

let start_board = {
    'white': [27, 36],
    'black': [28, 35],
    'turn': 0
};

let board_history = [];

let state = {
    'board': {},
    'players': {
        'white': 'human',
        'black': 'human'
    }
};

$(function(){
    for(let y=0; y<8;y++){
        let row = $('<tr></tr>');
        $('#board').append(row);
        for (let x=0; x<8; x++){
            $(row, 'tr').append('<td></td>');
        }
    }

    $('#board td').append('<img src="./svg/field" />');

    // deep copy
    state.board = JSON.parse(JSON.stringify(start_board));

    update_fields(state.board);

    if (ws) {
        return false;
    }

    let ws_protocol = 'wss://';
    if(window.location.protocol == 'http:') {
        ws_protocol =  'ws://';
    }

    ws = new WebSocket(ws_protocol + window.location.host + window.location.pathname + 'ws');

    ws.onclose = function(evt) {
    };

    ws.onmessage = function(evt) {
        console.log('RESPONSE: ' + evt.data);
        let message = JSON.parse(evt.data);
        switch(message.event){
            case 'bot_move_reply':
                board_history.push(JSON.parse(JSON.stringify(state.board)));
                state.board = message.data.state;
                update_fields(state.board);
                if(get_valid_moves(state.board).length === 0){
                    state.board.turn = 1-state.board.turn;
                    if(get_valid_moves(state.board).length !== 0){
                        request_ws_move();
                    }
                } else {
                    request_ws_move();
                }
                break;
            case 'analyze_move_reply':
                if(!(
                    arraysEqual(state.board.white, message.data.board.white) &&
                    arraysEqual(state.board.black, message.data.board.black) &&
                    state.turn === message.data.turn)){

                    console.warn("Received outdated analyze_move_reply message:", message);
                    break;
                }

                let move = message.data.move;
                let heuristic = message.data.heuristic;

                let img_url = window.location.protocol + "//" + window.location.host + window.location.pathname + 'svg/field?text=' + heuristic;
                console.log(state.board.turn);
                if(state.board.turn === 1){
                    img_url += "&textcolor=white"
                }

                $('#board img').eq(move).attr('src', img_url);
                break;
            case 'xot_reply':
                state.board = message.data.state;
                update_fields(state.board);
                request_ws_move();
                break;
            default:
                console.warn("Unhandled ws event ", message.event)
        }
    };
    ws.onerror = function(evt) {
        console.log('ERROR: ' + evt.data);
    };
    return false;
});

$(document).on('mousedown', '#board td', function () {
    let y = $(this).parent().index();
    let x = $(this).index();
    let cell_id = 8*y + x;

    if(get_player_to_move(state) === 'bot'){
        return false;
    }

    let flipped = get_flippable_discs(state.board, cell_id);

    if(flipped.length === 0){
        return false;
    }

    request_analysis_stop();
    board_history.push(JSON.parse(JSON.stringify(state.board)));

    if(state.board.turn === 0){
        state.board.black.push(cell_id, ...flipped);
        state.board.white = state.board.white.filter(x => !flipped.includes(x));
    } else {
        state.board.white.push(cell_id, ...flipped);
        state.board.black = state.board.black.filter(x => !flipped.includes(x));
    }
    state.board.turn = 1-state.board.turn;

    if(get_valid_moves(state.board).length === 0){
        state.board.turn = 1-state.board.turn;
    }

    update_fields(state.board);
    request_ws_move();

    return false;
});

$(document).on('change', 'select', function() {
    let selected = $(this).val();
    let name = $(this).attr('name');

    switch(name){
        case 'white_player':
            state.players.white = selected;
            break;
        case 'black_player':
            state.players.black = selected;
            break;
        default:
            console.log('Unhandled name: ', name);
            return false;
    }

    request_analysis_stop();
    request_ws_move();
});

$(document).on('click', 'button#new_game', function(){

    // deep copy
    state.board = JSON.parse(JSON.stringify(start_board));
    board_history = [];

    request_analysis_stop();
    update_fields(state.board);
    request_ws_move();
});

$(document).on('click', 'button#xot_game', function(){

    let message = {
        'event': 'xot_request'
    };

    ws.send(JSON.stringify(message));

    request_analysis_stop();
    update_fields(state.board);
    request_ws_move();
});

$(document).on('click', 'button#undo_move', function(){

    for(let i = board_history.length - 1; i >= 0; i--){
        let turn = board_history[i].turn;
        if((turn === 1 && state.players.white !== 'bot') || (turn === 0 && state.players.black !== 'bot')){
            request_analysis_stop();
            state.board = board_history[i];
            board_history = board_history.slice(0, i);
            update_fields(state.board);
            request_ws_move();
            return;
        }
    }

});
