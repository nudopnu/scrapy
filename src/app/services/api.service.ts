import { HttpClient, HttpErrorResponse } from "@angular/common/http";
import { inject, Injectable } from "@angular/core";
import { Router } from "@angular/router";
import { catchError, throwError } from "rxjs";
import { CreateSearchAgent } from "../models/requests";
import {
  AgentResponse,
  LoginResponse,
  RefreshResponse,
} from "../models/responses";

@Injectable({
  providedIn: "root",
})
export class ApiService {
  BASE_URL = "http://localhost:8080/api/v1";
  router = inject(Router);
  http = inject(HttpClient);

  checkHealth() {
    const url = `${this.BASE_URL}/healthz`;
    return this.http.get<string>(url, { responseType: "text" as "json" })
      .pipe(catchError(this.handleError.bind(this)));
  }

  login(username: string, password: string) {
    const url = `${this.BASE_URL}/login`;
    return this.http.post<LoginResponse>(url, { username, password })
      .pipe(catchError(this.handleError.bind(this)));
  }

  refresh(sessionToken: string) {
    const url = `${this.BASE_URL}/refresh`;
    const headers = { "Authorization": `Bearer ${sessionToken}` };
    return this.http.post<RefreshResponse>(url, null, { headers })
      .pipe(catchError(this.handleError.bind(this)));
  }

  getAgents() {
    const url = `${this.BASE_URL}/agents`;
    return this.http.get<AgentResponse[]>(url)
      .pipe(catchError(this.handleError.bind(this)));
  }

  createAgent(createAgentRequest: CreateSearchAgent) {
    const url = `${this.BASE_URL}/agents`;
    return this.http.post<AgentResponse>(url, createAgentRequest)
      .pipe(catchError(this.handleError.bind(this)));
  }

  private handleError(error: HttpErrorResponse) {
    if (error.error) {
      return throwError(() => error.error);
    }
    switch (error.status) {
      case 0:
        this.router.navigate(["error"], {
          queryParams: {
            status: 0,
            redirectTo: `/${(this.router as any).location._basePath}`,
          },
        });
        return throwError(() =>
          new Error("Unauthorized. Redirecting to login page.")
        );
      case 401:
        this.router.navigate(["login"], {
          queryParams: {
            redirectTo: `/${(this.router as any).location._basePath}`,
          },
        });
        return throwError(() =>
          new Error("Unauthorized. Redirecting to login page.")
        );
      default:
        break;
    }

    return throwError(() => error);
  }
}
