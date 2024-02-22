**Video Games Management API**
**Introduction**
This project is a simple RESTful API for managing video games. It allows users to retrieve information about games, add new games, update existing games, and delete games. The API is built with Go using the Gorilla Mux router for handling HTTP requests.

**Setup**
**Requirements**
*Go (version 1.14 or higher)
*Gorilla Mux

**Installation**
Clone the repository to your local machine:
`git clone https://yourrepositoryurl.com/path/to/repo`

Navigate to the project directory:
`cd path/to/project`

Install dependencies:
`go get -u github.com/gorilla/mux`

**Running the Server**
To start the server on port 8000, run:
`go run .`

**API Endpoints**

**Get All Games**
 *GET /games
 *Retrieves a list of all video games.
 
**Get a Single Game**
 *GET /games/{id}
 *Retrieves details of a specific game by ID.
 
**Create a New Game**
*POST /games
*Adds a new game to the collection. The request body should contain the game details in JSON format.

**Update an Existing Game**
*PUT /games/{id}
*Updates the details of an existing game. The request body should contain the updated game details in JSON format.

**Delete a Game**
*DELETE /games/{id}
*Removes a game from the collection by ID.

**Data Model
Game**
*ID: Unique identifier for the game.
*Isbn: ISBN number of the game.
*Title: Title of the game.
*Developer: Object containing the first and last name of the game's developer.

**Developer**
*FirstName: First name of the developer.
*LastName: Last name of the developer.






