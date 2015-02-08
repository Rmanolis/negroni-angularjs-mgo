//loading Lodash library in angular
var lodash = angular.module('lodash', []);
lodash.factory('_', function () {
    return window._; // assumes underscore has already been loaded on the page
});


var app = angular.module('App', ['ngRoute', 
    'ui.bootstrap']);

app.run(function($rootScope,$location,AuthSrv){
    // enumerate routes that don't need authentication
    var routesThatDontRequireAuth = ['/login','/register'];

    // check if current location matches route
    var routeClean = function (route) {
        return _.find(routesThatDontRequireAuth,
            function (noAuthRoute) {
                return route === noAuthRoute;
            });
    };

    $rootScope.$on('$routeChangeStart', function (event, next, current) {
        // if route requires auth and user is not logged in
        AuthSrv.auth_user().error(function(){
            if (!routeClean($location.url()) ) {
                // redirect back to login
                $location.path('/login');
            }
        }).success(function(){
                  })

    });
});


app.config(function ($routeProvider) {
    $routeProvider
        .when('/', {
            templateUrl: '/static/app/pages/home.html',
            controller: 'HomeCtrl'
        })

        .when('/login', {
            templateUrl: '/static/app/pages/login.html',
            controller: 'LoginCtrl'
        })

        .when('/register', {
          templateUrl: '/static/app/pages/register.html',
          controller: 'RegisterCtrl'
        })
 
        .when('/users/profile',{
          templateUrl: '/static/app/pages/profile.html',
          controller: 'ProfileCtrl'
        })


        
});

