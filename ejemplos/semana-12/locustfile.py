from locust import HttpUser, TaskSet, task, between
import random
import json

class MyTasks(TaskSet):
    
    @task(1)
    def engineering(self):
        # List of random names
        names = ["guatemala", "mexico", "panama", "inglaterra", "francia", "italia", "españa", "argentina", "chile", "colombia"]

        climas = ["soleado", "nublado", "lluvioso"]
    
        # Student data
        weather_data = {
            "name": random.choice(names),  # Random name
            "temperatura": random.randint(18, 28),  # Random temperature between 18 and 28
            "humedad": random.randint(40, 80),  # Random humidity between 40 and 80
            "clima": random.choice(climas)  # Random weather condition
        }
        
        # Send JSON as POST to the /engineering route
        headers = {'Content-Type': 'application/json'}
        self.client.post("/clima", json=weather_data, headers=headers)

    

class WebsiteUser(HttpUser):
    tasks = [MyTasks]
    wait_time = between(1, 5)  # Wait time between tasks (1 to 5 seconds)