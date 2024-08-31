import argparse
import socket

parser = argparse.ArgumentParser(description='A tutorial of argparse!')
parser.add_argument("--port")
parser.add_argument("--ip")


args = parser.parse_args()

server = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
server.bind((args.ip, int(args.port)))
server.listen(1)

print("Сервер запущен и ожидает подключений...")

client_socket, client_address = server.accept()
print(f"Подключение установлено с {client_address}")

order = ""

while True:
    char = client_socket.recv(1)
    if char == b'\n':
        print("Анализируем файл")
        print(order)
        order = ""
    else:
        order += char
        