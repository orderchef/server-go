
angular.module('orderchef')
.controller('CategoriesCtrl', ['$scope', '$http', 'TestService', function ($scope, $http, TestService) {
	var tests = [];
	tests.push({
		name: "Categories",
		tests: [{
			name: "Add Category",
			test: function (done) {
				$http.post('/categories', {
					name: "name.",
					description: "desc."
				}).success(function (data) {
					done(true, []);
				}).error(function () {
					done(false, []);
				});
			}
		}, {
			name: "Get All",
			test: function (done) {
				$http.get('/categories').success(function (data) {
					done(true, data);
				}).error(function () {
					done(false, []);
				});
			}
		}, {
			name: "Get Single",
			test: function (done) {
				$http.get('/categories/' + tests[0].tests[1].results[0].id).success(function (data) {
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
					$http.delete('/categories/' + category.id).success(function (data) {
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
