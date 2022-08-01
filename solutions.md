Language: GOLANG

API Router Implementation:  goriella/mux

Github Repo:  mikenewlin/fearless-trial-project
              git clone https://github.com/mikenewlin/fearless-trial-project <local directory>

Local build:  install go and execute in the terminal - go run main.go
              the rest /items api can be accessed with localhost:3000/items
Docker build: install docker and retrieve the application from github.
              docker build -t <local directory of the git clone> .

Docker run:   docker run -p 3000:3000 fearless-trial-project

Testing:      Used postman from the testing:
              
API:          GET - http://localhost:3000/items
              DELETE - http://localhost:3000/items
              POST - http://localhost:3000/items using the following JSON
              [{id: "1", name: "books"}, {id: "2", name: "cards"}] 
              GET - http://localhost:3000/items/{id}
              DELETE - http://localhost:3000/items/{id}            

Future Updates:
    1.  implement the update route - to change the name for an id
    2.  include additional error handling in the code
    3.  include returning the proper HTTP Response Code when the item doesn't exist in the data
    4.  add the database access with the proper SQL statements for each route
    5.  build unit testing and integration testing
    6.  code includes creating the array with one item at startup - remove
    7.  apply company coding standards - style and naming variables and etc...
    
