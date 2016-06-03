'use strict';

angular.module('myApp', ['ngSanitize','ui.router'])
	.config(function($stateProvider,$urlRouterProvider){

		$stateProvider
			.state('patients', {
				url: '/patients', //"root" directory
				templateUrl: '/patients',
				controller: 'MyCtrl'
			})

			.state('appointment', {
				url: '/appointment', //"root" directory
				templateUrl: '/appointment',
				controller: 'MyCtrl'
			})

			.state('registration', {
				url: '/register',
				templateUrl: '/register',
				controller: 'MyCtrl'
			})

		// For any unmatched url, redirect to "home"
		$urlRouterProvider.otherwise('/');

	})
		
  .controller('MyCtrl', ['$scope', '$http','$state', function($scope, $http, $state) { 
  	
  	$scope.hideImg = false;

    $scope.split = function(string, nb) {
    	var array = string.split(' ');
    	return array[nb];
	}

	$scope.printName = function() {
		console.log($scope.name)
		if($scope.name) {
			$scope.firstName = $scope.split($scope.name,0);
			$scope.lastName = $scope.split($scope.name,1); 
		}
	}

	$scope.hideIt = function() {
		$scope.hideImg = true;
	}

	$scope.showIt = function() {
		$scope.hideImg = false;
	}

   //function called to fetch tracks based on the scope's query
    $scope.hideBackgroud = function() {
      angular.element('homeBackground').visibility
    };


    $scope.searchIt = function() {
    	$scope.hideIt();
		$scope.firstQuery = firstQuery;
		$scope.printName();
		// console.log($scope.firstName);
		// console.log($scope.lastName);

		$http({
			 method: 'POST',
	 		 data: {'firstname': $scope.firstName, 'lastname': $scope.lastName}, 
	 		 url: '/patientQuery'
			}).then(function successCallback(response) {
		    	var queryDiv = angular.element($scope.firstQuery);
		    	// queryDiv.append(data[]);
		    	queryDiv.empty();
		    	queryDiv.append(response["data"]);
		    	// 		    console.log(data)
		    	// console.log(response["data"])
		    	// $scope.firstQuery.(response);
		    }, function errorCallback(response) {
		    // called asynchronously if an error occurs
		    // or server returns response with an error status.
		});
    }

    // $httpProvider.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded;charset=utf-8';

	  // var param = function(obj) {
	  //   var query = '', name, value, fullSubName, subName, subValue, innerObj, i;

	  //   for(name in obj) {
	  //     value = obj[name];

	  //     if(value instanceof Array) {
	  //       for(i=0; i<value.length; ++i) {
	  //         subValue = value[i];
	  //         fullSubName = name + '[' + i + ']';
	  //         innerObj = {};
	  //         innerObj[fullSubName] = subValue;
	  //         query += param(innerObj) + '&';
	  //       }
	  //     }
	  //     else if(value instanceof Object) {
	  //       for(subName in value) {
	  //         subValue = value[subName];
	  //         fullSubName = name + '[' + subName + ']';
	  //         innerObj = {};
	  //         innerObj[fullSubName] = subValue;
	  //         query += param(innerObj) + '&';
	  //       }
	  //     }
	  //     else if(value !== undefined && value !== null)
	  //       query += encodeURIComponent(name) + '=' + encodeURIComponent(value) + '&';
	  //   }
	  //   return query.length ? query.substr(0, query.length - 1) : query;
	  // };

	  // // Override $http service's default transformRequest
	  // $httpProvider.defaults.transformRequest = [function(data) {
	  //   return angular.isObject(data) && String(data) !== '[object File]' ? param(data) : data;
	  // }];
}])
