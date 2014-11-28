'use strict';

require('angular');
require('angular-route');
var controllers = require('./controllers/controllers');

var app = angular.module('MyApp', ['ngRoute']);

// Set up routes
app.config(['$routeProvider', function($routeProvider){
    $routeProvider.when('/', {
        templateUrl: 'views/landing.html',
        controller: 'landingController',
    });
}]);

// Add controller
app.controller('landingController', [
    '$scope', '$http',
    controllers.LandingController]);
