var app = angular.module('orderchef');

app.controller('OrderTypeCtrl', function ($scope, $http, $state, OrderType) {
  $scope.orderType = OrderType;

  $scope.save = function () {
    var p;
    if ($scope.orderType.id) {
      p = $http.put('/config/order-type/' + $scope.orderType.id, $scope.orderType);
    } else {
      p = $http.post('/config/order-types', $scope.orderType);
    }

    p.success(function () {
      $state.go('config.order_types');
    }).error(function () {
      alert('Cannot save order type');
    })
  }
});