import { Injectable, signal } from "@angular/core";

@Injectable({
  providedIn: "root",
})
export class AuthService {
  accessToken?: string;
  currentUser?: string;
  loggedIn = signal(false);

  getRefreshToken(): string | null {
    return localStorage.getItem("refresh_token") ||
      sessionStorage.getItem("refresh_token");
  }

  setRefreshToken(refreshToken: string, remember = false) {
    if (remember) {
      localStorage.setItem("refresh_token", refreshToken);
    } else {
      sessionStorage.setItem("refresh_token", refreshToken);
    }
  }

  setAccessToken(accessToken: string) {
    this.accessToken = accessToken;
    this.loggedIn.set(true);
  }

  setUserInfo(username: string) {
    this.currentUser = username;
  }

  logout() {
    localStorage.removeItem("refresh_token");
    sessionStorage.removeItem("refresh_token");
    this.loggedIn.set(false);
  }
}
