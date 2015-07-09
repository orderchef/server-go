var app = angular.module('orderchef', ['ui.router']);

function errorCb (cb) {
	return function () {
		cb(false, []);
	}
}

app.config(function ($httpProvider) {
	$httpProvider.interceptors.push(function ($q) {
		return {
			'request': function (config) {
				if (config.url.indexOf('/public') === -1) config.url = '/api' + config.url;

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
});