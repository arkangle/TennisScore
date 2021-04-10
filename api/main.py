#!/usr/bin/python3 -u

import json
from typing import List
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
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

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_methods=["*"],
    allow_headers=["*"],
    allow_credentials=True
)
@app.get("/status")
async def status():
    return {"status": "OK"}

@app.post("/score")
def posted_score(score: Score):
    _key = bytes(str(score.id), 'utf-8')
    _value = bytes(json.dumps(score.dict()),'utf-8')
    producer = KafkaProducer(bootstrap_servers=['kafka:9093'])
    producer.send('score', key=_key, value=_value)
    return {"status": "CREATED"}
