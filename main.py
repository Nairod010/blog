import psycopg2

conn = psycopg2.connect("user=postgres password=pass host=localhost port=5432")

#Below I used a temporary memory AKA a cursor to check the version of the psql
'''crsr = conn.cursor()
crsr.execute('SELECT version();')
result = crsr.fetchone()
print(result)
print(type(crsr))'''

