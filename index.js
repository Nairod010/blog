const{Pool,Client}= require('pg');

const connectionString='postgressql://dorian:pass@localhost:5432/test'

const client= new Client({
	connectionString: connectionString
})

client.connect()
client.query('Select * from test;',(err,res)=> {
	console.log(err,res)
	client.end()
})
