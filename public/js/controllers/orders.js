var app = angular.module('orderchef');

app.controller('OrdersCtrl', function ($scope, $http, OrderGroup, OrderTypes, Items, Modifiers) {
  $scope.group = OrderGroup;

  $scope.refresh = function () {
    async.waterfall([
      function(cb) {
        Items.get().then(function (items) {
          $scope.items = items;
          cb();
        });
      },
      function(cb) {
        OrderTypes.get().then(function (types) {
          $scope.orderTypes = types;
          cb();
        });
      },
      function(cb) {
        Modifiers.get().then(function (modifiers) {
          $scope.modifiers = modifiers;

          async.each($scope.modifiers, function (mod, cb) {
            $http.get('/config/modifier/' + mod.id + '/items')
            .success(function (items) {
              mod.items = items;
              cb();
            }).error(function () {
              cb();
            });
          }, cb);
        });
      },
      function (cb) {
        $http.get('/order-group/' + OrderGroup.id + '/orders').success(function (orders) {
          $scope.orderItems = orders;
          cb();
        });
      },
      function (cb) {
        for (var i = 0; i < $scope.orderItems.length; i++) {
          for (var ii = 0; ii < $scope.orderTypes.length; ii++) {
            if ($scope.orderItems[i].type_id == $scope.orderTypes[ii].id) {
              $scope.orderItems[i].type = $scope.orderTypes[ii];
              break;
            }
          }
        }

        async.each($scope.orderItems, function (orderItem, cb) {
          $http.get('/order/' + orderItem.id + '/items').success(function (items) {
            orderItem.items = items;

            for (var i = 0; i < items.length; i++) {
              for (var ii = 0; ii < $scope.items.length; ii++) {
                if (items[i].item_id == $scope.items[ii].id) {
                  items[i].item = $scope.items[ii];
                  break;
                }
              }
            }

            async.each(orderItem.items, function (item, cb) {
              $http.get('/order/' + orderItem.id + '/item/' + item.id + '/modifiers')
              .success(function (data) {
                item.modifiers = data;
                for (var i = 0; i < item.modifiers.length; i++) {
                  var m = item.modifiers[i];

                  for (var ii = 0; ii < $scope.modifiers.length; ii++) {
                    var group = $scope.modifiers[ii];
                    if (m.modifier_group_id != group.id) continue;

                    for (var mi = 0; mi < group.items.length; mi++) {
                      if (m.modifier_id == group.items[mi].id) {
                        m.modifier = group.items[mi];
                        m.modifier_group = group;

                        break;
                      }
                    }
                  }
                }
                cb();
              }).error(function () {
                cb();
              });
            }, cb);
          }).error(function () {
            cb();
          }, cb);
        });
      }
    ]);
  }

  $scope.refresh();

  $http.get('/table/' + OrderGroup.table_id).success(function (table) {
    $scope.table = table;
  });
});