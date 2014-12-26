angular.module('orderchef')
.controller('ConfigCtrl', ['$scope', '$http', 'TestService', function ($scope, $http, TestService) {
	var tests = [];
	tests.push({
		name: "Config Table Type",
		tests: [{
			name: "Add",
			test: function (done) {
				$http.post('/api/config/table-types', {
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
				$http.get('/api/config/table-types')
				.success(function (data) {
					done(true, data);
				}).error(function (data) {
					done(false, data);
				});
			}
		}, {
			name: "Get Single",
			test: function (done) {
				$http.get('/api/config/table-type/' + tests[0].tests[1].results[0].id)
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
					$http.delete('/api/config/table-type/' + obj.id)
					.success(function (data) {
						cb(null);
					}).error(function (data) {
						cb(data);
					});
				}, function(err) {
					if (err) {
						return done(false, { response: err })
					}

					done(true, null);
				});
			}
		}]
	});

	TestService.runTests($scope, tests);
}]);