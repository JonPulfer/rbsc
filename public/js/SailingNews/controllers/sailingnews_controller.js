SailingNews.ArticlesController = Ember.ArrayController.extend({

});

/*
SailingNews.WeatherController = Ember.ObjectController.extend({	
	weatherJSON: function() {
		return this.get('json')
	}.property('weatherJSON')
});

var weather = SailingNews.Weather.create();
var weatherController = SailingNews.WeatherController.create({ model: weather });

SailingNews.WeatherView.create({ controller: weatherController });

jQuery.getJSON("http://api.openweathermap.org/data/2.5/weather?q=rollesby,uk&units=metric", function(json) {
  weather.setProperties(json);
});
*/