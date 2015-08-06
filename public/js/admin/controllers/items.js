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
      alert('Cannot remove item');
    });
  }

  $scope.refresh();
});

app.controller('ItemCtrl', function ($scope, $http, $state, Item, Categories, Modifiers) {
  $scope.item = Item;
  Categories.get().then(function (cats) {
    $scope.categories = cats;
  });

  $scope.refresh = function () {
    Modifiers.get().then(function (mods) {
      $scope.allModifiers = mods;
    });

    $http.get('/item/' + $scope.item.id + '/modifiers')
    .success(function (modifiers) {
      $scope.modifiers = modifiers;

      for (var i = 0; i < $scope.modifiers.length; i++) {
        for (var ii = 0; ii < $scope.allModifiers.length; ii++) {
          if ($scope.allModifiers[ii].id == $scope.modifiers[i]) {
            $scope.modifiers[i] = $scope.allModifiers[ii];
            $scope.allModifiers.splice(ii, 1);
            break;
          }
        }
      }
    });
  }
  if ($scope.item.id) $scope.refresh();

  $scope.addModifier = function (mod) {
    if (!mod) return;

    $scope.modifiers.push(mod);
    $scope.selectedModifier = null;

    for (var i = 0; i < $scope.allModifiers.length; i++) {
      if ($scope.allModifiers[i].id == mod.id) {
        $scope.allModifiers.splice(i, 1);
        break;
      }
    }

    $http.post('/item/' + $scope.item.id + '/modifiers', {
      modifier_group_id: mod.id
    });
  }

  $scope.removeModifier = function (mod) {
    $http.delete('/item/' + $scope.item.id + '/modifier/' + mod.id);

    $scope.allModifiers.push(mod);
    
    for (var i = 0; i < $scope.modifiers.length; i++) {
      if ($scope.modifiers[i].id == mod.id) {
        $scope.modifiers.splice(i, 1);
        break;
      }
    }
  }

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
    });
  }
});