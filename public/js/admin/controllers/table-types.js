var app = angular.module('orderchef');

app.controller('TableTypesCtrl', function ($scope, $http, TableTypes) {
  $scope.refresh = function () {
    TableTypes.get().then(function (types) {
      $scope.tableTypes = types;
    });
  }

  $scope.remove = function (obj) {
    $http.delete('/config/table-type/' + obj.id)
    .success(function () {
      $scope.refresh();
    }).error(function () {
      alert('Cannot remove table type')
    });
  }

  $scope.refresh();
});