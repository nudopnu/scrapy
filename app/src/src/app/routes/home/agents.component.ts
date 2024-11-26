import { Component, inject, OnInit } from "@angular/core";
import { Router, RouterModule } from "@angular/router";
import { AgentCardComponent } from "../../components/agent-card/agent.component";
import { MockAgents as MOCK_AGENTS } from "../../mock/agents.mock";
import { AgentResponse } from "../../models/responses";
import { ApiService } from "../../services/api.service";
import { environment } from "../../../environments/environment";

@Component({
  selector: "fs-agents",
  imports: [AgentCardComponent, RouterModule],
  templateUrl: "./agents.component.html"
})
export class AgentsComponent implements OnInit {
  router = inject(Router);
  agents: AgentResponse[] = [];

  constructor(private apiService: ApiService) { }

  ngOnInit(): void {
    if (environment.mock) {
      this.mockAgents();
      return;
    }
    this.apiService.getAgents().subscribe((agents) => {
      this.agents = agents;
      console.log(this.agents);
    });
  }

  onClick(agent: AgentResponse) {
    console.log(agent);
    this.router.navigate([`/agents/${agent.id}/ads`]);
  }

  private mockAgents() {
    this.agents = MOCK_AGENTS;
  }
}
