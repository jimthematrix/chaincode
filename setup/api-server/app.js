var express = require('express');

var app = express();

var counter = 0;

app.get("/", function(req, res) {
	res.setHeader('Content-Type', 'application/json');
	res.status(200).send(
		JSON.stringify(
			{
				result: counter++
			}
		)
	);
})

app.listen(3000, function() {
	console.log("API Server running and listening on port 3000!");
})