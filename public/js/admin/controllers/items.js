var app = angular.module('orderchef');

app.controller('ItemsCtrl', function ($scope, $http) {
  $scope.refresh = function () {
    $http.get('/items').success(function (items) {
      $scope.items = items;
    });
  }

  $scope.remove = function (obj) {
    $http.delete('/item/' + obj.id)
    .success(function () {
      $scope.refresh();
    }).error(function () {
      alert('Cannot remove item')
    });
  }

  $scope.refresh();
});

app.controller('ItemCtrl', function ($scope, $http, $state, Item) {
  $scope.item = Item;

  $scope.save = function () {
    var p;
    if ($scope.item.id) {
      p = $http.put('/item/' + $scope.item.id, $scope.item);
    } else {
      p = $http.post('/items', $scope.item);
    }

    p.success(function () {
      $state.go('config.items');
    }).error(function () {
      alert('Cannot save item');
    })
  }
});