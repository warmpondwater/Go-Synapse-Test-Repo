import os
import sqlite3
from typing import Dict, Any

# (DEAD CODE)
def abandoned_python_utility():
    unused_var = 1234
    print(f"Dead code: {unused_var}")

class AnalyticsEngine:
    def __init__(self, db_path: str = "analytics.db"):
        self.db_path = db_path
        self.event_count = 0

    # (SOURCE)
    def fetch_untrusted_event(self) -> Dict[str, Any]:
        return {"event_name": "login_attempt", "payload": "user'; DROP TABLE logs;--"}

    # (SANITIZER)
    def sanitize_event_payload(self, payload: str) -> str:
        if not payload:
            return ""
        return payload.replace("'", "''").replace(";", "")

    # (SINK)
    def persist_analytics_event(self, event_name: str, clean_payload: str) -> bool:
        print(f"[PYTHON SINK DB]: Persisting {event_name} -> {clean_payload}")
        return True

    def track_event(self, event_name: str, payload: str) -> bool:
        self.event_count += 1
        clean = self.sanitize_event_payload(payload)
        self.persist_analytics_event(event_name, clean)
        return self._write_log(event_name, clean)

    def _write_log(self, event_name: str, data: str) -> bool:
        conn = sqlite3.connect(self.db_path)
        cursor = conn.cursor()
        cursor.execute("CREATE TABLE IF NOT EXISTS events (name TEXT, data TEXT)")
        cursor.execute("INSERT INTO events VALUES (?, ?)", (event_name, data))
        conn.commit()
        conn.close()
        return True

if __name__ == "__main__":
    engine = AnalyticsEngine()
    event_data = engine.fetch_untrusted_event()
    engine.track_event(event_data["event_name"], str(event_data["payload"]))

