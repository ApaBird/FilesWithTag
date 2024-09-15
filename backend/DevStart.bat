cd MetaDataModule
start cmd.exe /k "uvicorn main:app --port 8051"
cd ..
go build 
start cmd.exe /k "FilesWithTag.exe"