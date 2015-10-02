'use strict'

angular.module('gameit', []) 
	.config(function($routeProvider)) {
		$routeProvider
			.when('/', {
				templateUrl: 'components/lol/home/home.html',
				controller: 'homeCtrl'
			})
			// .when('/tournament', {
			// 	templateUrl: 'components/lol/tournament/tournament.html'
			// 	controller: 'tournamentCtrl'
			// })
			.otherwise({
				redirectTo: '/'
			});
	}