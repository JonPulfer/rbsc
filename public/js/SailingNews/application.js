window.App = Ember.Application.create();

App.Weather = Ember.Object.extend({
	init: function() {
		var t = this
		Ember.$.getJSON("http://api.openweathermap.org/data/2.5/weather?q=rollesby,uk&units=metric").then(function(responseJSON) {
			t.set('wx', responseJSON);	
		});
	},
});

Ember.Handlebars.helper('windSpeed', function(value) {
	return value/(1609.44/3600)
})