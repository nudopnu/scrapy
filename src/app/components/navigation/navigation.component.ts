import { Component, inject, OnChanges, SimpleChanges } from "@angular/core";
import { RouterLink } from "@angular/router";
import { initFlowbite } from "flowbite";
import { AuthService } from "../../services/auth.service";
import { AvatarComponent } from "../avatar/avatar.component";

@Component({
  selector: "fs-navigation",
  standalone: true,
  imports: [RouterLink, AvatarComponent],
  templateUrl: "./navigation.component.html",
})
export class NavigationComponent implements OnChanges {
  auth = inject(AuthService);

  ngOnChanges(changes: SimpleChanges): void {
    initFlowbite();
  }
}
