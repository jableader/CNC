(function () {

    angular.module('baseApp')
        .controller('mainController', ['$mdSidenav', '$rootScope', '$location', mainController]);

    function mainController($mdSidenav, $rootScope, $location) {
        var self = this;

        self.isRoutingMode = false;
        self.focus = 'summary';

        self.activateRouted = function(route, element) {
          console.log("Now activating section: " + element + " on route: " + route);
          self.focus = element;
          $rootScope.$broadcast('component.changed', element);
          $mdSidenav('left').close()
          $location.path(route);
          self.isRoutingMode = true;
        };

        self.activate = function (element) {
          console.log("Now activating section: " + element);
          self.focus = element;
          self.isRoutingMode = false;
          $rootScope.$broadcast('component.changed', element);
          $mdSidenav('left').close()
          $location.path("/");
        };

        self.logout = function() {
          window.location.href = '/logout';
        };

        self.toggle = function () {
            $mdSidenav('left').toggle();
        };
    }
})();
