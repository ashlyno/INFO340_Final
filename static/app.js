'use strict';

angular.module('myApp', ['ngSanitize','ui.router'])
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


		$http({
			 method: 'POST',
	 		 data: {'firstname': $scope.firstName, 'lastname': $scope.lastName}, 
	 		 url: '/addressQuery'
			}).then(function successCallback(response) {
		    	var queryDiv = angular.element($scope.firstQuery);
		    	// queryDiv.append(data[]);
		    	queryDiv.append(response["data"]);
		    	// 		    console.log(data)
		    	// console.log(response["data"])
		    	// $scope.firstQuery.(response);
		    }, function errorCallback(response) {
		    // called asynchronously if an error occurs
		    // or server returns response with an error status.
		});	
    }

}])






