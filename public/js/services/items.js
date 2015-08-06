var app = angular.module('orderchef');

app.service('Items', function ($q, $http) {
  this.get = function () {
    var p = $q.defer();

    $http.get('/items')
    .success(function(data) {
      p.resolve(data);
    }).error(function() {
      p.reject();
    });

    return p.promise;
  }
})