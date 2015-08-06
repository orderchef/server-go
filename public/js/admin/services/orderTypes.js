var app = angular.module('orderchef');

app.service('OrderTypes', function ($q, $http) {
  this.get = function () {
    var p = $q.defer();

    $http.get('/config/order-types')
    .success(function(data) {
      p.resolve(data);
    }).error(function() {
      p.reject();
    });

    return p.promise;
  }
})