angular.module('orderchef')
.controller('ConfigModifiersCtrl', ['$scope', '$http', 'TestService', function ($scope, $http, TestService) {
	var tests = [];
	tests.push({
		name: "Config Modifiers",
		tests: [{
			name: "Add",
			test: function (done) {
				$http.post('/config/modifiers', {
					name: "Modifier Group",
					required: false
				}).success(function (data) {
					done(true, data);
				}).error(errorCb(done));
			}
		}, {
			name: "Get all",
			test: function (done) {
				$http.get('/config/modifiers')
				.success(function (data) {
					done(true, data);
				}).error(errorCb(done));
			}
		}, {
			name: "Get Single",
			test: function (done) {
				console.log(tests[0].tests[1])
				$http.get('/config/modifier/' + tests[0].tests[1].results[0].id)
				.success(function (data) {
					done(true, data);
				}).error(errorCb(done));
			}
		}, {
			name: "Add Modifier to group",
			test: function (done) {
				$http.post('/config/modifier/' + tests[0].tests[1].results[0].id + '/items', {
					name: "Medium",
					price: 0
				}).success(function (data) {
					done(true, data);
				}).error(errorCb(done));
			}
		}, {
			name: "Get all modifiers",
			test: function (done) {
				$http.get('/config/modifier/' + tests[0].tests[1].results[0].id + '/items')
				.success(function (data) {
					done(true, data);
				}).error(errorCb(done));
			}
		}, {
			name: "Get single modifier",
			test: function (done) {
				$http.get('/config/modifier/' + tests[0].tests[1].results[0].id + '/item/' + tests[0].tests[4].results[0].id)
				.success(function (data) {
					done(true, data);
				}).error(errorCb(done));
			}
		}, {
			name: "Update single modifier",
			test: function (done) {
				tests[0].tests[4].results[0].name = "Large";
				$http.put('/config/modifier/' + tests[0].tests[1].results[0].id + '/item/' + tests[0].tests[4].results[0].id, tests[0].tests[4].results[0])
				.success(function (data) {
					done(true, data);
				}).error(errorCb(done));
			}
		}, {
			name: "Get single modifier",
			test: function (done) {
				$http.get('/config/modifier/' + tests[0].tests[1].results[0].id + '/item/' + tests[0].tests[4].results[0].id)
				.success(function (data) {
					done(true, data);
				}).error(errorCb(done));
			}
		}, {
			name: "Remove single modifier",
			test: function (done) {
				$http.delete('/config/modifier/' + tests[0].tests[1].results[0].id + '/item/' + tests[0].tests[4].results[0].id)
				.success(function (data) {
					done(true, data);
				}).error(errorCb(done));
			}
		}, {
			name: "Get all modifiers (should be none)",
			test: function (done) {
				$http.get('/config/modifier/' + tests[0].tests[1].results[0].id + '/items')
				.success(function (data) {
					if (data.length != 0) return done(false, data);
					done(true, data);
				}).error(errorCb(done));
			}
		}, {
			name: "Delete all",
			test: function (done) {
				var objs = tests[0].tests[1].results;
				async.eachSeries(objs, function (obj, cb) {
					$http.delete('/config/modifier/' + obj.id)
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