<?php

namespace App\Services;

class NotificationManager {
    public function sendNotification(string $userId, string $message): bool {
        echo "Notifying User {$userId}: {$message}\n";
        return $this->logNotification($userId, $message);
    }

    private function logNotification(string $userId, string $message): bool {
        return true;
    }
}

$manager = new NotificationManager();
$manager->sendNotification("user_55", "Your security audit is complete.");
