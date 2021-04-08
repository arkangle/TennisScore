#!/usr/bin/python3 -u

import json
from typing import List
from fastapi import FastAPI
from pydantic import BaseModel
from kafka import KafkaProducer

app = FastAPI()

class State(BaseModel):
    game:  List[int]
    deuce: int
    set:   List[int]
    match: List[int]

class Score(BaseModel):
    state:  State
    winner: int
    id:     int

@app.get("/status")
async def status():
    return {"status": "OK"}

@app.post("/score")
def posted_score(score: Score):
    _key = bytes(score.id, 'utf-8')
    _value = bytes(json.dumps(score.dict()),'utf-8')
    producer = KafkaProducer(bootstrap_servers=['kafka:9092'])
    producer.send('score', key=_key, value=_value)
    return {"status": "CREATED"}
