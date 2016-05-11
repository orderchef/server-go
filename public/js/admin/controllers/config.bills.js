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
	$scope.kitchen_receipt_quantity = {
		name: 'kitchen_receipt_quantity',
		value: ''
	};
	$scope.customer_bill_quantity = {
		name: 'customer_bill_quantity',
		value: ''
	};

	$http.get('/config/settings/kitchen_receipt').success(function (receipt) {
		$scope.kitchen_receipt = receipt;
	});
	$http.get('/config/settings/customer_bill').success(function (receipt) {
		$scope.customer_bill = receipt;
	});
	$http.get('/config/settings/kitchen_receipt_quantity').success(function (receipt) {
		receipt.value = parseInt(receipt.value);
		$scope.kitchen_receipt_quantity = receipt;
	});
	$http.get('/config/settings/customer_bill_quantity').success(function (receipt) {
		receipt.value = parseInt(receipt.value);
		$scope.customer_bill_quantity = receipt;
	});

	$scope.save = function () {
		var krq = JSON.parse(JSON.stringify($scope.kitchen_receipt_quantity));
		krq.value += "";
		var cbq = JSON.parse(JSON.stringify($scope.customer_bill_quantity));
		cbq.value += "";


		$http.put('/config/settings/kitchen_receipt_quantity', krq)
		.success(function () {
			$http.put('/config/settings/kitchen_receipt', $scope.kitchen_receipt).error(function () {
				alert('Cannot save Kitchen Receipt');
			});
		}).error(function () {
			alert('Cannot save Kitchen Receipt');
		});

		$http.put('/config/settings/customer_bill_quantity', cbq).success(function () {
			$http.put('/config/settings/customer_bill', $scope.customer_bill).error(function () {
				alert('Cannot save Customer Bill');
			});
		}).error(function () {
			alert('Cannot save Customer Bill');
		});
	}
});