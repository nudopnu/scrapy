import { Routes } from "@angular/router";
import { ErrorComponent } from "./routes/error/error.component";
import { HomeComponent } from "./routes/home/home.component";
import { LoginComponent } from "./routes/login/login.component";

export const routes: Routes = [
    { path: "", component: HomeComponent },
    { path: "login", component: LoginComponent },
    { path: "error", component: ErrorComponent },
    { path: "**", redirectTo: "/error?status=404" },
];
