
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

var iframe_port = 8081;

function setupiFrame() {
	// setup the iframe
	$.get("/api/domain", function(data) {
		console.log("Setting iFrame src to " +"http://" + data + ":" + iframe_port );
		document.getElementById('cameraFrame').src = "http://" + data + ":" + iframe_port;
	});
}

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

// do things on keypresses!
$(document).keydown(function(e) {

	switch(e.which) {
		case 32: // stop
			move('stop');
		break;

		case 37: // left
			move('left');
		break;

		case 38: // up
			move('forwards');
		break;

		case 39: // right
			move('right');
		break;

		case 40: // down
			move('reverse');
		break;

		case 82: // r
			resetServos();
		break;

		case 87: // w
			incServo('tilt', +5);
		break;

		case 65: // a
			incServo('pan', -5);
		break;

		case 83: // s
			incServo('tilt', -5);
		break;

		case 68: // d
			incServo('pan', +5);
		break;

		default: return; // exit this handler for other keys
	}

	e.preventDefault(); // prevent the default action (scroll / move caret)
});
