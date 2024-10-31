import { Component, inject } from "@angular/core";
import { ApiService } from "../../services/api.service";
import { CreateAgentComponent } from "../../components/create-agent/create-agent.component";

@Component({
  selector: "fs-agents",
  standalone: true,
  imports: [CreateAgentComponent],
  templateUrl: "./agents.component.html",
  styleUrl: "./agents.component.css",
})
export class AgentsComponent {
  apiService = inject(ApiService);
  agents = [];
}
