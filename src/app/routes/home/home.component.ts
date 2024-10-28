import { JsonPipe } from "@angular/common";
import { Component, inject, OnInit } from "@angular/core";
import { toSignal } from "@angular/core/rxjs-interop";
import { Router } from "@angular/router";
import { ApiService } from "../../services/api.service";

@Component({
  selector: "fs-home",
  standalone: true,
  imports: [JsonPipe],
  templateUrl: "./home.component.html",
  styleUrl: "./home.component.css",
})
export class HomeComponent implements OnInit {
  router = inject(Router);
  agents: any;

  constructor(private apiService: ApiService){
    this.agents = toSignal(this.apiService.getAgents());
  }

  ngOnInit(): void {
    
    // console.log(this.router.navigate(["login"]));
  }
}
