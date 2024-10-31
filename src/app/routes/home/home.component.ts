import { JsonPipe } from "@angular/common";
import { Component, inject, OnInit } from "@angular/core";
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
  agents = [];

  constructor(private apiService: ApiService) {}

  ngOnInit(): void {
    this.apiService.getAgents().subscribe((res) => {
      console.log(res);
    });
  }
}
