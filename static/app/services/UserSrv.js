app.factory('UserSrv', function($http){
  var obj = {};

 
  obj.show_profile = function(){
    return $http.get('/users/profile')
  }

  obj.register = function(data){
    return $http.post('/users/register',data)
  }
  
  
  return obj;

})
