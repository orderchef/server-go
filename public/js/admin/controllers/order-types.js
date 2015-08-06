var app = angular.module('orderchef');

app.controller('OrderTypesCtrl', function ($scope, $http, OrderTypes) {
  $scope.refresh = function () {
    OrderTypes.get().then(function (types) {
      $scope.orderTypes = types;
    });
  }

  $scope.remove = function (obj) {
    $http.delete('/config/order-type/' + obj.id)
    .success(function () {
      $scope.refresh();
    }).error(function () {
      alert('Cannot remove order type')
    });
  }

  $scope.refresh();
});