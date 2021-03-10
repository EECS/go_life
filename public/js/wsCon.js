var conn;

if (window["WebSocket"]) {
    //ws://localhost:8080/ws/123?v=1.0
  conn = new WebSocket("ws://" + document.location.host + "/ws/");

  conn.onclose = function (evt) {
    output("!! -- Connection closed. -- !!");
  };

  conn.onopen = function (evt) {
    output("!! -- Connection opened. -- !!");
  }

  conn.onmessage = function (evt) {
    var messages = evt.data.split("\n");
    for (var i = 0; i < messages.length; i++) {
      messageReceived(messages[i]);
    }
  };

} else {
  output("!! -- Your browser does not support WebSockets. -- !!");
}

function messageReceived(msg){
  output(msg);
  
  // defined in index.js
  handleMessage(msg);
}

function sendMessage(msg){
  output(msg);
  conn.send(msg);
}

function output(str){
  console.log(str);
}