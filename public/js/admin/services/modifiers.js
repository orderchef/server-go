var app = angular.module('orderchef');

app.service('Modifiers', function ($q, $http) {
  this.get = function () {
    var p = $q.defer();

    $http.get('/config/modifiers')
    .success(function(data) {
      p.resolve(data);
    }).error(function() {
      p.reject();
    });

    return p.promise;
  }
})