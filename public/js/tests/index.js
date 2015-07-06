angular.module('orderchef')
.controller('TestsCtrl', ['$scope', '$http', 'TestService', function ($scope, $http, TestService) {
  $scope.specs = window.orderchef_specs;

  async.each($scope.specs, function (spec, cb) {
    spec.data = {};
    spec.dataStringified = {};

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
      test.run($http, spec, test, function (data, status, url) {
        spec.data[test.id] = data;
        spec.dataStringified[test.id] = JSON.stringify(data, null, 2);

        test.lag = (Date.now() - start) + 'ms';
        test.responseCode = status;
        test.dataType = typeof spec.dataStringified[test.id];
        test.state = 0;

        if (test.expect && ((typeof test.expect !== 'function' && status != test.expect) || (typeof test.expect == 'function' && test.expect(spec, test, data, status) !== true))) {
          var orig = spec.dataStringified[test.id];

          test.state = 1;
          test.dataType = 'string';

          spec.dataStringified[test.id] = status + ' ' + url + ' failed. Expectation ';
          if (typeof test.expect !== 'function')
            spec.dataStringified[test.id] += '(HTTP ' + test.expect + ')';
          else if (typeof test.expect == 'function')
            spec.dataStringified[test.id] += '(func: ' + test.expect(spec, test, data, status) + ')';

          spec.dataStringified[test.id] += '\n\n' + orig;
        } else if (!test.expect && ((!data && status != 204) || status >= 400)) {
          test.state = 1;
          test.dataType = 'string';

          spec.dataStringified[test.id] = status + ' ' + url + ' failed. \n\n' + spec.dataStringified[test.id];
        }

        cb();
      });
    });
  }
}]);