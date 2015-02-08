app.factory('AuthSrv', function($http){
  var obj = {};

  obj.login = function(data){
    console.log(data)
    return $http.post('/auth/login', data)
  }


  obj.register = function(data){
    console.log(data)
    return $http.post('/auth/register', data)
  }


 
  obj.auth_user = function(){
    return $http.get('/auth/user')
  }

  return obj;

})
