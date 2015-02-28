
angular.module('orderchef')
.controller('ItemsCtrl', ['$scope', '$http', 'TestService', function ($scope, $http, TestService) {
	var tests = [];
	tests.push({
		name: "Items",
		tests: [{
			name: "Add Item",
			test: function (done) {
				$http.post('/categories', {
					name: 'hello'
				}).success(function (data) {
					$http.get('/categories').success(function (categories) {
						$http.post('/items', {
							name: "name.",
							description: "desc.",
							price: 2.0,
							category_id: categories[0].id
						}).success(function (data) {
							done(true, []);
						}).error(errorCb(done));
					}).error(errorCb(done));
				}).error(errorCb(done));
			}
		}, {
			name: "Get All",
			test: function (done) {
				$http.get('/items').success(function (data) {
					done(true, data);
				}).error(errorCb(done));
			}
		}, {
			name: "Get Single",
			test: function (done) {
				$http.get('/item/' + tests[0].tests[1].results[0].id).success(function (data) {
					done(true, data);
				}).error(errorCb(done));
			}
		}, {
			name: "Remove All",
			test: function (done) {
				var categories = tests[0].tests[1].results;
				async.eachSeries(categories, function (category, cb) {
					$http.delete('/item/' + category.id).success(function (data) {
						cb(null);
					}).error(errorCb(done));
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
