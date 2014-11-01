var debugSidebar = angular.module("debugSidebar", []);

debugSidebar.controller('debugSidebarMainCtl', function($scope) {
	$scope.visible = false;

	$scope.Clicked = function() {
		if ($scope.visible) {
			$scope.visible = false;
		} else {
			$scope.visible = true;
		}
	};
});