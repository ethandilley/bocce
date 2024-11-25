from pydantic import BaseModel


class Food(BaseModel):
    name: str
    protein: int
    fat: int
    carb: int


class DB_Food(BaseModel):
    id: int
    name: str
    protein: int
    fat: int
    carb: int
