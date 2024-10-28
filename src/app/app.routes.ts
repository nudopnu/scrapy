import { Routes } from "@angular/router";
import { HomeComponent } from "./routes/home/home.component";
import { LoginComponent } from "./routes/login/login.component";

export const routes: Routes = [
    { path: "", component: HomeComponent },
    { path: "login", component: LoginComponent },
];
