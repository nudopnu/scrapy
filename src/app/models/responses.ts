export interface LoginResponse {
    id:            number;
    username:      string;
    created_at:    Date;
    updated_at:    Date;
    access_token:  string;
    refresh_token: string;
}

export interface RefreshResponse {
    access_token: string;
}

export interface AgentResponse {
    id:              number;
    name:            string;
    user_id:         number;
    last_fetched_at: NullTime;
    created_at:      Date;
    updated_at:      Date;
}

export interface NullTime {
    Time:   Date;
    Valid:  boolean;
}