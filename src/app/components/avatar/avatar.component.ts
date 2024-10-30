import { AfterViewChecked, AfterViewInit, Component, inject } from "@angular/core";
import { AuthService } from "../../services/auth.service";
import { initFlowbite } from "flowbite";

@Component({
  selector: "fs-avatar",
  standalone: true,
  imports: [],
  templateUrl: "./avatar.component.html",
  styleUrl: "./avatar.component.css",
})
export class AvatarComponent  implements AfterViewInit{
  auth = inject(AuthService);

  onClickSignOut() {
    this.auth.logout();
  }

  ngAfterViewInit(): void {
    initFlowbite();
  }
}
