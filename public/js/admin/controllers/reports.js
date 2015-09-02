var app = angular.module('orderchef');

app.controller('ReportBillsCtrl', function ($scope, $http, reportDates) {
	$scope.dates = reportDates;

	reportDates.setup();

	$http.get('/config/payment-methods').success(function (payment_methods) {
		$scope.payment_methods = payment_methods;
		$scope.refreshData();
	});

	$scope.refreshData = function () {
		$http.get('/reports/bills' + reportDates.getQuery()).success(function(bills) {
			$scope.total = bills.total;
			$scope.bills = bills.bills;

			for (var i = 0; i < $scope.bills.length; i++) {
				for (var x = 0; x < $scope.payment_methods.length; x++) {
					if ($scope.bills[i].payment_method_id == $scope.payment_methods[x].id) {
						$scope.bills[i].payment_method = $scope.payment_methods[x];
					}
				}
			}
		})
	}
});