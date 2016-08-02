"use strict";

var apiUrl = "http://localhost:4444/"

window.addEventListener('load',(function(e){

	$("#message").innerText = "Now Loading...";
	getJson(apiUrl + "status", function(r){
		if (!r) {
			$("#message").innerText = "device not found.";
			return;
		}
		if (!r.serial_connected) {
			$("#message").innerText = "Disconnected";
			return;
		}
		$("#message").innerText = "Connected! X=" + r.x + "  Y=" + r.y;
	});
}),false);
