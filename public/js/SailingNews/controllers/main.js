App.HomeController = Ember.ObjectController.extend({
	init: function() {
		this._super();
		var wx = App.Weather.create();
		this.set('weatherData', wx);
	},
});