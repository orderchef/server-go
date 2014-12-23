
var app = angular.module('orderchef', [])

app.controller('TestCtrl', ['$scope', '$http', function($scope, $http) {
	$scope.tests = [];
	$scope.testResults = [];

	$scope.tests.push({
		name: "Tables",
		tests: [{
			name: "Add Table",
			test: function (done) {
				$http.post('/api/tables', {
					type_id: 1,
					name: "Test Table",
					table_number: "two",
					location: "Internets"
				}).success(function (data) {
					done(true, []);
				}).error(function () {
					done(false, []);
				});
			}
		}, {
			name: "Get All",
			test: function (done) {
				$http.get('/api/tables')
				.success(function (data) {
					done(true, data);
				});
			}
		}, {
			name: "Get Single",
			test: function (done) {
				$http.get('/api/table/' + $scope.testResults[1].results[0].id)
				.success(function (data) {
					done(true, data);
				})
				.error(function (data) {
					done(false, data);
				});
			}
		}, {
			name: "Remove All",
			test: function (done) {
				var tables = $scope.testResults[1].results;
				async.eachSeries(tables, function (table, cb) {
					$http.delete('/api/table/' + table.id)
					.success(function (data) {
						cb(null);
					}).error(function (data) {
						cb(data);
					});
				}, function(err) {
					if (err) {
						return done(false, err)
					}

					done(true, null);
				});
			}
		}]
	});

	$scope.tests.push({
		name: "Config Table Type",
		tests: [{
			name: "Get all table types",
			test: function (done) {
				$http.get('/api/config/table-types')
				.success(function (data) {
					done(true, data);
				})
			}
		}]
	});

	$scope.runTest = function (test, cb) {
		if (typeof test.tests == 'object') {
			for (var i = 0; i < test.tests.length; i++) {
				test.tests[i].name = test.name + ' » ' + test.tests[i].name;
			}

			return async.eachSeries(test.tests, $scope.runTest, cb);
		}

		test.test(function (success, results, err) {
			if (err) return cb(err);

			test.hasRun = true;
			test.success = success;
			test.results = results;
			if (Object.prototype.toString.call(test.results) != '[object Array]') {
				test.results = [results];
			}

			$scope.testResults.push(test);
			if (!$scope.$$phase) {
				$scope.$digest();
			}

			cb();
		});
	}

	async.eachSeries($scope.tests, $scope.runTest, function (err) {
		if (err) throw err;
	});
}])
