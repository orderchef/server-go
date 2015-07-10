var app = angular.module('orderchef');

app.controller('ModifiersCtrl', function ($scope, $http, Modifiers) {
  $scope.refresh = function () {
    Modifiers.get().then(function (modifiers) {
      $scope.modifiers = modifiers;
    });
  }

  $scope.remove = function (obj) {
    $http.delete('/config/modifier/' + obj.id)
    .success(function () {
      $scope.refresh();
    }).error(function () {
      alert('Cannot remove modifier')
    });
  }

  $scope.refresh();
});

app.controller('ModifierCtrl', function ($scope, $http, $state, Modifier) {
  $scope.modifier = Modifier;

  $scope.refresh = function () {
    $http.get('/config/modifier/' + $scope.modifier.id + '/items')
    .success(function (modifiers) {
      $scope.modifiers = modifiers;
    });
  }

  if (Modifier.id) $scope.refresh();

  $scope.save = function () {
    var p;
    if ($scope.modifier.id) {
      p = $http.put('/config/modifier/' + $scope.modifier.id, $scope.modifier);
    } else {
      p = $http.post('/config/modifiers', $scope.modifier);
    }

    p.success(function () {
      $state.go('config.modifiers');
    }).error(function () {
      alert('Cannot save modifier');
    })
  }

  $scope.saveModifier = function (modifier) {
    if (!$scope.modifier.id) return alert('Save group first');

    var p;
    if (modifier.id) {
      p = $http.put('/config/modifier/' + $scope.modifier.id + '/item/' + modifier.id, modifier)
    } else {
      p = $http.post('/config/modifier/' + $scope.modifier.id + '/items', modifier)
    }

    p.success(function (data) {
      $scope.refresh();
      modifier.id = data.id;
    }).error(function () {
      alert('cannot save modifier');
    });
  }

  $scope.removeModifier = function (modifier) {
    if (!$scope.modifier.id) return alert('Save group first');

    $http.delete('/config/modifier/' + $scope.modifier.id + '/item/' + modifier.id)
    .success(function () {
      $scope.refresh();
    }).error(function () {
      alert('cannot remove modifier');
    })
  }
});