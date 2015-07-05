angular.module('orderchef')
.controller('TestsCtrl', ['$scope', '$http', 'TestService', function ($scope, $http, TestService) {
  $scope.specs = window.orderchef_specs;

  async.each($scope.specs, function (spec, cb) {
    spec.data = {};
    for (var i = 0; i < spec.tests.length; i++) {
      var test = spec.tests[i];
      test.name = ' .' + test.id;
      test.state = -1;
      test.lag = 'n/a';
      spec.data[test.id] = []
    }
  });

  $scope.runSpec = function (spec) {
    async.eachSeries(spec.tests, function (test, cb) {
      var start = Date.now();
      test.run($http, spec, test, function (err, data, status) {
        spec.data[test.id] = err || data;

        test.lag = (Date.now() - start) + 'ms';
        test.responseCode = status;
        test.dataType = typeof spec.data[test.id];
        test.state = 0;

        if (Object.prototype.toString.call(data) == '[object Array]') {
          test.dataType = 'array';
        }

        if (err) {
          test.state = 1;
        }

        cb();
      });
    });
  }
}]);