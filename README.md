# guardian/clj

Clojure project for checking health on several services.

## Features

* Define services with .yaml files.
* Define connectors with .yaml files. Connectors are used for sending alarms to several targets (email, slack, telegram...)
* Stores information internally with SQLite.
* Admin Dashboard made with reagent with all the services and its data.
* Real time data.
