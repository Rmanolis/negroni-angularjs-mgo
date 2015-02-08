app.controller('RegisterCtrl', function($scope, $routeParams,
      $location,AuthSrv){
  $scope.user = {
    name : "",
    email : "",
    password : "" };

  $scope.save = function(user){
    AuthSrv.register(user)
      .success(function(data){
        if(data.error){
          alert(data.error)
        }else{
          $location.path('/')
        }
      });

  };

  $scope.cancel = function(){
    $location.path('/')
  };

})
