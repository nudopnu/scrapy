import { Component, Input, LOCALE_ID } from "@angular/core";
import { AgentResponse } from "../../models/responses";
import { DATE_PIPE_DEFAULT_OPTIONS, DatePipe, JsonPipe } from "@angular/common";

@Component({
  selector: "fs-agent",
  standalone: true,
  imports: [JsonPipe, DatePipe],
  templateUrl: "./agent.component.html",
  styleUrl: "./agent.component.css",
})
export class AgentComponent {
  @Input()
  agent!: AgentResponse;
}
