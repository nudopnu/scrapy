import { Component, inject } from "@angular/core";
import { FormBuilder, ReactiveFormsModule, Validators } from "@angular/forms";
import { InputComponent } from "../input/input.component";
import { SelectComponent } from "../select/select.component";

@Component({
  selector: "fs-create-agent",
  standalone: true,
  imports: [InputComponent, ReactiveFormsModule, SelectComponent],
  templateUrl: "./create-agent.component.html",
  styleUrl: "./create-agent.component.css",
})
export class CreateAgentComponent {
  form = inject(FormBuilder).group({
    name: ["", [Validators.required, Validators.minLength(3)]],
    searchTerm: ["", [Validators.required, Validators.minLength(3)]],
    postalCode: ["", [Validators.required, Validators.minLength(3)]],
    distance: [0, [Validators.required]],
  });
  isSubmitting = false;

  onClickSubmit() {
    Object.values(this.form.controls).forEach((control) =>
      control.updateValueAndValidity()
    );
    const distance = this.form.value.distance;
    const name = this.form.value.name;
    const postalCode = this.form.value.postalCode;
    const searchTerm = this.form.value.searchTerm;
    if (!this.form.valid || !distance || !name || !postalCode || !searchTerm) {
      return;
    }
    console.log(this.form.value);
  }
}
