(function () {

    angular.module('baseApp')
        .controller('pluginListController', ['$scope', '$rootScope', '$http', '$mdDialog', pluginListController]);

    function pluginListController($scope, $rootScope, $http, $mdDialog) {
        var self = this;
        $scope.plugins = [];
        $scope.showLoading = false;

        self.refresh = function() {
          console.log("pluginListController.refresh()")
          $scope.showLoading = true;
          $http.get('/plugins', {}).then(function (response) { //get entities data to display in table
            plugins = response.data;
            $scope.running = plugins.Running;
            $scope.stopped = plugins.Disabled;
            $scope.plugins = [];
            for(var i = 0; i < $scope.running.length ;i++)
              $scope.plugins[$scope.plugins.length] = $scope.running[i];
            for(var i = 0; i < $scope.stopped.length ;i++)
              $scope.plugins[$scope.plugins.length] = $scope.stopped[i];
            $scope.showLoading = false;
            console.log($scope.plugins);
          });
        }

        $scope.refresh = self.refresh;
        self.refresh();
    };

})();
