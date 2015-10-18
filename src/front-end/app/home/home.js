(function() {
	'use strict';

	angular.module('app.home')

	.controller('HomeCtrl', Home);

	Home.$inject = ['$q'];

	function Home($q) {
		$q.ddd = "test"
	}

})();