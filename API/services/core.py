import socket

def get_hostname() -> str:
    return socket.gethostname()

def get_client_ip(request) -> str:
    return request.client.host

