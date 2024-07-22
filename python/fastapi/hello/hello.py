from fastapi import FastAPI

app = FastAPI()

@app.get("/")
def hi():
    return {"Hello": "fast api"}