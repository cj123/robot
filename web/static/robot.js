$(document).ready(function() {
	// get the readings at an interval of 1s
	setInterval(function() {
		// sonar
		$.get("/api/sonar/distance", function(data) {
			$("td#sonar").text(data + "cm");
		});

		// infrared readings
		$.getJSON("/api/ir", function(response) {
			$("td#irleft").text(response.left);
			$("td#irright").text(response.right);
			$("td#irleftl").text(response.leftline);
			$("td#irrightl").text(response.rightline);
		})

		// get the pan and tilt values of the servos
		$.get("/api/servos/pan/get", function(data) {
			$("span#servo-pan").text(data);
		});

		$.get("/api/servos/tilt/get", function(data) {
			$("span#servo-tilt").text(data);
		});
	}, 1000);
});

function move(direction) {
	$.get("/api/motors/" + direction);
}

function incServo(name, val) {
	$.get("/api/servos/" + name + "/inc/" + val);
}

function resetServos() {
	$.get("/api/servos/tilt/reset");
	$.get("/api/servos/pan/reset");
}