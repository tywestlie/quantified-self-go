package main

import (

)

func migrateDB() {
  database.Exec(createFoodsTable)
  database.Exec(createMealsTable)
  database.Exec(createMealFoodsTable)
}


const createFoodsTable = `CREATE TABLE IF NOT EXISTS foods
(
id SERIAL,
name TEXT NOT NULL,
calories INT,
CONSTRAINT food_pkey PRIMARY KEY (id)
)`

const createMealsTable = `CREATE TABLE IF NOT EXISTS meals
(
id SERIAL,
name TEXT NOT NULL,
CONSTRAINT meals_pkey PRIMARY KEY (id)
)`

const createMealFoodsTable = `CREATE TABLE IF NOT EXISTS meal_foods
(
id SERIAL,
meal_id INT,
food_id INT,
CONSTRAINT meal_foods_pkey PRIMARY KEY (id)
)`
