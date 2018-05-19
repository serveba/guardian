# Guardian

Project for health-checking different services (datasources) configured through forms or yaml files. The healthchecks can be connected to several appenders in order to send notifications when something bad happens.

## Features

* **Datasource**; endpoints and datasources we want to monitor.
* **Appender**; Those objects are used to abstract the concept of sending alarms through several services (slack, email, telegram...)
* Stores information internally with SQLite.
* Admin Dashboard made with datatables and bootstrap all the datasources and its healthchecks.
* Real time data.

## Building guardian

TODO

## Deploying guardian

TODO

## Connectors

In the guardian project we can have several type of services to check. For now we have the following ones:

* **Datasource**: This service abstracts a database service. We can check the connection of the database.
* **Endpoint**: This service abstracts a typical REST/SOAP service endpoint. Through this kind of service we can make request to some status endpoint to see if the service responds correctly, for example: HTTP 200 code, less time than some fixed value...

We are going to describe the structure of the different services through this document.

### Datasources

A datasource has the following data attributes:

* **datasource_type**: this is the supported database type, for now we support 'postgresql', 'mysql' and 'mongodb'.
* **polling_frequency**: integer number in seconds, i.e: 120 -> every 2 minutes we make a connection ping to see if its working
* **warning_timeout**: the max time in seconds it takes to consider a warning
* **error_timeout**: the max time in seconds it takes to consider bad behaviour.

### Appenders

* **type**: this is the supported endpoint type, for now we support 'rest'.


## Roadmap

We have a public Trello board to see the current development status regarding features:

https://trello.com/b/96ncBgpS/guardian

Bugs and feature requests will be addressed through issues.