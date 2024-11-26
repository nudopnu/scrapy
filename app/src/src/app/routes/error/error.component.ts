import { AfterContentInit, Component, Input, input } from "@angular/core";

@Component({
  selector: "fs-error",
  standalone: true,
  imports: [],
  templateUrl: "./error.component.html",
})
export class ErrorComponent implements AfterContentInit {
  @Input()
  status = "";
  statusText = "";

  ngAfterContentInit(): void {
    this.statusText = {
      500: "Sorry, there is an error on server.",
      0: "The server seems to be offline. Please try again later.",
    }[this.status] ?? "";
  }
}
