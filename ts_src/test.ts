interface UserProfile {
    id: string;
    username: string;
    email: string;
}

class UserManager {
    private users: Map<string, UserProfile> = new Map();

    public registerUser(profile: UserProfile): void {
        console.log(`Registering user ${profile.username}`);
        this.users.set(profile.id, profile);
        this.sendWelcomeEmail(profile.email);
    }

    private sendWelcomeEmail(email: string): void {
        console.log(`Sending welcome email to ${email}`);
    }
}

const manager = new UserManager();
manager.registerUser({ id: "usr_101", username: "alex", email: "alex@example.com" });
