var app = angular.module('orderchef');

app.controller('TableTypeCtrl', function ($scope, $http, $state, TableType) {
  $scope.tableType = TableType;

  $scope.save = function () {
    var p;
    if ($scope.tableType.id) {
      p = $http.put('/config/table-type/' + $scope.tableType.id, $scope.tableType);
    } else {
      p = $http.post('/config/table-types', $scope.tableType);
    }

    p.success(function () {
      $state.go('config.table_types');
    }).error(function () {
      alert('Cannot save table type');
    })
  }
});