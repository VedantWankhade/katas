const stompClient = new StompJs.Client({
    brokerURL: 'ws://localhost:8080/gs-guide-websocket'
});

stompClient.onConnect = frame => {
    setConnected(true);
    console.log("connected " + frame);
    stompClient.subscribe('/topic/messages', msg => {

        let text = new TextDecoder("utf-8").decode(msg._binaryBody);
        console.log("Received:", text);
        pushMsg(text);
    });
}

stompClient.onWebSocketError = err => {
    console.error('ERR: ' + err);
}

stompClient.onStompError = frame => {
    console.error("ERR: " + frame.headers['message']);
    console.error("ERR: " + frame.body);
}

function setConnected(connected) {
    document.getElementById('connect').disabled = connected;
    document.getElementById('disconnect').disabled = !connected;

    if (connected) {
        document.getElementById("msg").style.display = "block";
        document.getElementById("q").style.display = "block";
    }
    else {
        document.getElementById("msg").style.display = "none";
        document.getElementById("q").style.display = "none";
    }
}

function connect() {
    stompClient.activate();
}

function disconnect() {
    stompClient.deactivate();
    setConnected(false);
    console.log("Disconnected");
}

function pushMsg(msg) {
    let li = document.createElement("li");
    li.appendChild(document.createTextNode(msg));
    document.getElementById("messages").appendChild(li);
}

document.getElementById('connect').addEventListener('click', ev => {
    connect();
})
document.getElementById('disconnect').addEventListener('click', ev => {
    disconnect();
})