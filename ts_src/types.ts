export interface UserProfile {
    id: string;
    username: string;
    email: string;
}

export interface IUserManager {
    registerUser(profile: UserProfile): void;
    getUserCount(): number;
}
