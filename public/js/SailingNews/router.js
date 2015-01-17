App.Router.map(function() {
	this.resource('articles', { path: '/' });
	this.route('article', { path: '/article/:article_id'});
});

App.ArticlesRoute = Ember.Route.extend({
 	model: function() {
    	// the model is an Array of all of the articles
    	return this.store.find('article');
  	},
	
	renderTemplate: function() {
		this.render({
			into: 'application',
			outlet: 'content',
		});
	},
	
});

App.ArticleRoute = Ember.Route.extend({
	model: function(params) {
		return this.store.find('article', params.article_id);
	},
	renderTemplate: function() {
   		this.render({
			into: 'application',
			outlet: 'content',
		});
  	},
});
