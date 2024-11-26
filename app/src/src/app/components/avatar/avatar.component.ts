import { AfterViewInit, Component, inject } from "@angular/core";
import { initFlowbite } from "flowbite";
import { AuthService } from "../../services/auth.service";

@Component({
    selector: "fs-avatar",
    imports: [],
    templateUrl: "./avatar.component.html"
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
