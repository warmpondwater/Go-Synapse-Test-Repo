import { UserProfile, IUserManager } from './types';


// (DEAD CODE)
export class AbandonedTsHelper {
    public discard(): void {
        const isolated = 777;
        console.log("Dead TS code", isolated);
    }
}

export class UserManager implements IUserManager {
    private users: Map<string, UserProfile> = new Map();
    private registrationCount: number = 0;

    // (SOURCE)
    public getIncomingRegistrationPayload(): UserProfile {
        return {
            id: "usr_101",
            username: "<script>alert('xss')</script>",
            email: "alex@example.com"
        };
    }

    // (SANITIZER)
    public sanitizeProfile(profile: UserProfile): UserProfile {
        return {
            id: profile.id,
            username: profile.username.replace(/</g, "&lt;").replace(/>/g, "&gt;"),
            email: encodeURI(profile.email)
        };
    }

    // (SINK)
    public persistUserToDatabase(profile: UserProfile): void {
        console.log(`[TS SINK DB]: Persisting clean user profile: ${profile.id} -> ${profile.username}`);
    }

    public registerUser(profile: UserProfile): void {
        this.registrationCount++;
        const cleanProfile = this.sanitizeProfile(profile);
        this.persistUserToDatabase(cleanProfile);
        this.users.set(cleanProfile.id, cleanProfile);
        this.sendWelcomeEmail(cleanProfile.email);
    }

    public getUserCount(): number {
        return this.registrationCount;
    }

    private sendWelcomeEmail(email: string): void {
        console.log(`Sending welcome email to ${email}`);
    }
}

const manager: IUserManager = new UserManager();
const concreteManager = new UserManager();
const rawProfile = concreteManager.getIncomingRegistrationPayload();
manager.registerUser(rawProfile);

