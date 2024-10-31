import { HttpClient } from "@angular/common/http";
import { inject, Injectable, signal } from "@angular/core";
import { Router } from "@angular/router";
import { catchError, of, tap, throwError } from "rxjs";
import { LoginResponse, RefreshResponse } from "../models/responses";
import { ApiService } from "./api.service";

@Injectable({
  providedIn: "root",
})
export class AuthService {
  private REFRESH_TOKEN_KEY = "refresh_token";
  accessToken?: string;
  currentUser?: string;
  isLoggedIn = signal(false);
  isAlive = signal(false);
  http = inject(HttpClient);
  router = inject(Router);
  BASE_URL = inject(ApiService).BASE_URL;

  checkHealth() {
    const url = `${this.BASE_URL}/healthz`;
    return this.http.get(url)
      .pipe(
        catchError((err) => {
          console.log(err);
          this.router.navigate(["error"], { queryParams: { status: 500 } });
          this.isAlive.set(false);
          return throwError(() => err);
        }),
        tap(() => {
          this.isAlive.set(true);
        }),
      );
  }

  login(username: string, password: string, remember: boolean) {
    const url = `${this.BASE_URL}/login`;
    return this.http.post<LoginResponse>(url, { username, password }).pipe(
      tap((res) => {
        this.setUserInfo(res.username);
        this.setAccessToken(res.access_token);
        this.setRefreshToken(res.refresh_token, remember);
      }),
    );
  }

  refreshToken(sessionToken: string) {
    const url = `${this.BASE_URL}/refresh`;
    const headers = { "Authorization": `Bearer ${sessionToken}` };
    return this.http.post<RefreshResponse>(url, null, { headers }).pipe(
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
    this.isLoggedIn.set(false);
    // TODO: Revoke refresh token
  }

  private parseJWT(token: string) {
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
