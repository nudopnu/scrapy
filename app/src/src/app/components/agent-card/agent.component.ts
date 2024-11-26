import { DatePipe } from "@angular/common";
import { Component, Input } from "@angular/core";
import { AgentResponse } from "../../models/responses";

@Component({
    selector: "fs-agent",
    imports: [DatePipe],
    templateUrl: "./agent-card.component.html"
})
export class AgentCardComponent {
  @Input()
  agent!: AgentResponse;
}
