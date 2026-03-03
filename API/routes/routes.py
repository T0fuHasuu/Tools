from fastapi import APIRouter
from services.core import *
from fastapi import Request

from services.speedtest import run_speedtest

router = APIRouter()

@router.get("/hostname")
def hostname_route(request: Request):
    return {get_hostname() : get_client_ip(request)}

@router.get("/speedtest")
def speedtest_route():
    dl_spd, ul_spd, ping, duration = run_speedtest()
    return {
        "download_speed": f"{round(dl_spd, 2)} Mbps",
        "upload_speed": f"{round(ul_spd, 2)} Mbps",
        "latency": f"{round(ping, 2)} ms",
        "test_duration": f"{round(duration, 2)} seconds"
    }