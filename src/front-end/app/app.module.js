(function() {
	'use strict';
	angular.module('app', ['ngRoute', 'app.home', 'app.summoner']);
		

	angular.module('app')
			.config(function($routeProvider) {
				$routeProvider
					.when('/home', {
						templateUrl: 'app/home/home.html',
						controller: 'HomeCtrl'
					})
					.when('/summoner', {
						templateUrl: 'app/summoner/summoner.html',
						controller: 'SummonerCtrl'
					})
					// .when('/tournament', {
					// 	templateUrl: 'components/lol/tournament/tournament.html',
					// 	controller: 'tournamentCtrl'
					// })
					// .otherwise({
					// 	redirectTo: '/test'
					// })
			});

	angular.module('app').controller('MainCtrl', ['$scope', function($scope) {
		console.log("main");
		$scope.message = "test main";
	}])

})();



// .controller('AbcCtrl', ['$scope', function($scope) {
// 	console.log("abc");
// 	$scope.message = "test Abc main";
// }]);