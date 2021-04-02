#!/usr/bin/python3 -u

from fastapi import FastAPI
from pydantic import BaseModel
from kafka import KafkaProducer
import json

app = FastAPI()

class Score(BaseModel):
    game: str

class Point(BaseModel):
    score: Score
    won: int
    id: str

@app.get("/status")
async def status():
    return {"status": "OK"}

@app.post("/point")
def posted_point(point: Point):
    _key = bytes(point.id, 'utf-8')
    _value = bytes(json.dumps(point.dict()),'utf-8')
    producer = KafkaProducer(bootstrap_servers=['kafka:9092'])
    producer.send('point', key=_key, value=_value)
    return {"status": "CREATED"}
