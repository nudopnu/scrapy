import { JsonPipe } from "@angular/common";
import { Component, inject, OnInit } from "@angular/core";
import { Router } from "@angular/router";
import { AgentCardComponent } from "../../components/agent-card/agent.component";
import { MockAgents as MOCK_AGENTS } from "../../mock/agents.mock";
import { AgentResponse } from "../../models/responses";
import { ApiService } from "../../services/api.service";

@Component({
  selector: "fs-agents",
  standalone: true,
  imports: [JsonPipe, AgentCardComponent],
  templateUrl: "./agents.component.html",
})
export class AgentsComponent implements OnInit {
  router = inject(Router);
  agents: AgentResponse[] = [];

  constructor(private apiService: ApiService) {}

  ngOnInit(): void {
    this.apiService.getAgents().subscribe((agents) => {
      this.agents = agents;
      console.log(this.agents);
    });
    // this.mockAgents();
  }

  onClick(agent: AgentResponse) {
    console.log(agent);
    this.router.navigate([`/agents/${agent.id}/ads`]);
  }

  private mockAgents() {
    this.agents = MOCK_AGENTS;
  }
}
