import time
import requests

def send_requests():
    status_429_counter = 0
    status_200_counter = 0
    while True:
        for j in range(150):
            response = requests.get('http://localhost:8888')
            if response.status_code == 429:
                status_429_counter += 1
            elif response.status_code == 200:
                status_200_counter += 1
        print(f"Got {status_429_counter} http.StatusTooManyRequests [429] | Got {status_200_counter} http.StatusOK [200]")
        status_429_counter = 0
        status_200_counter = 0
        time.sleep(1)

if __name__ == "__main__":
    send_requests()
