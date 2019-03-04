update_fields = function(board) {

    let valid_moves = get_valid_moves(board);

    for(let i=0; i<64; i++){
        let image = '';

        if(board.white.includes(i)){
            image = 'white.svg';
        } else if(board.black.includes(i)){
            image = 'black.svg';
        } else if(valid_moves.includes(i)){
            if(board.turn === 0){
                image = 'move_black.svg';
            }  else {
                image = 'move_white.svg';
            }
        } else {
            image = 'empty.svg';
        }

        $('#board img').eq(i).attr('src', 'static/' + image);
    }
};

get_flippable_discs = function(board, move) {
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
};

get_valid_moves = function(board){
    let valid_moves = [];
    for(let move=0; move<64; move++){
        let flippable = get_flippable_discs(board, move);
        if(flippable.length > 0){
            valid_moves.push(move);
        }
    }
    return valid_moves;
};

get_player_to_move = function(state) {
    if(state.board.turn === 0){
      return state.players.black;
    }
    return state.players.white;
};

request_ws_move = function(){

    let message;

    switch(get_player_to_move(state)){
        case 'human':
            break;
        case 'bot':
            message = {
                'event': 'bot_move',
                'bot_move': {
                    'state': state.board
                }
            };
            ws.send(JSON.stringify(message));
            break;
        case 'analyzer':
            message = {
                'event': 'analyze_move',
                'analyze_move': {
                    'state': state.board
                }
            };
            ws.send(JSON.stringify(message));
            break;
    }
};

let ws;

let start_board = {
    'white': [27, 36],
    'black': [28, 35],
    'turn': 0
};

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

    $('#board td').append('<img src="static/empty.svg" />');

    // deep copy
    state.board = JSON.parse(JSON.stringify(start_board));

    update_fields(state.board);

    if (ws) {
        return false;
    }
    ws = new WebSocket('ws://localhost:8080/ws');

    ws.onclose = function(evt) {
        // TODO try to reconnect
    };

    ws.onmessage = function(evt) {
        console.log('RESPONSE: ' + evt.data);
        let message = JSON.parse(evt.data);
        switch(message.event){
            case 'bot_move_reply':
                state.board = message.bot_move_reply.state;
                update_fields(state.board);
                if(get_valid_moves(state.board).length === 0){
                    state.board.turn = 1-state.board.turn;
                    if(get_valid_moves(state.board).length !== 0){
                        request_ws_move();
                    }
                } else {
                    request_ws_move();
                }
            case 'analyze_move_reply':
                let move = message.analyze_move_reply.move;
                let heuristic = message.analyze_move_reply.heuristic;
                $('#board img').eq(move).attr('src', window.location.origin + '/svg/?text=' + heuristic);
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

    request_ws_move();
});

$(document).on('click', 'button', function(){

    // deep copy
    state.board = JSON.parse(JSON.stringify(start_board));

    update_fields(state.board);
    request_ws_move();
});