import { Component, Input, LOCALE_ID } from "@angular/core";
import { AgentResponse } from "../../models/responses";
import { DatePipe, JsonPipe } from "@angular/common";

@Component({
  selector: "fs-agent",
  standalone: true,
  imports: [JsonPipe, DatePipe],
  providers: [
    { provide: LOCALE_ID, useValue: "de-DE" },
  ],
  templateUrl: "./agent.component.html",
  styleUrl: "./agent.component.css",
})
export class AgentComponent {
  @Input()
  agent!: AgentResponse;
}
