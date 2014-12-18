window.SailingNews = Ember.Application.create();

SailingNews.WeatherInfoComponent = Ember.Component.extend({
	weatherPromise: Ember.$.getJSON('http://api.openweathermap.org/data/2.5/weather?q=rollesby,uk&units=metric'),
	
	windSpeedMPH: function() {
		return this.weatherPromise.responseJSON.wind.speed*2.237;
	}.property('windSpeedMPH'),
});