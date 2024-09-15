import fastapi
import metaReader

app = fastapi.FastAPI()

@app.get("/GetMeta")
def getMeta(Path: str):
    json = metaReader.ReadMeta(Path)

    return fastapi.responses.JSONResponse(content=json)
