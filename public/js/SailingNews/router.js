SailingNews.Router.map(function() {
	this.resource('articles', { path: '/' });
});

SailingNews.ArticlesRoute = Ember.Route.extend({
 	model: function() {
    // the model is an Array of all of the articles
    return this.store.find('article');
  }
});