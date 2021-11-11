# food
This project is for the Go project for CL.
Project name is "Food".

# Description

Pulling food data(specifically burgers) which includes several variables for each item, calculate the items that meet the input criteria.
The program has used 3 ranges of time to prepare, then list the items and present in the UX.   You can also click on the item and open the specific recipe.

# Requirements met:

1)  I have uploaded the project to my GitHub repository and made at least 5 commits during the changes.

2)  Project includes a README file has the following: 
    2.1  Descripton of what my project is about. 
    2.2  The 3+ features I included from the GO Project Requirements (I have included at least 5).
    2.3  The special instructions reqired for the reviewer to run my project.

3)  Project includes at least one package and one struct.
    3.1 Packages inlcude Main, rapidapi, router
    3.2 Struct is in the models.go file

4)  Create and call at least 3 functions or methods, at least one of which must return a value that is used in my application.
    4.1 There is more than 3 created and called.


5)  Choose at least 3 items on the Features List and implement them in the project.   I have listed 5 below that have been implemented.


    5.1 The program will be pulling from an API called Receipe - Food - Nutrition from the rapidapi site "Recipe - Food - Nutrition"
        This will satisfy the "Connect to an external/3rd party API and read data into your app." requirement.

    5.2 The program will create a list, populate it with several values retrieved from the API.  Then the program will retrieve a value from the list and use in the program.
        This will satisfy the "Create a dictionary or list, poplate it with several values, retrieve at least one value and use it in your program" requirement.

    5.3 The program will calculate and display data based on minutes to cook from the list pulled.
        This will satisfy the "Calculate and display data based on an external factor (ex: get the current data, and display how many days remaining until some event).

    5.4 The program will also include 3 unit tests for the application.
        This will satisfy the "Create 3 or more unit tests for your application".

    5.5 The program also includes a UX at localhost:8888 that will display the burgers that meet the prep time criteria and their specific recieipe.
        This will satisfy the "Visualize data in a graph, chart, or other visual representation of data" requirement.


# Instructions 

- clone repository: `git clone git@github.com:dcruselex/food.git`
- run the program: `go run ./cmd/main.go`
- open the browser and go to [http://localhost:8888](http://localhost:8888)
- choose time to prep from the pull down and click on the GO button
