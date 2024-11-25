import pytest
import requests

BASE_URL = "http://localhost:8000"


def test_post_food():
    endpoint = "/api/v1/food"
    url = BASE_URL + endpoint
    food = {"name": "bam-chipotle", "protein": 100, "fats": 100, "carbs": 100}
    response = requests.post(url, json=food)

    assert response.status_code == 201

    data = response.json()
    assert data["name"] == food["name"]
    assert data["protein"] == food["protein"]
    assert data["fats"] == food["fats"]
    assert data["carbs"] == food["carbs"]
    assert "id" in data
