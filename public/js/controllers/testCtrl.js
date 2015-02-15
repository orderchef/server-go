
var app = angular.module('orderchef', [])

function errorCb (cb) {
	return function () {
		cb(false, []);
	}
}

app.config(function ($httpProvider) {
	$httpProvider.interceptors.push(function ($q) {
		return {
			'request': function (config) {
				config.url = '/api' + config.url;
				return config || $q.when(config);
			}
		}
	});
});

app.service('TestService', function() {
	var self = this;

	self.tests = [];
	self.testResults = [];

	self.runTest = function (test, cb) {
		if (typeof test.tests == 'object') {
			for (var i = 0; i < test.tests.length; i++) {
				test.tests[i].name = test.name + ' » ' + test.tests[i].name;
			}

			return async.eachSeries(test.tests, self.runTest, cb);
		}

		test.test(function (success, results, err) {
			if (err) return cb(err);

			test.hasRun = true;
			test.success = success;
			test.results = results;
			if (Object.prototype.toString.call(test.results) != '[object Array]') {
				test.results = [results];
			}

			self.testResults.push(test);

			cb();
		});
	}

	self.runTests = function (scope, tests) {
		scope.testResults = self.testResults;
		self.tests = tests;
		async.eachSeries(self.tests, self.runTest, function (err) {
			if (err) throw err;
		});
	}

	return self;
})
