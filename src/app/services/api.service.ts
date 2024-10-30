import { HttpClient } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { LoginResponse } from "../models/responses";

@Injectable({
  providedIn: "root",
})
export class ApiService {
  BASE_URL = "http://localhost:8080/api/v1";

  constructor(private http: HttpClient) {}

  getAgents() {
    const url = `${this.BASE_URL}/agents`;
    return this.http.get(url);
  }

  isAlive() {
    const url = `${this.BASE_URL}/healthz`;
    return this.http.get(url);
  }

}
