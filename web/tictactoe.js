$(document).ready(function () {
  $("#main div").click(function () {
    if ($(this).html() !== "X" && $(this).html() !== "O") {
      $(this).html("X");
      readAndPost();
    }
  });

  $("#replay").click(function () {
    $("#main div").each(function () {
      $(this).html("");
    });
    $("#result").html("")
  });
});

function readAndPost() {
  //loop over divs and post board position to the backend
  pos = '{"State":"';
  $("#main div").each(function () {
    line = "-";
    if ($(this).html() !== undefined && $(this).html() !== "") {
      line = $(this).html();
    }

    pos += line;
  });
  pos += '", "Player":"'
  pos += $('input[name="player"]:checked').val(); 
  pos += '"}';

  console.log(pos);

  $.post("/api/", pos, function (data) {
    jsonData = JSON.parse(data);

    $("#main div").each(function (index) {
      cellState = jsonData.State[index];
      if (cellState !== "-") {
        $(this).html(cellState);
      }
    });
    if(jsonData.Result !== "NotEnd"){
        if(jsonData.Result === "XWon"){
            $("#result").html("Player X won!")
        }
        if(jsonData.Result === "OWon"){
            $("#result").html("Player O won!")
        }
        if(jsonData.Result === "Tie"){
            $("#result").html("Tie!")
        }    
    }
  });
}
