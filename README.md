# JetDev
This project is based on domain driven development which uses 
following technologies :
Gin web framework
GORM for mysql 

Execution is done using a docker-compose to remove local 
mysql installation.
repo >> service >> controller >> routing >> server

It is divided into repository layer at based derived from 
models which is then using service layer for communication to 
controller layer;
Controller layer is then takes input from user and processes it further
Route layer takes api routes to be defined from handler and then served using the 
server.