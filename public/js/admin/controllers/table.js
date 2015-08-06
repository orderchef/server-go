var app = angular.module('orderchef');

app.controller('TableEditCtrl', function ($scope, $http, $state, Table, TableTypes) {
  TableTypes.get().then(function (types) {
    $scope.tableTypes = types;
  });

  $scope.table = Table;

  $scope.saveTable = function () {
    var p;
    if ($scope.table.id) {
      p = $http.put('/table/' + $scope.table.id, $scope.table);
    } else {
      p = $http.post('/tables', $scope.table);
    }

    p.success(function () {
      $state.go('tables');
    }).error(function () {
      alert('Cannot save table');
    })
  }
});