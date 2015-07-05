// spec
function makeTest(name, url, action, data) {
  return {
    id: name,
    run: function ($http, spec, test, cb) {
      $http[action](url, data).success(function (data, status) {
        if (true || (!data && status != 204) || status >= 400) {
          return cb('Spec failed: status(' + status + '), data: \'' + data + '\'');
        }

        cb(null, data, status);
      }).error(function(data, status) {
        cb('Spec failed: status(' + status + '), data: \'' + data + '\'', data, status);
      });
    }
  };
}

function makeRecursiveTest(name, url, action, testSource, accessor) {
  return {
    id: 'recursive_' + name,
    run: function ($http, spec, test, cb) {
      var results = [];
      async.each(spec.data[testSource], function (d, callback) {
        makeTest(name, url.replace(':' + accessor, d[accessor]), action)
        .run($http, spec, test, function (err, data, status) {
          if (err) return callback(err);

          results.push(data);
          callback(null, data, status);
        });
      }, function (err) {
        cb(err, results, '?[multi]');
      });
    }
  }
}

function makeParameterTest(name, url, action, testSource, accessor) {
  return {
    id: 'param_' + name,
    run: function ($http, spec, test, cb) {
      makeTest(name, url.replace(':' + accessor, spec.data[testSource][accessor]), action)
      .run($http, spec, test, cb);
    }
  };
}

window.orderchef_specs = [
  {
    name: "Optional Test Name",
    tests: [
      makeTest('add', '/categories', 'post', {
        name: 'name..',
        description: 'desc'
      }),
      makeTest('getAll', '/categories', 'get'),
      makeRecursiveTest('get', '/category/:id', 'get', 'getAll', 'id'),
      makeParameterTest('delete_original', '/category/:id', 'delete', 'add', 'id')
    ]
  }
];