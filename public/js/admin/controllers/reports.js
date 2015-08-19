var app = angular.module('orderchef');

app.controller('PastOrdersReportCtrl', function ($scope, $http, reportDates) {
	$scope.dates = reportDates;

	reportDates.setup();
});

app.controller('SalesReportCtrl', function ($scope, $http, reportDates) {
	$scope.dates = reportDates;

	reportDates.setup();
});

app.controller('PopularDishesReportCtrl', function ($scope, $http, reportDates) {
	$scope.dates = reportDates;

	reportDates.setup();
});

app.controller('CashReportCtrl', function ($scope, $http, reportDates) {
	$scope.dates = reportDates;

	reportDates.setup();
});
