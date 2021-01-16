$(document).ready(function () {
  $("#main div").click(function () {
    if ($(this).html() !== "X" && $(this).html !== "O") {
      $(this).html("X");
    }
    readAndPost();
  });

  $("#replay").click(function () {
    $("#main div").each(function () {
      $(this).html("");
    });
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
  pos += '"}';

  console.log(pos);

  $.post("/api/", pos, function (data) {
    console.log(data);
    json = JSON.parse(data);

    $("#main div").each(function (index) {
      cellState = json.State[index];
      if (cellState !== "-") {
        $(this).html(cellState);
      }
    });
    console.log(json.Result);
  });
}
