var app = angular.module('orderchef', ['ui.router', 'ui.bootstrap']);

app.config(function ($httpProvider) {
	$httpProvider.interceptors.push(function ($q) {
		return {
			'request': function (config) {
				if (!(config.url.indexOf('/public') != -1 || config.url.indexOf('template/') != -1)) config.url = '/api' + config.url;

				return config || $q.when(config);
			}
		}
	});
});

app.config(function($stateProvider, $urlRouterProvider, $locationProvider) {
  var base = '/public/html/admin';

	$locationProvider.html5Mode({
		enabled: false,
		requireBase: true
	});

	$urlRouterProvider.otherwise('/tables');

	$stateProvider
	.state('tables', {
		url: '/tables',
		templateUrl: base + '/tables.html',
		controller: 'TablesCtrl'
	})
  .state('table', {
    url: '/table/:table_id',
    abstract: true,
    template: '<ui-view></ui-view>',
    resolve: {
      Table: function ($q, $http, $stateParams) {
        if ($stateParams.table_id == 'new') return {};

        var d = $q.defer();

        $http.get('/table/' + $stateParams.table_id).success(function (data) {
          d.resolve(data);
        });

        return d.promise;
      }
    }
  })
  .state('table.edit', {
    url: '/edit',
    templateUrl: base + '/table_edit.html',
    controller: 'TableEditCtrl'
  })

	.state('config', {
		abstract: true,
		url: '/config',
		template: '<ui-view></ui-view>'
	})
	.state('config.table_types', {
		url: '/table-types',
		templateUrl: base + '/config.table_types.html',
		controller: 'TableTypesCtrl'
	})
	.state('config.table_type', {
		url: '/table-type/:table_type_id',
		templateUrl: base + '/config.table_type.html',
		controller: 'TableTypeCtrl',
		resolve: {
      TableType: function ($q, $http, $stateParams) {
        if ($stateParams.table_type_id == 'new') return {};

        var d = $q.defer();

        $http.get('/config/table-type/' + $stateParams.table_type_id).success(function (data) {
          d.resolve(data);
        });

        return d.promise;
      }
    }
	})
	.state('config.order_types', {
		url: '/order-types',
		templateUrl: base + '/config.order_types.html',
		controller: 'OrderTypesCtrl'
	})
	.state('config.order_type', {
		url: '/order-type/:order_type_id',
		templateUrl: base + '/config.order_type.html',
		controller: 'OrderTypeCtrl',
		resolve: {
      OrderType: function ($q, $http, $stateParams) {
        if ($stateParams.order_type_id == 'new') return {};

        var d = $q.defer();

        $http.get('/config/order-type/' + $stateParams.order_type_id).success(function (data) {
          d.resolve(data);
        });

        return d.promise;
      }
    }
	})
	.state('config.modifiers', {
		url: '/modifiers',
		templateUrl: base + '/config.modifiers.html',
		controller: 'ModifiersCtrl'
	})
	.state('config.modifier', {
		url: '/modifier/:modifier_id',
		templateUrl: base + '/config.modifier.html',
		controller: 'ModifierCtrl',
		resolve: {
      Modifier: function ($q, $http, $stateParams) {
        if ($stateParams.modifier_id == 'new') return {};

        var d = $q.defer();

        $http.get('/config/modifier/' + $stateParams.modifier_id).success(function (data) {
          d.resolve(data);
        });

        return d.promise;
      }
    }
	})
	.state('config.items', {
		url: '/items',
		templateUrl: base + '/config.items.html',
		controller: 'ItemsCtrl'
	})
	.state('config.item', {
		url: '/item/:item_id',
		templateUrl: base + '/config.item.html',
		controller: 'ItemCtrl',
		resolve: {
			Item: function ($q, $http, $stateParams) {
				if ($stateParams.item_id == 'new') return {};

				var d = $q.defer();

				$http.get('/item/' + $stateParams.item_id).success(function (data) {
					d.resolve(data);
				});

				return d.promise;
			}
		}
	})
	.state('config.bills', {
		url: '/bills',
		templateUrl: base + '/config.bills.html',
		controller: 'ConfigBillsCtrl'
	})

	.state('categories', {
		url: '/categories',
		templateUrl: base + '/categories.html',
		controller: 'CategoriesCtrl'
	})
	.state('category', {
		url: '/category/:category_id',
		templateUrl: base + '/category.html',
		controller: 'CategoryCtrl',
		resolve: {
      Category: function ($q, $http, $stateParams) {
        if ($stateParams.category_id == 'new') return {};

        var d = $q.defer();

        $http.get('/category/' + $stateParams.category_id).success(function (data) {
          d.resolve(data);
        });

        return d.promise;
      }
    }
	})

	.state('reports', {
		url: '/reports',
		abstract: true,
		template: '<ui-view></ui-view>'
	})
	.state('reports.bills', {
		url: '/bills',
		abstract: true,
		template: '<ui-view></ui-view>'
	})
	.state('reports.bills.bills', {
		url: '',
		templateUrl: base + '/reports.bills.html',
		controller: 'ReportBillsCtrl'
	})
	.state('reports.bills.extras', {
		url: '/extras',
		templateUrl: base + '/reports.bills.extras.html',
		controller: 'BillExtrasCtrl'
	})
	.state('reports.cash', {
		url: '/cash',
		templateUrl: base + '/reports.cash.html',
		controller: 'ReportCashCtrl'
	})
	.state('reports.popularItems', {
		url: '/popularItems',
		templateUrl: base + '/reports.popularItems.html',
		controller: 'PopularItemsReportCtrl'
	})
	.state('reports.popularItemsCategories', {
		url: '/popularItemsCategories',
		templateUrl: base + '/reports.popularItems.categories.html',
		controller: 'PopularItemsReportCategoriesCtrl'
	})
	.state('reports.sales', {
		url: '/sales',
		templateUrl: base + '/reports.sales.html',
		controller: 'SalesReportCtrl'
	});
});