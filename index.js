// pg package was installed using npm install pg
// initialise the pg package in the Node.js script to get the Client from it
const { Client } = require('pg');
// below I am creating a psql client object that contains the database credentials
const client = new Client({
	user:'postgres',
	password:'pass',
	host:'localhost',
	post:'5432'
});

//below i am using the client to connect to the database using the connect method 
client.connect()
//below .then() and .catch() methods to check that the connect method worked successfully or had an error.
	.then(() => {
		console.log('Connected to the database');
	})
	.catch((err) => {
		console.error('Error occured',err);
	});



