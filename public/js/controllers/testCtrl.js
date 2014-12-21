
var app = angular.module('orderchef', [])

app.controller('TestCtrl', ['$scope', '$http', function($scope, $http) {
	$scope.tests = [];

	$scope.tests.push({
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
		name: "Get Tables",
		test: function (done) {
			$http.get('/api/tables')
			.success(function (data) {
				done(true, data);
			});
		}
	}, {
		name: "Remove Table",
		test: function (done) {
			$http.delete('/api/table/' + $scope.tests[1].results[0].id)
			.success(function (data) {
				done(true, data);
			})
			.error(function (data) {
				done(false, data)
			});
		}
	});

	async.eachSeries($scope.tests, function (test, cb) {
		test.test(function (success, results, err) {
			if (err) return cb(err);

			test.hasRun = true;
			test.success = success;
			test.results = results;

			cb();
		});
	}, function (err) {
		if (err) throw err;
	});
}])
