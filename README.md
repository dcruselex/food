# food
This project is for the Go project for CL.
Project name is "Food".
# Overview:
Pulling food data which includes several variables for each item, calculate with items meet the input criteria, then list the items.

Requirements met:
The program will be pulling from an API called Receipe - Food - Nutrition from the rapidapi site "Recipe - Food - Nutrition"
This will satisfy the "Connect to an external/3rd party API and read data into your app." requirement.

The program will create a list, populate it with several values retrieved from the API.  Then the program will retrieve a value from the list and use in the program.
This will satisfy the "Create a dictionary or list, poplate it with several values, retrieve at least one value and use it in your program" requirement.

The ptograme will calculate and display data based on minutes to cook from the list pulled.
This will satisfy the "Calculate and display data based on an external factor (ex: get the current data, and display how many days remaining until some event).

The program will also include 3 unit tests for the application.
This will satisfy the "Create 3 or more unit tests for your application".

The program also includes a UX at localhost:8888 that will display the burgers that meet the prep time criteria and their specific recieipe.
This will satisfy the "Visualize data in a graph, chart, or other visual representation of data" requirement.

# Instructions 

- clone repository: `git clone git@github.com:dcruselex/food.git`
- run the program: `go run ./cmd/main.go`
- open the browser and go to [http://localhost:8888](http://localhost:8888)
- choose time to prep from the pull down and click on the GO button
