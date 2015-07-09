var app = angular.module('orderchef');

app.controller('TablesCtrl', function ($scope, $http, TableTypes) {
  $scope.refresh = function () {
    TableTypes.get().then(function (types) {
      $scope.tableTypes = types;

      $http.get('/tables').success(function (data) {
        for (var i = 0; i < data.length; i++) {
          for (var x = 0; x < types.length; x++) {
            if (types[x].id == data[i].type_id) {
              data[i].table_type = types[x];
              break;
            }
          }
        }

        $scope.tables = data;
      });
    });
  }

  $scope.deleteTable = function (table) {
    $http.delete('/table/' + table.id)
    .success(function () {
      $scope.refresh();
    }).error(function () {
      alert('Cannot remove table')
    });
  }

  $scope.refresh();
});