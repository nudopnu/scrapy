import { Component, inject } from '@angular/core';
import { RouterModule } from '@angular/router';
import { AuthService } from '../../../services/auth.service';
import { AvatarComponent } from "../../avatar/avatar.component";

@Component({
  selector: 'fs-top-navigation',
  imports: [RouterModule, AvatarComponent],
  templateUrl: './top-navigation.component.html',
})
export class TopNavigationComponent {

  auth = inject(AuthService);

}
