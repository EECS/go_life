window.onload = function () {
    var conn;
    var msgInputBox = document.getElementById("msg");
    var log = document.getElementById("log");

    function appendLog(item) {
      var doScroll =
        log.scrollTop > log.scrollHeight - log.clientHeight - 1;
      log.appendChild(item);
      if (doScroll) {
        log.scrollTop = log.scrollHeight - log.clientHeight;
      }
    }

    document.getElementById("form").onsubmit = function () {
      if (!conn) {
        return false;
      }
      if (!msgInputBox.value) {
        return false;
      }
      addMessage(msgInputBox.value);
      conn.send(msgInputBox.value);      
      msgInputBox.value = "";
      return false;
    };

    var id = Date.now();

    if (window["WebSocket"]) {
        //ws://localhost:8080/ws/123?v=1.0
      conn = new WebSocket("ws://" + document.location.host + "/ws/" + id);
      conn.onclose = function (evt) {
        var item = document.createElement("div");
        item.innerHTML = "<b>Connection closed.</b>";
        appendLog(item);
      };
      conn.onmessage = function (evt) {
        console.log("msg:");
        console.log(evt);
        var messages = evt.data.split("\n");
        for (var i = 0; i < messages.length; i++) {
          addMessage(messages[i]);
        }
      };
    } else {
      var item = document.createElement("div");
      item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
      appendLog(item);
    }

    function addMessage(msg){
      console.log(msg);
      var item = document.createElement("div");
          item.innerText = msg;
          appendLog(item);
    }

  };

