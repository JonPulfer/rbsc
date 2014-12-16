SailingNews.ArticlesController = Ember.ArrayController.extend({
	weatherPromise: Ember.$.getJSON('http://api.openweathermap.org/data/2.5/weather?q=rollesby,uk&units=metric'),
	
	windSpeedMPH: function() {
		var speedMPS = this.get('weatherPromise.responseJSON.wind.speed');
		var speedMPH = Math.round(speedMPS * 3600 / 1610.3*1000)/1000;
		console.log('did the function');
		return speedMPH
	}.property('calculatedWindSpeedMPH')
});


