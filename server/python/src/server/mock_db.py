from server.models import DB_Food

DB: list[DB_Food] = [
    DB_Food(id=1, name="dummy", protein=1, carb=1, fat=1),
    DB_Food(id=2, name="fake", protein=1, carb=1, fat=1),
]
