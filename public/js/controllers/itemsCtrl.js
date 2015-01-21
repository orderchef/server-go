
angular.module('orderchef')
.controller('ItemsCtrl', ['$scope', '$http', 'TestService', function ($scope, $http, TestService) {
	var tests = [];
	tests.push({
		name: "Items",
		tests: [{
			name: "Add Item",
			test: function (done) {
				$http.post('/items', {
					name: "name.",
					description: "desc.",
					price: 2.0
				}).success(function (data) {
					done(true, []);
				}).error(function () {
					done(false, []);
				});
			}
		}, {
			name: "Get All",
			test: function (done) {
				$http.get('/items').success(function (data) {
					done(true, data);
				}).error(function () {
					done(false, []);
				});
			}
		}, {
			name: "Get Single",
			test: function (done) {
				$http.get('/items/' + tests[0].tests[1].results[0].id).success(function (data) {
					done(true, data);
				}).error(function (data) {
					done(false, data);
				});
			}
		}, {
			name: "Remove All",
			test: function (done) {
				var categories = tests[0].tests[1].results;
				async.eachSeries(categories, function (category, cb) {
					$http.delete('/items/' + category.id).success(function (data) {
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
}]);
