#!/usr/bin/python3 -u

from fastapi import FastAPI

app = FastAPI()

@app.get("/status")
async def status():
    return {"status": "OK"}
