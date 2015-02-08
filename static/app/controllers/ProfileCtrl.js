app.controller('ProfileCtrl', function($scope, $routeParams, $location ,UserSrv){
  $scope.user = {}
  UserSrv.show_profile()
    .success(function(data){
      $scope.user = data;
      
    })
    .error(function(){
       $location.path('/');
    })



})
