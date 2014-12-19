
var app = angular.module('orderchef', [])

app.controller('TableCtrl', ['$scope', '$http', function($scope, $http) {
	$scope.tables = [];

	$http.get('/api/tables')
	.success(function (data) {
		$scope.tables = data
	})
}])
