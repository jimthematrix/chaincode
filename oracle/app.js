var express = require('express'),
	request = require('request-promise');

var app = express();

var counter = 0;
var storage = {};

app.get("/oracle", function(req, res) {
	console.log(req.originalUrl);

	var url = req.query.url;
	var transactionId = req.query.transactionId;

	if (!url)
		res.status(404).send("Request to the oracle service must include query parameter 'url'")
	else {
		var value;
		if (transactionId && transactionId != "") {
			value = storage[transactionId];

			if (value != null) {
				res.status(200).send(formatResponse(value));
				return;
			}
		}

		request({
			uri: url
		}).then(function(body) {
			if (transactionId && transactionId != "") {
				storage[transactionId] = body;
			}

			res.status(200).send(formatResponse(body));
		}).catch(function(err) {
			res.status(500).send(err);
		});
	}
})

function formatResponse(result) {
	return '{"name": "oracle service", "result": ' + result + '}';
}

app.listen(3010, function() {
	console.log("API Server running and listening on port 3010!");
})