var app = angular.module('orderchef');

app.service('TableTypes', function ($q, $http) {
  this.get = function () {
    var p = $q.defer();

    $http.get('/config/table-types')
    .success(function(data) {
      p.resolve(data);
    }).error(function() {
      p.reject();
    });

    return p.promise;
  }
})