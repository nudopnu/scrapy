import { Component, inject } from "@angular/core";
import { FormBuilder, ReactiveFormsModule, Validators } from "@angular/forms";
import { CreateSearchAgent } from "../../models/requests";
import { ApiService } from "../../services/api.service";
import { InputComponent } from "../input/input.component";
import { SelectComponent } from "../select/select.component";

@Component({
  selector: "fs-create-agent",
  standalone: true,
  imports: [InputComponent, ReactiveFormsModule, SelectComponent],
  templateUrl: "./create-agent.component.html",
})
export class CreateAgentComponent {
  form = inject(FormBuilder).group({
    name: ["", [Validators.required, Validators.minLength(3)]],
    searchTerm: ["", [Validators.required, Validators.minLength(3)]],
    postalCode: ["", [Validators.required, Validators.minLength(3)]],
    distance: [0, [Validators.required]],
  });
  apiService = inject(ApiService);
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
    this.submit({
      distance,
      keyword: searchTerm,
      name,
      postal_code: postalCode,
    });
  }

  submit(agent: CreateSearchAgent) {
    this.apiService.createAgent(agent).subscribe((res) => console.log(res));
  }
}
