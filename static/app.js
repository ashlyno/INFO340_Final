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
		// console.log($scope.firstName);
		// console.log($scope.lastName);
		if($scope.type == "patient"){
			$http({
				 method: 'POST',
		 		 data: {'firstname': $scope.firstName, 'lastname': $scope.lastName}, 
		 		 url: '/patientQuery'
				}).then(function successCallback(response) {
			    	var queryDiv = angular.element($scope.firstQuery);
			    	// queryDiv.append(data[]);
			    	queryDiv.append("<h2>"+$scope.firstName+" "+$scope.lastName+"</h2>")
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
		} else {
			$http({
				 method: 'POST',
		 		 data: {'firstname': $scope.firstName, 'lastname': $scope.lastName}, 
		 		 url: '/dentistQuery'
				}).then(function successCallback(response) {
			    	var queryDiv = angular.element($scope.firstQuery);
			    	// queryDiv.append(data[]);
			    	queryDiv.append("<h2>"+$scope.firstName+" "+$scope.lastName+"</h2>")
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

    	$scope.registerList = registerList;

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
		    	var regiDiv = angular.element($scope.registerList);
		    	// queryDiv.append(data[]);
		    	regiDiv.append(response["data"]);
		    	// console.log(response["data"]);
		    }, function errorCallback(response) {
		    // called asynchronously if an error occurs
		    // or server returns response with an error status.
		});
	}


 //    $scope.appointIt = function() {	
 //    	var date = new Date();
 //    	$scope.today = date.getFullYear() + '-' + ('0' + (date.getMonth() + 1)).slice(-2) + '-' + ('0' + date.getDate()).slice(-2);
 //    	console.log(today)
		
	// 	$http({
	// 		 method: 'POST',
	//  		 data: {'firstname': $scope.firstName, 
	//  		 		'lastname': $scope.lastName,
	//  		 		'email': $scope.email,
	//  		 		'today':$scope.today}, 
	//  		 url: '/appointQuery'
	// 		}).then(function successCallback(response) {
	// 	    	var queryDiv = angular.element($scope.appointList);
	// 	    	// queryDiv.append(data[]);
	// 	    	queryDiv.append("<h2> Today is"+$scope.today+"</h2>")
	// 	    	queryDiv.append(response["data"]);
	// 	    	// console.log(data)
	// 	    	// console.log(response["data"])
	// 	    	// $scope.firstQuery.(response);
	// 	    }, function errorCallback(response) {
	// 	    // called asynchronously if an error occurs
	// 	    // or server returns response with an error status.
	// 	});
	// }
}])

// .controller('secondCtrl', ['$scope', '$http','$state', function($scope, $http, $state) { 
  	
//     $scope.appointIt = function() {	
//     	var date = new Date();
//     	$scope.today = date.getFullYear() + '-' + ('0' + (date.getMonth() + 1)).slice(-2) + '-' + ('0' + date.getDate()).slice(-2);
//     	console.log(today)
		
// 		$http({
// 			 method: 'POST',
// 	 		 data: {'firstname': $scope.firstName, 
// 	 		 		'lastname': $scope.lastName,
// 	 		 		'email': $scope.email,
// 	 		 		'today':$scope.today}, 
// 	 		 url: '/appointQuery'
// 			}).then(function successCallback(response) {
// 		    	var queryDiv = angular.element($scope.appointList);
// 		    	// queryDiv.append(data[]);
// 		    	queryDiv.append("<h2> Today is"+$scope.today+"</h2>")
// 		    	queryDiv.append(response["data"]);
// 		    	// console.log(data)
// 		    	// console.log(response["data"])
// 		    	// $scope.firstQuery.(response);
// 		    }, function errorCallback(response) {
// 		    // called asynchronously if an error occurs
// 		    // or server returns response with an error status.
// 		});
// 	}

//     $scope.registerIt = function() {	
//     	console.log("register")
//     	console.log($scope.firstname)
//     	console.log($scope.lastName)
//     	    	console.log($scope.email)
//     	    	    	console.log($scope.dob)
    	    	    	
//     	var date = $scope.birth
//     	$scope.dob = date.getFullYear() + '-' + ('0' + (date.getMonth() + 1)).slice(-2) + '-' + ('0' + date.getDate()).slice(-2);

// 		$http({
// 			 method: 'POST',
// 	 		 data: {'firstname': $scope.firstName, 
// 	 		 		'lastname': $scope.lastName,
// 	 		 		'email': $scope.email,
// 	 		 		'date': $scope.dob}, 
// 	 		 url: '/registerQuery'
// 			}).then(function successCallback(response) {
// 		    	var queryDiv = angular.element($scope.registerList);
// 		    	// queryDiv.append(data[]);
// 		    	queryDiv.append(response["data"]);
// 		    	// console.log(data)
// 		    	// console.log(response["data"])
// 		    	// $scope.firstQuery.(response);
// 		    }, function errorCallback(response) {
// 		    // called asynchronously if an error occurs
// 		    // or server returns response with an error status.
// 		});
// 	}
// }])






