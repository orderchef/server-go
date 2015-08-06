var app = angular.module('orderchef');

app.service('Categories', function ($q, $http) {
  this.get = function () {
    var p = $q.defer();

    $http.get('/categories')
    .success(function(data) {
      p.resolve(data);
    }).error(function() {
      p.reject();
    });

    return p.promise;
  }
})