var app = angular.module('orderchef');

app.controller('ReportBillsCtrl', function ($scope, $http, reportDates) {
	$scope.dates = reportDates;

	reportDates.setup(function () {
		$scope.refreshData();
	});

	$http.get('/config/payment-methods').success(function (payment_methods) {
		$scope.payment_methods = payment_methods;
		$scope.refreshData();
	});

	$scope.refreshData = function () {
		$http.get('/reports/bills' + reportDates.getQuery()).success(function(bills) {
			$scope.total = bills.total;
			$scope.bills = bills.bills;
			$scope.unclearedTables = bills.unclearedTables;

			$scope.coversTotal = 0;

			for (var i = 0; i < $scope.bills.length; i++) {
				var bill = $scope.bills[i];
				bill.printed_atFormatted = moment(bill.printed_at).format('Do MMM YYYY hh:mm:ss');
				bill.totalFormatted = (Math.round(bill.total * 100) / 100).toFixed(2);

				$scope.coversTotal += bill.covers;

				for (var x = 0; x < $scope.payment_methods.length; x++) {
					if (bill.payment_method_id == $scope.payment_methods[x].id) {
						bill.payment_method = $scope.payment_methods[x];
					}
				}
			}
		})
	}
});

app.controller('BillExtrasCtrl', function ($scope, $http, reportDates) {
	$scope.dates = reportDates;

	reportDates.setup(function () {
		$scope.refreshData();
	});

	$http.get('/config/payment-methods').success(function (payment_methods) {
		$scope.payment_methods = payment_methods;
		$scope.refreshData();
	});

	$scope.refreshData = function () {
		$http.get('/reports/bills/extras' + reportDates.getQuery()).success(function(bills) {
			$scope.total = bills.total;
			$scope.bills = bills.bills;
			$scope.unclearedTables = bills.unclearedTables;

			$scope.coversTotal = 0;

			for (var i = 0; i < $scope.bills.length; i++) {
				var bill = $scope.bills[i];
				bill.printed_atFormatted = moment(bill.printed_at).format('Do MMM YYYY hh:mm:ss');
				bill.totalFormatted = (Math.round(bill.total * 100) / 100).toFixed(2);

				$scope.coversTotal += bill.covers;

				for (var x = 0; x < $scope.payment_methods.length; x++) {
					if (bill.payment_method_id == $scope.payment_methods[x].id) {
						bill.payment_method = $scope.payment_methods[x];
					}
				}
			}
		})
	}
});

app.controller('ReportCashCtrl', function ($scope, $http, $modal, $rootScope, reportDates, $interpolate) {
	$scope.dates = reportDates;

	reportDates.setup(function () {
		$scope.refreshData();
	});

	$http.get('/config/settings/cashup_categories').success(function (categories) {
		try {
			$scope.categories = JSON.parse(categories.value);
		} catch (e) {
			$scope.categories = [];
		}
	}).error(function () {
		$scope.categories = [];
	});

	$scope.formula = 'Not Set Up';
	$http.get('/config/settings/cashup_formula').success(function (formula) {
		$scope.formula = formula.value;
		$scope.refreshData();
	});

	$scope.refreshData = function () {
		$http.get('/reports/cash' + reportDates.getQuery()).success(function(cash) {
			$scope.cash = cash;

			var exp = $interpolate($scope.formula);
			$scope.gross = exp({ c: cash });
		});
	}

	$scope.settings = function () {
		var scope = $rootScope.$new();
		scope.categories = $scope.categories;

		var modal = $modal.open({
			templateUrl: '/public/html/admin/reports.cash.settings.html',
			scope: scope,
			controller: function ($scope, $http, $modalInstance) {
				$scope.formula = '';
				$http.get('/config/settings/cashup_formula').success(function (formula) {
					$scope.formula = formula.value;
				});

				$scope.save = function () {
					$http.put('/config/settings/cashup_categories', {
						name: 'cashup_categories',
						value: JSON.stringify($scope.categories)
					}).success(function () {
						$http.put('/config/settings/cashup_formula', {
							name: 'cashup_formula',
							value: $scope.formula
						}).success(function () {
							$modalInstance.close();
						}).error(function () {
							alert('Could not save settings');
						});
					}).error(function () {
						alert('Could not save settings');
					});
				}
			}
		});
	}

	$scope.createReport = function () {
		scope = $rootScope.$new();
		scope.categories = $scope.categories;

		var modal = $modal.open({
			scope: scope,
			templateUrl: '/public/html/admin/reports.cash.new.html',
			controller: function ($scope, $modalInstance) {
				$scope.report = {
					date: new Date(),
					data: {}
				};
				$scope.createReport = function (report) {
					$modalInstance.close(report);
				}
			}
		});

		modal.close = function (report) {
			var data = [];
			for (var key in report.data) {
				if (!report.data.hasOwnProperty(key)) continue;

				data.push({
					date: report.date,
					category: key,
					amount: report.data[key]
				});
			}

			async.eachSeries(data, function (rep, cb) {
				$http.post('/reports/cash', rep).success(function () {
					cb();
				}).error(function () {
					alert('Failed to add report');
					cb('');
				});
			}, $scope.refreshData);
		};
	}
});

