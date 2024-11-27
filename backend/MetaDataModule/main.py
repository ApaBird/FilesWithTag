import fastapi
import metaReader
import os

app = fastapi.FastAPI()

@app.get("/GetMeta")
def getMeta(Path: str):
    json = metaReader.ReadMeta(Path)

    return fastapi.responses.JSONResponse(content=json)
