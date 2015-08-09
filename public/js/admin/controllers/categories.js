var app = angular.module('orderchef');

app.controller('CategoriesCtrl', function ($scope, $http, Categories) {
  $scope.refresh = function () {
    Categories.get().then(function (categories) {
      $scope.categories = categories;
    });
  }

  $scope.remove = function (obj) {
    $http.delete('/category/' + obj.id)
    .success(function () {
      $scope.refresh();
    }).error(function () {
      alert('Cannot remove category')
    });
  }

  $scope.refresh();
});

app.controller('CategoryCtrl', function ($scope, $http, $state, Category) {
  $scope.category = Category;

  $scope.refresh = function () {
    $http.get('/config/printers').success(function (printers) {
      $scope.availablePrinters = printers;

      $http.get('/category/' + Category.id + '/printers').success(function (printers) {
        $scope.printers = printers;
      });
    });
  }

  if (Category.id) $scope.refresh();

  $scope.addPrinter = function (printer) {
    if (!printer) return;
    $http.post('/category/' + $scope.category.id + '/printers/' + printer)
    .success(function () {
      $scope.refresh();
    })
  }

  $scope.removePrinter = function (printer) {
    $http.delete('/category/' + $scope.category.id + '/printers/' + printer).then($scope.refresh)
  }

  $scope.save = function () {
    var p;
    if ($scope.category.id) {
      p = $http.put('/category/' + $scope.category.id, $scope.category);
    } else {
      p = $http.post('/categories', $scope.category);
    }

    p.success(function () {
      $state.go('categories');
    }).error(function () {
      alert('Cannot save category');
    });
  }
});