// DAVE 2.0 Javascript
// (c) 2015 Callum Jones
// MIT Licensed - See LICENSE file attached

var crashCount = 0;
var crashTime = 0;
var cameraPort = 8090;

$(document).ready(function() {

	// setup the camera
	$.get("/api/domain", function(url) {
		document.getElementById('cameraFrame').src = "http://" + url + ":" + cameraPort;
	});

	setInterval(function() {

		if(getCollisionAvoidance() == true) {
			$("#collisionAvoidanceOn").removeClass("btn-default");
			$("#collisionAvoidanceOff").removeClass("btn-success");

			$("#collisionAvoidanceOn").addClass("btn-success");
			$("#collisionAvoidanceOff").addClass("btn-default");
		} else {
			$("#collisionAvoidanceOn").removeClass("btn-success");
			$("#collisionAvoidanceOff").removeClass("btn-default");

			$("#collisionAvoidanceOn").addClass("btn-default");
			$("#collisionAvoidanceOff").addClass("btn-success");
		}

		$.getJSON("/api/ir", function(response) {

			crashTime++

			if (response.frontleft || response.frontright || response.backleft || response.backright) {
				//move('stop'); // prevent a collision
				crashTime = 0;
				$("#crashPanel .panel-body").text(++crashCount + ": A crash was detected. We stopped you from breaking anything!");
				$("#crashPanel").show();
			} else if(crashTime == 50) {
				$("#crashPanel").hide();
			}
		}).fail(function() {
			// error, we timed out
			$("#timeoutAlert").show();
		});

		// update log panel
		$.get("/log", function(data) {
			$("#logOutput").text(data);
			tailScroll();
		});
	}, 100);

	// get the readings at an interval of 1s
	setInterval(function() {
		// sonar
		$.get("/api/sonar/distance", function(data) {
			$("td#sonar").text(data + "cm");
		});

		$(".collisionSwitch").prop("checked", getCollisionAvoidance());

		// infrared readings
		$.getJSON("/api/ir", function(response) {
			$("td#irleft").text(response.frontleft);
			$("td#irright").text(response.frontright);
			$("td#irbleft").text(response.backleft);
			$("td#irbright").text(response.backright);
			$("td#irleftl").text(response.leftline);
			$("td#irrightl").text(response.rightline);
		});

		// get the pan and tilt values of the servos
		$.get("/api/servos/pan/get", function(data) {
			$("span#servo-pan").text(data);
		});

		$.get("/api/servos/tilt/get", function(data) {
			$("span#servo-tilt").text(data);
		});
	}, 1000);

});

function enableCollisionAvoidance() {
	$.get("/api/collision/on");

	return false;
}

function disableCollisionAvoidance() {
	$.get("/api/collision/off");
	move('stop');
	resetServos();
	return false;
}

function getCollisionAvoidance() {
	var enabled = null;
	$.get("/api/collision/get", function(data) {
		enabled = data;
	}).complete(function() {
		return enabled === "true";
	});

	return enabled;
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

// keep at the bottom of the div
function tailScroll() {
    var height = $("#logOutput").get(0).scrollHeight;
    $("#logOutput").animate({
        scrollTop: height
    }, 500);
}
