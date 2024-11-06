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
    thumbnail_url?:  string;
}

export interface NullTime {
    Time:   Date;
    Valid:  boolean;
}

export interface Ad {
    id:          string;
    title:       string;
    price:       string;
    description: string;
    location:    string;
    postal_code: string;
    category_id: string;
    posted_at:   string;
    url:         string;
    created_at:  Date;
    updated_at:  Date;
    images:      Image[];
}

export interface Image {
    image_url:    string;
    image_number: number;
}
