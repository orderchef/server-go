var app = angular.module('orderchef');

app.controller('TablesCtrl', function ($scope, $http, TableTypes) {
  $scope.refresh = function () {
    $http.get('/tables/open')
    .success(function(data) {
      $scope.openTables = data;

      for (var i = 0; i < data.length; i++) {
        for (var x = 0; x < $scope.tableTypes.length; x++) {
          if ($scope.tableTypes[x].id == data[i].type_id) {
            data[i].table_type = $scope.tableTypes[x];
            break;
          }
        }
      }
    })
  }

  TableTypes.get().then(function (types) {
    $scope.tableTypes = types;

    $http.get('/tables').success(function (data) {
      for (var i = 0; i < data.length; i++) {
        for (var x = 0; x < $scope.tableTypes.length; x++) {
          if ($scope.tableTypes[x].id == data[i].type_id) {
            data[i].table_type = $scope.tableTypes[x];
            if (!$scope.tableTypes[x].tables) {
              $scope.tableTypes[x].tables = [];
            }

            $scope.tableTypes[x].tables.push(data[i]);
          }
        }
      }

      $scope.tables = data;
    });

    $scope.refresh();
  });

  $scope.openTable = function (table) {
    $http.get('/table/' + table.id + '/group').success(function () {
      $scope.refresh();
    });
  }
});