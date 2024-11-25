from server.models import Food, DB_Food
from server.mock_db import DB

from fastapi import FastAPI

app = FastAPI()


def new_id() -> int:
    highest = -1
    for food in DB:
        if food.id > highest:
            highest = food.id
    return highest + 1


@app.get("/")
def read_root():
    return {"Hello": "World"}


@app.get("/api/v1/food/")
def get_foods() -> list[dict]:
    response = [food.model_dump() for food in DB]
    return response


@app.get("/api/v1/food/{food_id}")
def get_food(food_id: int) -> dict:
    print(DB)
    for food in DB:
        print(food)
        if food.id == food_id:
            return food.model_dump()
    return {"no": "shot"}


@app.post("/api/v1/food/")
def create_food(food: Food) -> dict:
    id = new_id()
    db_entry = DB_Food(
        id=id, name=food.name, protein=food.protein, fat=food.fat, carb=food.carb
    )
    DB.append(db_entry)

    for dbfood in DB:
        if dbfood.id == id:
            return dbfood.model_dump()
    return {"no": "shot"}


@app.delete("/api/v1/food/{food_id}")
def delete_food(food_id: int) -> dict:
    remove_food = None
    for i, dbfood in enumerate(DB):
        if dbfood.id == food_id:
            remove_food = dbfood
            DB.pop(i)
            print(remove_food)
            break
    if not remove_food:
        return {"no": "shot"}
    return remove_food.model_dump()
