import { inject, Injectable, signal } from "@angular/core";
import { Router } from "@angular/router";
import { tap } from "rxjs";
import { ApiService } from "./api.service";

@Injectable({
  providedIn: "root",
})
export class AuthService {
  private REFRESH_TOKEN_KEY = "refresh_token";
  accessToken?: string;
  currentUser?: string;
  isLoggedIn = signal(false);
  router = inject(Router);
  apiService = inject(ApiService);

  login(username: string, password: string, remember: boolean) {
    return this.apiService.login(username, password).pipe(
      tap((res) => {
        this.setUserInfo(res.username);
        this.setAccessToken(res.access_token);
        this.setRefreshToken(res.refresh_token, remember);
      }),
    );
  }

  refreshToken(sessionToken: string) {
    return this.apiService.refresh(sessionToken).pipe(
      tap((res) => {
        this.setUserInfo(this.parseJWT(res.access_token).username);
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
    this.isLoggedIn.set(true);
  }

  private setUserInfo(username: string) {
    this.currentUser = username;
  }

  logout() {
    localStorage.removeItem(this.REFRESH_TOKEN_KEY);
    sessionStorage.removeItem(this.REFRESH_TOKEN_KEY);
    this.accessToken = undefined;
    this.isLoggedIn.set(false);
    this.router.navigate(["/"]);
    // TODO: Revoke refresh token
  }

  parseJWT(token: string) {
    const base64Url = token.split(".")[1];
    const base64 = base64Url.replace(/-/g, "+").replace(/_/g, "/");
    const jsonPayload = decodeURIComponent(
      atob(base64)
        .split("")
        .map((c) => "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2))
        .join(""),
    );

    return JSON.parse(jsonPayload);
  }
}
