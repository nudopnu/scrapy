export interface LoginResponse {
    id:            number;
    username:      string;
    created_at:    Date;
    updated_at:    Date;
    access_token:  string;
    refresh_token: string;
}
