import { Component, inject, OnInit } from "@angular/core";
import { Router, RouterOutlet } from "@angular/router";
import { initFlowbite } from "flowbite";
import { retry } from "rxjs";
import { BottomNavigationComponent } from "./components/bottom-navigation/bottom-navigation.component";
import { NavigationComponent } from "./components/navigation/navigation.component";
import { ApiService } from "./services/api.service";
import { AuthService } from "./services/auth.service";

@Component({
    selector: "fs-root",
    imports: [RouterOutlet, NavigationComponent, BottomNavigationComponent],
    templateUrl: "./app.component.html",
    styleUrl: "./app.component.css"
})
export class AppComponent implements OnInit {
  title = "flowbite-sample";
  apiService = inject(ApiService);
  router = inject(Router);
  auth = inject(AuthService);

  constructor() {}

  ngOnInit(): void {
    initFlowbite();
  }
}
