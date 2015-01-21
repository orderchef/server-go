angular.module('orderchef')
.controller('ConfigCtrl', ['$scope', '$http', 'TestService', function ($scope, $http, TestService) {
	var tests = [];
	tests.push({
		name: "Config Table Type",
		tests: [{
			name: "Add",
			test: function (done) {
				$http.post('/config/table-types', {
					name: "Table Type"
				}).success(function (data) {
					done(true, data);
				}).error(function (data) {
					done(false, data);
				});
			}
		}, {
			name: "Get all",
			test: function (done) {
				$http.get('/config/table-types')
				.success(function (data) {
					done(true, data);
				}).error(function (data) {
					done(false, data);
				});
			}
		}, {
			name: "Get Single",
			test: function (done) {
				$http.get('/config/table-types/' + tests[0].tests[1].results[0].id)
				.success(function (data) {
					done(true, data);
				}).error(function (data) {
					done(false, data);
				});
			}
		}, {
			name: "Delete all",
			test: function (done) {
				var objs = tests[0].tests[1].results;
				async.eachSeries(objs, function (obj, cb) {
					$http.delete('/config/table-types/' + obj.id)
					.success(function (data) {
						cb(null);
					}).error(function (data) {
						cb(new Error(obj.id + " Could not be deleted: " + data));
					});
				}, function(err) {
					if (err) {
						return done(false, { response: err.toString() })
					}

					done(true, null);
				});
			}
		}]
	});

	TestService.runTests($scope, tests);
}]);