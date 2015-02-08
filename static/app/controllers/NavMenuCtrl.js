app.controller("NavMenuCtrl", function($scope, $http) {
  var getNavigation;
  $scope.sitemap = [];
  getNavigation = function() {
    return $http.get('/sitemap').success(function(data) {
      return $scope.sitemap = data;
    });
  };
  getNavigation();
  $scope.$on("successful:login", getNavigation);
  $scope.$on("successful:logout", getNavigation);

});
