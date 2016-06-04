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
    	angular.element($scope.firstQuery).empty();

    	$scope.hideIt();
		$scope.firstQuery = firstQuery;
		$scope.printName();
		if($scope.type == undefined) {
			$scope.type = "patient"
		}
		// console.log($scope.firstName);
		// console.log($scope.lastName);

    	var queryDiv = angular.element($scope.firstQuery);

    	queryDiv.append("<h2 class='tableTop' style='margin-top: 60px'>"+$scope.firstName+" "+$scope.lastName+"</h2>")
    	console.log($scope.type)
		if($scope.type == "patient"){
			$http({
				 method: 'POST',
		 		 data: {'firstname': $scope.firstName, 'lastname': $scope.lastName}, 
		 		 url: '/patientQuery'
				}).then(function successCallback(response) {
			    	// queryDiv.append(data[]);
			    	queryDiv.append(response["data"]);
			    	// console.log(data)
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
			    	// queryDiv.append(data[]);
			    	queryDiv.append(response["data"]);
			    	// 		    console.log(data)
			    	// console.log(response["data"])
			    	// $scope.firstQuery.(response);
			    }, function errorCallback(response) {
			    // called asynchronously if an error occurs
			    // or server returns response with an error status.
			});	
		} else {
			$http({
				 method: 'POST',
		 		 data: {'firstname': $scope.firstName, 'lastname': $scope.lastName}, 
		 		 url: '/dentistQuery'
				}).then(function successCallback(response) {
			    	queryDiv.append(response["data"]);
			    	// console.log(data)
			    	// console.log(response["data"])
			    	// $scope.firstQuery.(response);
			    }, function errorCallback(response) {
			    // called asynchronously if an error occurs
			    // or server returns response with an error status.
			});

		}
    }


    $scope.registerIt = function() {	

    	var date = $scope.birth
    	$scope.dob = date.getFullYear() + '-' + ('0' + (date.getMonth() + 1)).slice(-2) + '-' + ('0' + date.getDate()).slice(-2);

		$http({
			 method: 'POST',
	 		 data: {'firstname': $scope.firstName, 
	 		 		'lastname': $scope.lastName,
	 		 		'email': $scope.email,
	 		 		'date': $scope.dob}, 
	 		 url: '/registerQuery'
			}).then(function successCallback(response) {

		    }, function errorCallback(response) {
		    // called asynchronously if an error occurs
		    // or server returns response with an error status.
		});
		$scope.registerShow();
		document.forms["regiForm"].reset();		
		// $scope.regiForm = regiForm;
		// $scope.regiForm.$setPristine();
	}

	$scope.registerShow = function() {
    	$scope.registerList = registerList;
    	var queryDiv = angular.element($scope.registerList).empty();

		$http({
			 method: 'POST', 
	 		 url: '/registerShow'
			}).then(function successCallback(response) {
		    	// queryDiv.append(data[]);
		    	queryDiv.append("<h2>All registered patients</h2>");
		    	queryDiv.append(response["data"]);
		    	// console.log(response["data"]);
		    }, function errorCallback(response) {
		    // called asynchronously if an error occurs
		    // or server returns response with an error status.
		});
	}

    $scope.appointIt = function() {	
		$http({
			 method: 'POST',
	 		 data: {'firstname': $scope.firstName, 
	 		 		'lastname': $scope.lastName,
	 		 		'date': $scope.visitDate}, 
	 		 url: '/appointQuery'
			}).then(function successCallback(response) {

		    }, function errorCallback(response) {
		    // called asynchronously if an error occurs
		    // or server returns response with an error status.
		});
	 	$scope.appointShow();
	 	document.forms["appointForm"].reset();
	}

	$scope.appointShow = function() {
    	var date = new Date();
    	$scope.today = date.getFullYear() + '-' + ('0' + (date.getMonth() + 1)).slice(-2) + '-' + ('0' + date.getDate()).slice(-2);
    	console.log($scope.today)

    	$scope.appointList = appointList;
    	var queryDiv = angular.element($scope.appointList).empty();

		$http({
			 method: 'POST', 
			 data: {'today':$scope.today},
	 		 url: '/appointShow'
			}).then(function successCallback(response) {
		    	// queryDiv.append(data[]);
		    	queryDiv.append("<h2> Today is "+('0' + (date.getMonth() + 1)).slice(-2) + '/' + ('0' + date.getDate()).slice(-2)+'/'+date.getFullYear()+"</h2>")
		    	queryDiv.append(response["data"]);
		    	// console.log(response["data"]);
		    }, function errorCallback(response) {
		    // called asynchronously if an error occurs
		    // or server returns response with an error status.
		});
	}

	if(document.querySelector('#registerList')){
		$scope.registerShow();
	}
	if(document.querySelector('#appointList')){
		$scope.appointShow();
	}

}])







