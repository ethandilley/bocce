import pytest
import requests

BASE_URL = "http://server:8000"
#
# def test_get_food():
#     endpoint = "/api/v1/food"
#     url = BASE_URL + endpoint
#     response = requests.get(url)
#
#     assert response.status_code == 200
#
#     json_data = response.json()
#     print(json_data)
#     assert json_data == []

def test_post_food():
    endpoint = "/api/v1/food"
    url = BASE_URL + endpoint
    food = {"name": "bam-chipotle", "protein": 100, "fat": 100, "carb": 100}
    response = requests.post(url, json=food)

    assert response.status_code == 201

    data = response.json()
    assert data["name"] == food["name"]
    assert data["protein"] == food["protein"]
    assert data["fat"] == food["fat"]
    assert data["carb"] == food["carb"]
