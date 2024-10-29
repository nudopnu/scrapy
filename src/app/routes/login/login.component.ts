import { Component, inject } from "@angular/core";
import {
  FormBuilder,
  FormControl,
  ReactiveFormsModule,
  Validators,
} from "@angular/forms";
import { catchError, of } from "rxjs";
import { InputComponent } from "../../components/input/input.component";
import { ApiService } from "../../services/api.service";

@Component({
  selector: "fs-login",
  standalone: true,
  imports: [ReactiveFormsModule, InputComponent],
  templateUrl: "./login.component.html",
  styleUrl: "./login.component.css",
})
export class LoginComponent {
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

  serverError = "";

  onClickSubmit() {
    console.log(this.form.value);
    const username = this.form.value.username;
    const password = this.form.value.password;
    if (this.form.errors || !username || !password) {
      return;
    }
    this.submit(username, password);
  }

  private submit(username: string, password: string) {
    this.isSubmitting = true;
    this.serverError = "";
    this.apiService.login(username, password).pipe(
      catchError((err) => {
        this.isSubmitting = false;
        if (err.message) {
          this.serverError = err.message;
        }
        return of();
      }),
    ).subscribe((res) => {
      this.isSubmitting = false;
    });
  }
  
}
