## 📝 Project
Best Route Travel

## Description :airplane:
With this application you can find the best route for yor travel.  
The application uses dijkstra's algorithm to test all possible routes and find the cheapest route for you.

### Endpoints

##### Get the cheapest route
`GET http://localhost:3000/travel/:where_from/:where_to/bestroute/`

##### Add a new route
`GET http://localhost:3000/travel/`


## ▶️ Start application 💻 

### ❗ First of all:
First you need to input a csv file with the routes and prices into the app.  
To do this, you need to put your csv file in the <b>same path</b> with the file *<b>mysolution</b>*.  
  
<b>❗Important:</b> the file must be a csv and need have a format like this:
```csv
GRU,BRC,10
BRC,SCL,5
GRU,CDG,75
GRU,SCL,20
GRU,ORL,56
ORL,CDG,5
SCL,ORL,20
```  
The format is:  *<b>WHERE FROM</b>*, *<b>WHERE TO</b>* and the *<b>PRICE</b>* of the the trip.  
  
To do the input of the file, you just need to run `./mysolution <filename.csv>` where the *<filename.csv>* is the name of your file

### The application has 2 interfaces:  
#### 1 - Command line interface:
  - To start this interface app, just need to run in your terminal: `$ ./mysolution cli`  
    
#### 2 - REST API interface
  - To start this rest app, just run the command in your terminal: `$ ./mysolution`
  
### Example:
#### 1 - Command line interface:  
- When you start the app, you will see a message like this:
  ```shell
  $ ./mysolution cli  
  Hi, with this program we'll find the cheapest route for your travel!  
  ---------------------  
  
  please enter the route:  
  ```
- You need to input the route respecting the format: `FROM-TO`  
  Example: `GRU-CDG` and the response will show for route the trip, with connections if it has and the total price:  
  `best route: GRU - BRC - SCL - ORL - CDG > $40`  
  
  
#### 2 - REST API interface:  
Start the app: `$ ./mysolution`  
  
With the [Postman](https://www.postman.com/) or [Insomnia](https://insomnia.rest/) to do the requests you can:  
  
- Get the best/cheapest route in the endpoint `http://localhost:3000/travel/:where_from/:where_to/bestroute/`  
  Example: `GET -> http://localhost:3000/travel/GRU/CDG/bestroute/` and the return will be a 200 http status code with the body:  
  ```shell
  {
    "route": "GRU - BRC - SCL - ORL - CDG",
    "price": 40
  }
  ```  
  
- Add a new route into the file in the endpoint `http://localhost:3000/travel/`  
  Example: `POST -> http://localhost:3000/travel/` and send a json body:
  ```shell
  {
    "where_from": "BRC",
    "where_to": "ORL",
    "price": 10
  }
  ```  
  If the request as process with success the response will be a 204 http status code without a body.  
  
## 💻 Tech infos: 
- The architecture of files and folders was based on domain driven design.
- I tried to use the best pratices of development 
  - Using interfaces and contract to not a create layers dependencies
  - Principles of single responsability
  - Unit tests

:bowtie:
<br><br>
