import argparse
import socket
import server as customServer
import metaReader as handlerMeta


def main():
    parser = argparse.ArgumentParser(description='A tutorial of argparse!')
    parser.add_argument("--port")
    parser.add_argument("--ip")

    args = parser.parse_args()

    handlers = {
        "GetMeta": handlerMeta
    }    

    server = customServer.Server(args.ip, int(args.port), handlers)
    server.StartServer()
            

if __name__ == "__main__":
    main()