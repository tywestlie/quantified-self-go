# Quantified Self: Go

This application is a Go/Mux back-end with a PostgreSQL database. It is designed to serve food, meal and calorie information via JSON payloads to a javaScript front-end found [here](http://roasted-throat.surge.sh/index.html). The Go/Mux API is also deployed [here](https://qs-go.herokuapp.com/).

# Requests
These are the possible requests from the application and their JSON payloads.
## Food Endpoints
`GET /api/v1/foods` Returns all the food in the database:
```javascript
[
    {
        id: 1,
        name: "Bacon",
        calories: 500
    },
    {
        id: 2,
        name: "More Bacon",
        calories: 700
    },
    {
        id: 3,
        name: "Ham Sammy",
        calories: 800
    },
    {
        id: 4,
        name: "Burrito",
        calories: 1200
    }
]
```
`GET /api/v1/foods/:id` Returns a single food object from the `:id` you've pased
```javascript
  {
    id: 4,
    name: "Burrito",
    calories: 1200
  }
```
`POST /api/v1/foods` Creates and returns a new food from the following parameters format
```JSON
{ "food": { "name": "Name of food here", "calories": "Calories here"} }
```
`PATCH /api/v1/foods/:id` Updates a food that already exists in the database based on the following parameters.
```JSON
{ "food": { "name": "Mint", "calories": "14"} }
```
`DELETE /api/v1/foods/:id` Deletes a food with the passed `:id`
## Meal Enpoints
`GET /api/v1/meals` Returns all meals ind the database along with their associated foods:
```JSON
[
    {
        "id": 1,
        "name": "Breakfast",
        "foods": [
            {
                "id": 1,
                "name": "Banana",
                "calories": 150
            },
            {
                "id": 6,
                "name": "Yogurt",
                "calories": 550
            },
            {
                "id": 12,
                "name": "Apple",
                "calories": 220
            }
        ]
    },
    {
        "id": 2,
        "name": "Snack",
        "foods": [
            {
                "id": 1,
                "name": "Banana",
                "calories": 150
            },
            {
                "id": 9,
                "name": "Gum",
                "calories": 50
            },
            {
                "id": 10,
                "name": "Cheese",
                "calories": 400
            }
        ]
    },
    {
        "id": 3,
        "name": "Lunch",
        "foods": [
            {
                "id": 2,
                "name": "Bagel Bites - Four Cheese",
                "calories": 650
            },
            {
                "id": 3,
                "name": "Chicken Burrito",
                "calories": 800
            },
            {
                "id": 12,
                "name": "Apple",
                "calories": 220
            }
        ]
    },
    {
        "id": 4,
        "name": "Dinner",
        "foods": [
            {
                "id": 1,
                "name": "Banana",
                "calories": 150
            },
            {
                "id": 2,
                "name": "Bagel Bites - Four Cheese",
                "calories": 650
            },
            {
                "id": 3,
                "name": "Chicken Burrito",
                "calories": 800
            }
        ]
    }
]
```
`GET /api/v1/meals/:meal_id/foods` Returns all the foods associated with a meal and  specified by `:meal_id`
```JSON
{
    "id": 1,
    "name": "Breakfast",
    "foods": [
        {
            "id": 1,
            "name": "Banana",
            "calories": 150
        },
        {
            "id": 6,
            "name": "Yogurt",
            "calories": 550
        },
        {
            "id": 12,
            "name": "Apple",
            "calories": 220
        }
    ]
}
```
`POST /api/v1/meals/:meal_id/foods/:id` Adds the food with `:id` to the meal with `:meal_id`.
```JSON
{
    "message": "Successfully added FOODNAME to MEALNAME"
}
```
`DELETE /api/v1/meals/:meal_id/foods/:id` Removes the food with `:id` from the meal
```JSON
{
    "message": "Successfully removed FOODNAME to MEALNAME"
}
```
