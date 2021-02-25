# APIMAN

## Description
Creates an API with predefined handlers. Handlers are listed in the endpoints table and always return the request result in JSON.
enpoints table:
- project - project, any name
- endpoint - project handler, any name
- query - user query to any other table in this database.
At startup, a string is formed from "project" and "endpoint" ("/project/edpoint"), which must be unique. Take this into when adding new lines to enpoints.

## how to start

* Create a database in postgres and a user with access to it. Specify connection details in appconfig.yml
* Apply migrate.sql to create the required structure and add custom data.
* go run apiman.go

## The output will be three links:

* http://localhost:8080/project1/getsrvprod
* http://localhost:8080/project1/getsrvtest
* http://localhost:8080/project1/getallsrv
