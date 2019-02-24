$(function(){
    for(let y=0; y<8;y++){
        let row = $("<tr></tr>");
        for (let x=0; x<8; x++){
            $(row, "tr").append("<td></td>");
        }
        $("#board").append(row);
    }
});

$(document).on("click", "#board td", function () {
    let y = $(this).parent().index();
    let x = $(this).index();
    let cell_id = 8*y + x;
    console.log(cell_id);
    return false;
});