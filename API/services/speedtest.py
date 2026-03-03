import speedtest
import time

def run_speedtest():

    start_time = time.time()
    
    st = speedtest.Speedtest()
    st.get_best_server()

    download_speed = st.download() / 1_000_000  
    upload_speed = st.upload() / 1_000_000 
    ping = st.results.ping
    duration = time.time() - start_time

    return download_speed, upload_speed, ping, duration

