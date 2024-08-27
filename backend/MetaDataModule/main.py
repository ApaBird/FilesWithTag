import tinytag
import fastapi


app = fastapi.FastAPI()


@app.get("/")
async def get():
    return fastapi.Response(content="Hello World", media_type="text/plain")