(function() {
	'use strict';

	angular.module('app.summoner')

	.controller('SummonerCtrl', Summoner);

	// Summoner.$inject = ['$scope'];

	function Summoner() {
		var summonerVm = this;
		summonerVm.name = "test summoner"
	}

})();