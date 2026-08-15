import os
import sqlite3

class AnalyticsEngine:
    def __init__(self, db_path: str = "analytics.db"):
        self.db_path = db_path

    def track_event(self, event_name: str, payload: dict):
        print(f"Tracking event: {event_name}")
        self._write_log(event_name, str(payload))

    def _write_log(self, event_name: str, data: str):
        conn = sqlite3.connect(self.db_path)
        cursor = conn.cursor()
        cursor.execute("CREATE TABLE IF NOT EXISTS events (name TEXT, data TEXT)")
        cursor.execute("INSERT INTO events VALUES (?, ?)", (event_name, data))
        conn.commit()
        conn.close()

if __name__ == "__main__":
    engine = AnalyticsEngine()
    engine.track_event("page_view", {"url": "/dashboard", "user": "dev_01"})
