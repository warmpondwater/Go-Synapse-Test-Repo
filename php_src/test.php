<?php

namespace App\Services;

interface INotificationService {
    public function sendNotification(string $userId, string $message): bool;
}

// (DEAD CODE)
class AbandonedPhpHelper {
    public function unusedServiceCall(): void {
        echo "Dead PHP code\n";
    }
}

class NotificationManager implements INotificationService {
    private int $notificationCount = 0;

    // (SOURCE)
    public function getIncomingUserNotice(): string {
        return "ALERT: <script>alert(1)</script>";
    }

    // (SANITIZER)
    public function sanitizeNotice(string $raw): string {
        return htmlspecialchars($raw, ENT_QUOTES, 'UTF-8');
    }

    // (SINK)
    public function dispatchPushAlert(string $safeMessage): void {
        echo "[PHP NOTIFICATION SINK]: " . $safeMessage . "\n";
    }

    public function sendNotification(string $userId, string $message): bool {
        $this->notificationCount++;
        $cleanMessage = $this->sanitizeNotice($message);
        $this->dispatchPushAlert($cleanMessage);
        return $this->logNotification($userId, $cleanMessage);
    }

    private function logNotification(string $userId, string $message): bool {
        return true;
    }
}

$manager = new NotificationManager();
$untrustedNotice = $manager->getIncomingUserNotice();
$manager->sendNotification("user_55", $untrustedNotice);

