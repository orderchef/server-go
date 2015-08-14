var redis = require('redis').createClient();
var id = "receipt"

redis.sadd("oc_printers", id);
redis.subscribe("oc_print." + id);

redis.on("message", function (chan, message) {
	var obj = JSON.parse(message);
	console.log(obj);
	console.log(obj.print);
});