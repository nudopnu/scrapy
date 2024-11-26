import { Component, inject, Input } from "@angular/core";
import {
  FormBuilder,
  FormControl,
  ReactiveFormsModule,
  Validators,
} from "@angular/forms";
import { Router } from "@angular/router";
import { catchError, of } from "rxjs";
import { InputComponent } from "../../components/input/input.component";
import { ApiService } from "../../services/api.service";
import { AuthService } from "../../services/auth.service";

@Component({
    selector: "fs-login",
    imports: [ReactiveFormsModule, InputComponent],
    templateUrl: "./login.component.html"
})
export class LoginComponent {
  @Input()
  redirectTo?: string;
  form = inject(FormBuilder).group({
    username: ["", [
      Validators.required,
    ]],
    password: ["", [
      Validators.required,
    ]],
    remember: new FormControl(false),
  });
  isSubmitting = false;
  apiService = inject(ApiService);
  authService = inject(AuthService);
  router = inject(Router);

  serverError = "";

  onClickSubmit() {
    const username = this.form.value.username;
    const password = this.form.value.password;
    const remember = !!this.form.value.remember;
    if (this.form.errors || !username || !password) {
      this.form.markAllAsTouched();
      return;
    }
    this.submit(username, password, remember);
  }

  private submit(username: string, password: string, remember: boolean) {
    this.isSubmitting = true;
    this.serverError = "";
    this.authService.login(username, password, remember).pipe(
      catchError((err) => {
        this.isSubmitting = false;
        if (err.status === 0) {
          this.serverError = "The server seems to be offline";
        }
        if (err.message) {
          this.serverError = err.message;
        }
        return of();
      }),
    ).subscribe((res) => {
      this.isSubmitting = false;
      if (this.redirectTo) {
        this.router.navigate([this.redirectTo]);
      }
    });
  }
}
