var app = angular.module('orderchef');

app.controller('ConfigBillsCtrl', function ($scope, $http) {
	$scope.kitchen_receipt = {
		name: 'kitchen_receipt',
		value: ''
	};
	$scope.customer_bill = {
		name: 'customer_bill',
		value: ''
	};
	$scope.restaurant_bill = {
		name: 'restaurant_bill',
		value: ''
	};

	$http.get('/config/settings/kitchen_receipt').success(function (receipt) {
		$scope.kitchen_receipt = receipt;
	});
	$http.get('/config/settings/customer_bill').success(function (receipt) {
		$scope.customer_bill = receipt;
	});
	$http.get('/config/settings/restaurant_bill').success(function (receipt) {
		$scope.restaurant_bill = receipt;
	});

	$scope.save = function () {
		$http.put('/config/settings/kitchen_receipt', $scope.kitchen_receipt).error(function () {
			alert('Cannot save Kitchen Receipt');
		});
		$http.put('/config/settings/customer_bill', $scope.customer_bill).error(function () {
			alert('Cannot save Customer Bill');
		});
		$http.put('/config/settings/restaurant_bill', $scope.restaurant_bill).error(function () {
			alert('Cannot save Restaurant Bill');
		});
	}
});