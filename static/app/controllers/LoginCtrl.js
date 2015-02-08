app.controller('LoginCtrl', function($scope, $location, $rootScope, AuthSrv){
   $scope.user = {}
   $scope.user.email = "";
   $scope.user.password = "";


   $scope.login = function (user) {
       AuthSrv.login({email: user.email, password: user.password}).
            success(function () {
                $rootScope.$broadcast("successful:login");
                $location.path('/');
            })
            .error(function(){
                alert('Wrong email and password');
            });


    };

});
