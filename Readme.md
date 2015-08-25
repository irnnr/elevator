#Elevator Control System

##Installation

`GOPATH=$(pwd) gpm install`


##Building and Running the App

`GOPATH=$(pwd) go build -o dist/elevator && ./dist/elevator`


## Testing

`GOPATH=$(pwd) go test`


## ToDos / Potential Improvements

* keep a set of requests
	* choose a set if order of requests should not matter (better to optimize usage of an elevator)
	* choose a queue if we want to process requests in the order they come in (easier to implement)
* randomly trigger new requests and add them to the set of requests
* re-introduce the direction argument for the Pickup method to allow easier scheduling/picking requests from the 
* add a sorted list of next stops to an elevator
* add a maximum capacity for an elevator 
* when stepping an elevator check the request queue for fitting pickups  
	* must go into the same direction
	* must still be ahead of the current floor
	* when a request meets the criteria, add it to the elevator's queue, remove from waiting list set
	* if a new request's goal is closer than the current goal, make that new request's goal the current goal and put the current goal into the elevator's queue
* when reaching the current goal floor pick the next goal from the queue
* assuming that per request we add one person to the elevator we should be able to predict the number of people in the elevator at each stop
	* depending on that we can also decline/not pick a fitting request if the elevator will be full at that time/stop
* since that would be unrealistic we would have to add a random number of people when reaching a goal floor and keep accepting all request
	* the elevator will stop, but may not accept any more people if it the number of people wanting to join would exceed the capacity
	* the request would go back to the pool of waiting requests

