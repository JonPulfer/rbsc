SailingNews.User = DS.Model.extend({
    firstName: DS.attr('string'),
    lastName: DS.attr('string'),
    email: DS.attr('string'),
    salutation: DS.attr('string'),
    password: DS.attr('string'),
    isAdmin: DS.attr('boolean')
});

/*

Expects JSON --

"users": [{

    "id": 1,
    "firstName": "foo",
    "lastName": "foo",
    "email": "foo",
    "salutation": "foo",
    "password": "foo",
    "isAdmin": true

}]

*/