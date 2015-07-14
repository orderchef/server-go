var app = angular.module('orderchef', ['ui.router']);

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
  var base = '/public/html';

	$locationProvider.html5Mode({
		enabled: false,
		requireBase: true
	});

	$urlRouterProvider.otherwise('/');

	$stateProvider
	.state('tables', {
		url: '/',
		templateUrl: base + '/tables.html',
		controller: 'TablesCtrl'
	})
	.state('orders', {
		url: '/table/:table_id',
		templateUrl: base + '/orders.html',
		controller: 'OrdersCtrl',
		resolve: {
			OrderGroup: function ($q, $http, $stateParams) {
				var d = $q.defer();

				$http.get('/table/' + $stateParams.table_id + '/group')
				.success(function (group) {
					d.resolve(group);
				}).error(function () {
					d.reject();
				});

				return d.promise;
			}
		}
	});
});