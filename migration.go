package main

import (

)

func migrateDB() {
  // database.Exec(createFoodsTable)
  database.Exec(createMealsTable)
  // database.Exec(createMealFoodsTable)
}

func seedDB() {
  // database.Exec("TRUNCATE TABLE meal_foods RESTART IDENTITY")
  // database.Exec("TRUNCATE TABLE foods RESTART IDENTITY")
  // database.Exec("TRUNCATE TABLE meals RESTART IDENTITY")

  database.Exec(seedMeals)
  // database.Exec(seedFoods)
  // database.Exec(seedMealFoods)
}

const seedMeals = `INSERT INTO meals (id, name)
VALUES (1, 'Breakfast'),
       (2, 'Snack'),
       (3, 'Lunch'),
       (4, 'Dinner')`

// const seedFoods = `INSERT INTO foods (id, name, calories)
// VALUES (1, 'Eggs', 90),
//        (2, 'Slim Jim', 150),
//        (3, 'Ham Sammich', 500),
//        (4, 'Cake', 1200)`
//
// const seedMealFoods = `INSERT INTO meal_foods (id, meal_id, food_id)
// VALUES (1, 1, 1),
//        (2, 2, 2),
//        (3, 3, 3),
//        (4, 4, 4)`

const createFoodsTable = `CREATE TABLE IF NOT EXISTS foods
(
id SERIAL,
name TEXT,
calories INT,
CONSTRAINT food_pkey PRIMARY KEY (id)
)`

const createMealsTable = `CREATE TABLE IF NOT EXISTS meals
(
id SERIAL,
name TEXT,
CONSTRAINT meals_pkey PRIMARY KEY (id)
)`

const createMealFoodsTable = `CREATE TABLE IF NOT EXISTS meal_foods
(
id SERIAL,
meal_id INT,
food_id INT,
CONSTRAINT meal_foods_pkey PRIMARY KEY (id)
)`
