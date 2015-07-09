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

  $scope.runAll = function (spec, done) {
    async.eachSeries($scope.specs, function (spec, cb) {
      $scope.runSpec(spec, cb);
    });
  }

  $scope.runSpec = function (spec, done) {
    if (typeof done != 'function') done = function(){};

    spec.running = true;
    spec.testsDone = 0;
    spec.testsErr = false;

    var runner = $scope.runTest(spec);
    var callback = function (err) {
      if (spec.post && !err) {
        return async.each(spec.post, runner, done);
      }

      done(err);
    }

    if (spec.pre) {
      async.each(spec.pre, runner, function (err) {
        if (err) throw err;

        async.eachSeries(spec.tests, runner, callback);
      });

      return;
    }

    async.eachSeries(spec.tests, runner, callback);
  }

  $scope.runTest = function (spec) {
    return function (test, cb) {
      var start = Date.now();
      test.run($http, spec, test, function (data, status, url) {
        spec.data[test.id] = data;
        spec.dataStringified[test.id] = JSON.stringify(data, null, 2);

        test.lag = (Date.now() - start) + 'ms';
        test.responseCode = status;
        test.dataType = typeof spec.dataStringified[test.id];
        test.state = 0;
        spec.testsDone++;

        if (test.expect && ((typeof test.expect !== 'function' && status != test.expect) || (typeof test.expect == 'function' && test.expect(spec, test, data, status) !== true))) {
          var orig = spec.dataStringified[test.id];

          spec.testsErr = true;
          test.state = 1;
          test.dataType = 'string';

          spec.dataStringified[test.id] = status + ' ' + url + ' failed. Expectation ';
          if (typeof test.expect !== 'function')
            spec.dataStringified[test.id] += '(HTTP ' + test.expect + ')';
          else if (typeof test.expect == 'function')
            spec.dataStringified[test.id] += '(func: ' + test.expect(spec, test, data, status) + ')';

          spec.dataStringified[test.id] += '\n\n' + orig;
        } else if (!test.expect && ((!data && status != 204) || status >= 400)) {
          spec.testsErr = true;
          test.state = 1;
          test.dataType = 'string';

          spec.dataStringified[test.id] = status + ' ' + url + ' failed. \n\n' + spec.dataStringified[test.id];
        }

        cb();
      });
    }
  }
}]);