app.controller('PopularItemsReportCtrl', function ($scope, $http, reportDates) {
	reportDates.setup(function () {
		$scope.refreshData();
	});

	$scope.refreshData = function () {
		$http.get('/reports/popularItems' + reportDates.getQuery()).success(function(popularItems) {
			$scope.popularItems = popularItems;
		});
	}

	$scope.refreshData();
});

app.controller('PopularItemsReportCategoriesCtrl', function ($scope, $http, reportDates) {
	reportDates.setup(function () {
		$scope.refreshData();
	});

	$scope.refreshData = function () {
		$http.get('/reports/popularItems' + reportDates.getQuery()).success(function(popularItems) {
			var categories = {};

			popularItems.forEach(function (popItem) {
				if (!categories[popItem.Category]) {
					categories[popItem.Category] = {
						items: [],
						name: popItem.Category
					}
				}

				categories[popItem.Category].items.push(popItem);
			});

			$scope.popularItems = [];
			for (var key in categories) {
				if (!categories.hasOwnProperty(key)) continue;

				$scope.popularItems.push(categories[key]);
			}
		});
	}

	$scope.refreshData();
});

app.controller('SalesReportCtrl', function ($scope, $http, $modal, $rootScope, reportDates, $interpolate) {
	$scope.dates = reportDates;

	reportDates.setup(function () {
		$scope.refreshData();
	});

	$http.get('/config/settings/sales_categories').success(function (categories) {
		try {
			$scope.categories = JSON.parse(categories.value);
		} catch (e) {
			$scope.categories = [];
		}
	}).error(function () {
		$scope.categories = [];
	});

	$scope.formula = 'Not Set Up';
	$http.get('/config/settings/sales_formula').success(function (formula) {
		$scope.formula = formula.value;
		$scope.refreshData();
	});

	$scope.refreshData = function () {
		$http.get('/reports/sales' + reportDates.getQuery()).success(function(cash) {
			$scope.cash = cash;

			var exp = $interpolate($scope.formula);
			$scope.gross = exp({ c: cash });
		});
	}

	$scope.settings = function () {
		var scope = $rootScope.$new();
		scope.categories = $scope.categories;

		var modal = $modal.open({
			templateUrl: '/public/html/admin/reports.cash.settings.html',
			scope: scope,
			controller: function ($scope, $http, $modalInstance) {
				$scope.formula = '';
				$http.get('/config/settings/sales_formula').success(function (formula) {
					$scope.formula = formula.value;
				});

				$scope.save = function () {
					$http.put('/config/settings/sales_categories', {
						name: 'sales_categories',
						value: JSON.stringify($scope.categories)
					}).success(function () {
						$http.put('/config/settings/sales_formula', {
							name: 'sales_formula',
							value: $scope.formula
						}).success(function () {
							$modalInstance.close();
						}).error(function () {
							alert('Could not save settings');
						});
					}).error(function () {
						alert('Could not save settings');
					});
				}
			}
		});
	}

	$scope.createReport = function () {
		scope = $rootScope.$new();
		scope.categories = $scope.categories;

		var modal = $modal.open({
			scope: scope,
			templateUrl: '/public/html/admin/reports.cash.new.html',
			controller: function ($scope, $modalInstance) {
				$scope.report = {
					date: new Date(),
					data: {}
				};
				$scope.createReport = function (report) {
					$modalInstance.close(report);
				}
			}
		});

		modal.close = function (report) {
			var data = [];
			for (var key in report.data) {
				if (!report.data.hasOwnProperty(key)) continue;

				data.push({
					date: report.date,
					category: key,
					amount: report.data[key]
				});
			}

			async.eachSeries(data, function (rep, cb) {
				$http.post('/reports/sales', rep).success(function () {
					cb();
				}).error(function () {
					alert('Failed to add report');
					cb('');
				});
			}, $scope.refreshData);
		};
	}
});