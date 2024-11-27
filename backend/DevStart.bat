go build 
start cmd.exe /k "FilesWithTag.exe"
cd MetaDataModule

start cmd.exe /k ".\venv\Scripts\activate & uvicorn main:app --port 8051"

