import os
from fastapi import FastAPI
from dotenv import load_dotenv
from routes.routes import router
from fastapi.responses import JSONResponse


# Variables
load_dotenv()

app = FastAPI(
    title="T0fu",
    description="API Tools For Personal Uses",
    version="1.0.0",
    contact={
        "name": "T0fu",
        "email": "dayuththy@gmail.com",
    },
)

app.include_router(router)

@app.get("/")
async def root():
    return {"message": "Root"}


if __name__ == "__main__":
    import uvicorn
    uvicorn.run("app:app", host=os.getenv("host"), port=int(os.getenv("port")), reload=True) 