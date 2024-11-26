import { Routes } from "@angular/router";
import { ErrorComponent } from "./routes/error/error.component";
import { AgentsComponent } from "./routes/home/agents.component";
import { LoginComponent } from "./routes/login/login.component";
import { AdsComponent } from "./routes/ads/ads.component";
import { CreateAgentComponent } from "./components/create-agent/create-agent.component";

export const routes: Routes = [
    { path: "login", component: LoginComponent },
    { path: "agents", component: AgentsComponent },
    { path: "agents/create", component: CreateAgentComponent },
    { path: "agents/:agentId/ads", component: AdsComponent },
    { path: "error", component: ErrorComponent },
    // { path: "**", redirectTo: "/error?status=404" },
];
