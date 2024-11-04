import { JsonPipe } from "@angular/common";
import { Component, inject, OnInit } from "@angular/core";
import { Router } from "@angular/router";
import { ApiService } from "../../services/api.service";
import { AgentResponse } from "../../models/responses";
import { AgentComponent } from "../../components/agent/agent.component";

@Component({
  selector: "fs-home",
  standalone: true,
  imports: [JsonPipe, AgentComponent],
  templateUrl: "./home.component.html",
  styleUrl: "./home.component.css",
})
export class HomeComponent implements OnInit {
  router = inject(Router);
  agents: AgentResponse[] = [];

  constructor(private apiService: ApiService) {}

  ngOnInit(): void {
    this.apiService.getAgents().subscribe((agents) => {
      this.agents = agents;
    });
    // this.mockAgents();
  }

  private mockAgents() {
    this.agents = [
      {
        "id": 1,
        "name": "Ebikes",
        "user_id": 1,
        "last_fetched_at": {
          "Time": new Date(Date.parse("0001-01-01T00:00:00Z")),
          "Valid": false,
        },
        "created_at": new Date(Date.parse("2024-10-31T16:11:44.119303Z")),
        "updated_at": new Date(Date.parse("2024-10-31T16:11:44.119303Z")),
      },
      {
        "id": 2,
        "name": "Pokemon Trier",
        "user_id": 1,
        "last_fetched_at": {
          "Time": new Date(Date.parse("0001-01-01T00:00:00Z")),
          "Valid": false,
        },
        "created_at": new Date(Date.parse("2024-10-31T16:27:26.499987Z")),
        "updated_at": new Date(Date.parse("2024-10-31T16:27:26.499987Z")),
      },
      {
        "id": 3,
        "name": "Monitor Trier",
        "user_id": 1,
        "last_fetched_at": {
          "Time": new Date(Date.parse("0001-01-01T00:00:00Z")),
          "Valid": false,
        },
        "created_at": new Date(Date.parse("2024-10-31T16:32:14.664766Z")),
        "updated_at": new Date(Date.parse("2024-10-31T16:32:14.664766Z")),
      },
    ];
  }
}
