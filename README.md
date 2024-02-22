**Video Games Management API**
**Introduction**
This project is a simple RESTful API for managing video games. It allows users to retrieve information about games, add new games, update existing games, and delete games. The API is built with Go using the Gorilla Mux router for handling HTTP requests.

**Setup**
**Requirements**
1. Go (version 1.14 or higher)
2. Gorilla Mux

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
 1. GET /games
 2. Retrieves a list of all video games.
 
**Get a Single Game**
 1. GET /games/{id}
 2. Retrieves details of a specific game by ID.
 
**Create a New Game**
1. POST /games
2. Adds a new game to the collection. The request body should contain the game details in JSON format.

**Update an Existing Game**
1. PUT /games/{id}
2. Updates the details of an existing game. The request body should contain the updated game details in JSON format.

**Delete a Game**
1. DELETE /games/{id}
2. Removes a game from the collection by ID.

**Data Model
Game**
1. ID: Unique identifier for the game.
2. Isbn: ISBN number of the game.
3. Title: Title of the game.
4. Developer: Object containing the first and last name of the game's developer.

**Developer**
1. FirstName: First name of the developer.
2. LastName: Last name of the developer.






