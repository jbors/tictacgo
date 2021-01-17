$(document).ready(function () {
  $("#main div").click(function () {
    if ($(this).html() !== "X" && $(this).html() !== "O" && $("#result").html() === "") {
      $(this).html("X");
      readAndPost();
    }
  });

  $("#replay").click(function () {
    $("#main div").each(function () {
      $(this).html("");
      $(this).removeClass("win")
    });
    $("#result").html("");
  });
});

function readAndPost() {
  //Build the board representation
  pos = "";
  $("#main div").each(function () {
    line = "-";
    if ($(this).html() !== undefined && $(this).html() !== "") {
      line = $(this).html();
    }
    pos += line;
  });
  player = $('input[name="player"]:checked').val();

  payload = JSON.stringify({ State: pos, Player: player });

  console.log(payload);

  $.post("/api/", payload, function (data) {
    jsonData = JSON.parse(data);

    $("#main div").each(function (index) {
      cellState = jsonData.State[index];
      if (cellState !== "-") {
        $(this).html(cellState);
      }
    });
    if (jsonData.Result !== "NotEnd") {
      if (jsonData.Result === "XWon") {
        $("#result").html("Player X won!");
        displayWin();
      }
      if (jsonData.Result === "OWon") {
        $("#result").html("Player O won!");
        displayWin();
      }
      if (jsonData.Result === "Tie") {
        $("#result").html("Tie!");
      }
    }
  });

  function displayWin() {
    var box1 = $("#box1"),
      box2 = $("#box2"),
      box3 = $("#box3"),
      box4 = $("#box4"),
      box5 = $("#box5"),
      box6 = $("#box6"),
      box7 = $("#box7"),
      box8 = $("#box8"),
      box9 = $("#box9");

    if (
      box1.html() !== "" &&
      box1.html() === box2.html() &&
      box1.html() === box3.html()
    )
      selectWinnerBoxes(box1, box2, box3);

    if (
      box4.html() !== "" &&
      box4.html() === box5.html() &&
      box4.html() === box6.html()
    )
      selectWinnerBoxes(box4, box5, box6);

    if (
      box7.html() !== "" &&
      box7.html() === box8.html() &&
      box7.html() === box9.html()
    )
      selectWinnerBoxes(box7, box8, box9);

    if (
      box1.html() !== "" &&
      box1.html() === box4.html() &&
      box1.html() === box7.html()
    )
      selectWinnerBoxes(box1, box4, box7);

    if (
      box2.html() !== "" &&
      box2.html() === box5.html() &&
      box2.html() === box8.html()
    )
      selectWinnerBoxes(box2, box5, box8);

    if (
      box3.html() !== "" &&
      box3.html() === box6.html() &&
      box3.html() === box9.html()
    )
      selectWinnerBoxes(box3, box6, box9);

    if (
      box1.html() !== "" &&
      box1.html() === box5.html() &&
      box1.html() === box9.html()
    )
      selectWinnerBoxes(box1, box5, box9);

    if (
      box3.html() !== "" &&
      box3.html() === box5.html() &&
      box3.html() === box7.html()
    )
      selectWinnerBoxes(box3, box5, box7);
  }

  function selectWinnerBoxes(b1, b2, b3){
    b1.addClass("win");
    b2.addClass("win");
    b3.addClass("win");
  }
}
