start cmd.exe /k "MetaDataModule\venv\Scripts\python.exe MetaDataModule\main.py --port=8051 --ip=127.0.0.1"
go build 
start cmd.exe /k "FilesWithTag.exe"