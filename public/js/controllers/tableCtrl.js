angular.module('orderchef')
.controller('TableCtrl', ['$scope', '$http', 'TestService', function ($scope, $http, TestService) {
	var tests = [];
	tests.push({
		name: "Tables",
		tests: [{
			name: "Add Table",
			test: function (done) {
				$http.post('/api/config/table-types', {
					name: "Table Type"
				}).success(function (data) {
					$http.get('/api/config/table-types')
					.success(function (data) {
						$http.post('/api/tables', {
							type_id: data[0].id,
							name: "Test Table",
							table_number: "two",
							location: "Internets"
						}).success(function (data) {
							done(true, []);
						}).error(function () {
							done(false, []);
						});
					})
				}).error(function (data) {
					done(false, data);
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
				$http.get('/api/table/' + tests[0].tests[1].results[0].id)
				.success(function (data) {
					done(true, data);
				}).error(function (data) {
					done(false, data);
				});
			}
		}, {
			name: "Remove All",
			test: function (done) {
				var tables = tests[0].tests[1].results;
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

	TestService.runTests($scope, tests);
}])