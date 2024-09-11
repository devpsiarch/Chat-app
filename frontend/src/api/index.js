//connetting to the server websocket
var socket = new WebSocket("ws://localhost:8080/ws");


//this is kind if a switch statement , each 'event' to has a response to
let connect = cb => {
	console.log("attempting connection ...");

	socket.onopen = () => {
		console.log("connection success!");
	};

	socket.onmessage = msg => {
		console.log(msg);
        cb(msg);
	};

	socket.onclose = event => {
		console.log("socket connection closed",event);
	};

	socket.onerror = error => {
		console.log("socket error: ",error);
	};
};


let send_msg = msg => {
	console.log("sending message: ",msg);
	socket.send(msg);
};


export {send_msg , connect};
