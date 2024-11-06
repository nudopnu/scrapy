import { DatePipe, JsonPipe } from "@angular/common";
import { Component, Input } from "@angular/core";
import { AgentResponse } from "../../models/responses";

@Component({
  selector: "fs-agent",
  standalone: true,
  imports: [JsonPipe, DatePipe],
  templateUrl: "./agent-card.component.html",
})
export class AgentCardComponent {
  @Input()
  agent!: AgentResponse;
}
