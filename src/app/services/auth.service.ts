import { HttpClient } from "@angular/common/http";
import { inject, Injectable, signal } from "@angular/core";
import { ApiService } from "./api.service";
import { LoginResponse, RefreshResponse } from "../models/responses";
import { tap } from "rxjs";

@Injectable({
  providedIn: "root",
})
export class AuthService {
  private REFRESH_TOKEN_KEY = "refresh_token";
  accessToken?: string;
  currentUser?: string;
  loggedIn = signal(false);
  http = inject(HttpClient);
  BASE_URL = inject(ApiService).BASE_URL;

  login(username: string, password: string, remember: boolean) {
    const url = `${this.BASE_URL}/login`;
    return this.http.post<LoginResponse>(url, { username, password }).pipe(
      tap((res) => {
        this.setAccessToken(res.access_token);
        this.setRefreshToken(res.refresh_token, remember);
      }),
    );
  }

  refresh(sessionToken: string) {
    const url = `${this.BASE_URL}/refresh`;
    const headers = { "Authorization": `Bearer ${sessionToken}` };
    return this.http.post<RefreshResponse>(url, null, { headers }).pipe(
      tap((res) => {
        this.setAccessToken(res.access_token);
      }),
    );
  }

  getRefreshToken(): string | null {
    return localStorage.getItem(this.REFRESH_TOKEN_KEY) ||
      sessionStorage.getItem(this.REFRESH_TOKEN_KEY);
  }

  private setRefreshToken(refreshToken: string, remember = false) {
    if (remember) {
      localStorage.setItem(this.REFRESH_TOKEN_KEY, refreshToken);
    } else {
      sessionStorage.setItem(this.REFRESH_TOKEN_KEY, refreshToken);
    }
  }

  private setAccessToken(accessToken: string) {
    this.accessToken = accessToken;
    this.loggedIn.set(true);
  }

  private setUserInfo(username: string) {
    this.currentUser = username;
  }

  private logout() {
    localStorage.removeItem(this.REFRESH_TOKEN_KEY);
    sessionStorage.removeItem(this.REFRESH_TOKEN_KEY);
    this.loggedIn.set(false);
  }
}
