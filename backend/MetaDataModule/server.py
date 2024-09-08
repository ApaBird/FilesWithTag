import socket
import sys
import asyncio


class Server:
    def __init__(self,host="localhost", port=8000, mapHandler = None):
        self.port = port
        self.host = host
        self.mapHandler = mapHandler

    def StartServer(self):
        self.server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        #self.server.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        self.server.bind((self.host, self.port))
        self.server.listen(1)

        print(f"Сервер запущен [{self.host}:{self.port}] и ожидает подключений...")

        self.client_socket, self.client_address = self.server.accept()
        print(f"Подключение установлено с {self.client_address}")

        self.HandlerSocket()

    def HandlerSocket(self):
        while True:
            order = self.ReadSocket()
            ans = None
            for handler in self.mapHandler:
                if order.find(handler) != -1:
                    ans = self.mapHandler[handler](order)
                    break
            
            if ans is not None:
                self.client_socket.send(ans.encode() + b"\n")
            else:
                self.client_socket.send("ok".encode() + b"\n")

                

    def ReadSocket(self)-> str:
        order = ""
        while True:
            char = self.client_socket.recv(1)
            if char == b'\n':
                return order
            else:
                try:
                    order += char.decode()
                except:
                    print(char)
       
