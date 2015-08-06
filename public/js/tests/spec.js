// spec
function makeTest(name, url, action, data, expect) {
  return {
    id: name,
    run: function ($http, spec, test, cb) {
      if (typeof data == 'function') {
        data = data(spec, test, url);
      }

      $http[action](url, data).success(function (data, status) {
        cb(data, status, url);
      }).error(function(data, status) {
        cb(data, status, url);
      });
    },
    expect: expect
  };
}

function makeRecursiveTest(name, url, action, testSource, accessor, expect) {
  return {
    id: name,
    run: function ($http, spec, test, cb) {
      var results = [];
      var allStatus = null;
      async.each(spec.data[testSource], function (d, callback) {
        makeTest(name, url.replace(':' + accessor, d[accessor]), action, null, expect)
        .run($http, spec, test, function (data, status, url) {
          if (allStatus == null) allStatus = status;
          if (allStatus != status) allStatus = 'multi';

          results.push(data);
          callback(null, data, status);
        });
      }, function (err) {
        cb(results, allStatus, url);
      });
    },
    expect: expect
  }
}

function makeParameterTest(name, url, action, data, testSource, accessor, expect) {
  return {
    id: name,
    run: function ($http, spec, test, cb) {
      makeTest(name, url.replace(':' + accessor, spec.data[testSource][accessor]), action, data, expect)
      .run($http, spec, test, cb);
    },
    expect: expect
  };
}

window.orderchef_specs = [];

// Categories
window.orderchef_specs.push({
  name: 'Categories',
  tests: [
    makeTest('add', '/categories', 'post', {
      name: 'name..',
      description: 'desc'
    }, 201),
    makeTest('get_all', '/categories', 'get', null, 200),
    makeRecursiveTest('get', '/category/:id', 'get', 'get_all', 'id', 200),
    makeParameterTest('delete_original', '/category/:id', 'delete', null, 'add', 'id', 204)
  ]
});

// Modifiers
window.orderchef_specs.push({
  name: 'Config Modifiers',
  tests: [
    makeTest('add', '/config/modifiers', 'post', {
      name: 'Modifier group',
      required: false
    }, 201),
    makeTest('get_all', '/config/modifiers', 'get', null, 200),
    makeRecursiveTest('get', '/config/modifier/:id', 'get', 'get_all', 'id', 200),
    makeParameterTest('add_modifier_to_group', '/config/modifier/:id/items', 'post', {
      name: 'Medium',
      price: 0
    }, 'add', 'id', 201),
    makeParameterTest('get_all_modifiers', '/config/modifier/:id/items', 'get', null, 'add', 'id', 200),
    {
      id: 'get_single_modifier',
      run: function ($http, spec, test, cb) {
        var url = '/config/modifier/:modifier_id/item/:item_id'.replace(':modifier_id', spec.data.add.id).replace(':item_id', spec.data.add_modifier_to_group.id);

        makeTest('get_single_modifier', url, 'get').run($http, spec, test, cb);
      },
      expect: 200
    },
    {
      id: 'update_single_modifier',
      run: function ($http, spec, test, cb) {
        var url = '/config/modifier/:modifier_id/item/:item_id'.replace(':modifier_id', spec.data.add.id).replace(':item_id', spec.data.add_modifier_to_group.id);

        var o = JSON.parse(JSON.stringify(spec.data.add_modifier_to_group));
        o.name = 'Large';
        o.price = 2.99;
        makeTest('update_single_modifier', url, 'put', o).run($http, spec, test, cb);
      },
      expect: 201
    },
    {
      id: 'get_updated_modifier',
      run: function ($http, spec, test, cb) {
        var url = '/config/modifier/:modifier_id/item/:item_id'.replace(':modifier_id', spec.data.add.id).replace(':item_id', spec.data.add_modifier_to_group.id);

        makeTest('get_updated_modifier', url, 'get').run($http, spec, test, cb);
      },
      expect: 200
    },
    {
      id: 'remove_modifier',
      run: function ($http, spec, test, cb) {
        var url = '/config/modifier/:modifier_id/item/:item_id'.replace(':modifier_id', spec.data.add.id).replace(':item_id', spec.data.add_modifier_to_group.id);

        makeTest('remove_modifier', url, 'delete').run($http, spec, test, cb);
      },
      expect: 204
    },
    makeParameterTest('get_all_modifiers_test', '/config/modifier/:id/items', 'get', null, 'add', 'id', function (spec, test, data, status) {
      if (status != 200) return 'expecting HTTP 202';
      if (data.length != spec.data.get_all_modifiers.length - 1) {
        return 'modifier not deleted';
      }

      return true
    }),
    makeParameterTest('delete_modifier', '/config/modifier/:id', 'delete', null, 'add', 'id', 204)
  ]
});

// Order Types
window.orderchef_specs.push({
  name: 'Config Order Types',
  tests: [
    makeTest('add', '/config/order-types', 'post', {
      name: 'Order Type',
      description: 'hmmm'
    }, 201),
    makeTest('get_all', '/config/order-types', 'get', null, 200),
    makeRecursiveTest('get', '/config/order-type/:id', 'get', 'get_all', 'id', 200),
    makeParameterTest('delete_original', '/config/order-type/:id', 'delete', null, 'add', 'id', 204)
  ]
});

// Table Types
window.orderchef_specs.push({
  name: 'Config Table Types',
  tests: [
    makeTest('add', '/config/table-types', 'post', {
      name: 'Table Type'
    }, 201),
    makeTest('get_all', '/config/table-types', 'get', null, 200),
    makeRecursiveTest('get', '/config/table-type/:id', 'get', 'get_all', 'id', 200),
    makeParameterTest('delete_original', '/config/table-type/:id', 'delete', null, 'add', 'id', 204)
  ]
});

// Items
window.orderchef_specs.push({
  name: 'Items',
  pre: [
    makeTest('add_category', '/categories', 'post', {
      name: 'name..',
      description: 'desc'
    }, 201),
  ],
  tests: [
    makeTest('add', '/items', 'post', function (spec, test, url) {
      return {
        name: 'Coca Cola',
        description: 'Glass 0.25l',
        price: 2.0,
        category_id: spec.data.add_category.id
      }
    }, 201),
    makeTest('get_all', '/items', 'get', null, 200),
    makeRecursiveTest('get', '/item/:id', 'get', 'get_all', 'id', 200),
    makeParameterTest('delete_original', '/item/:id', 'delete', null, 'add', 'id', 204)
  ],
  post: [
    makeParameterTest('delete_category', '/category/:id', 'delete', null, 'add_category', 'id', 204)
  ]
});

window.orderchef_specs.push({
  name: 'Tables',
  pre: [
    makeTest('add_table_type', '/config/table-types', 'post', {
      name: 'Table Type'
    }, 201)
  ],
  tests: [
    makeTest('add', '/tables', 'post', function (spec, test, url) {
      return {
        type_id: spec.data.add_table_type.id,
        name: 'Test Table',
        table_number: 'two',
        location: 'Internets'
      }
    }, 201),
    makeTest('get_all', '/tables', 'get', null, 200),
    makeRecursiveTest('get', '/table/:id', 'get', 'get_all', 'id', 200),
    makeParameterTest('delete_original', '/table/:id', 'delete', null, 'add', 'id', 204)
  ],
  post: [
    makeParameterTest('delete_table_type', '/config/table-type/:id', 'delete', null, 'add_table_type', 'id', 204)
  ]
});