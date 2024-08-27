import socket
import sys


class Server:
    def __init__(self,host="localhost", port=8000):
        self.port = port
        self.host = host

    def StartServer(self):
        self.server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
        self.server.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
        self.server.bind((self.host, self.port))
        self.server.listen(1)

