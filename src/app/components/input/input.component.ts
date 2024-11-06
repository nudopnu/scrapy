import { JsonPipe } from "@angular/common";
import { Component, effect, Input, signal } from "@angular/core";
import {
  AbstractControl,
  ControlValueAccessor,
  FormsModule,
  NG_VALIDATORS,
  NG_VALUE_ACCESSOR,
  ValidationErrors,
  Validator,
} from "@angular/forms";

@Component({
  selector: "fs-input",
  standalone: true,
  imports: [FormsModule, JsonPipe],
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      multi: true,
      useExisting: InputComponent,
    },
    {
      provide: NG_VALIDATORS,
      multi: true,
      useExisting: InputComponent,
    },
  ],
  templateUrl: "./input.component.html",
})
export class InputComponent implements ControlValueAccessor, Validator {
  @Input()
  type = "text";
  @Input()
  label = "";
  @Input()
  placeHolder = "";
  @Input()
  redirectTo = "";
  @Input()
  hasSibling: "none" | "left" | "right" | "both" = "none";

  messages: Map<string, string> = new Map();

  state: "default" | "success" | "invalid" | "disabled" = "default";
  initialized = false;
  errors = signal<any[]>([]);

  private _value: string = "";
  set value(val: string) {
    if (val !== this._value) {
      this._value = val;
      // this.onChange(val); // Update the form control value
    }
  }
  get value(): string {
    return this._value;
  }

  classList = {
    input: {
      default:
        "bg-gray-50 border border-gray-300 text-gray-900 text-sm focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500",
      disabled:
        "bg-gray-50 border border-gray-300 text-gray-900 text-sm focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500",
      success:
        "bg-green-50 border border-green-500 text-green-900 dark:text-green-400 placeholder-green-700 dark:placeholder-green-500 text-sm focus:ring-green-500 focus:border-green-500 block w-full p-2.5 dark:bg-gray-700 dark:border-green-500",
      invalid:
        "bg-red-50 border border-red-500 text-red-900 placeholder-red-700 text-sm focus:ring-red-500 dark:bg-gray-700 focus:border-red-500 block w-full p-2.5 dark:text-red-500 dark:placeholder-red-500 dark:border-red-500",
    },
  };
  borderClass = {
    "none": "rounded-lg",
    "left": "rounded-e-lg",
    "right": "rounded-s-lg",
    "both": "",
  };

  onChange = (ch: any) => {};
  onTouched = () => {};

  constructor() {
    effect(() => {
      if (this.errors().length > 0) {
        this.state = "invalid";
      } else {
        this.state = "default";
      }
    });
  }

  writeValue(v: any): void {
    this.value = v;
    console.log("setting", this.label, v);
  }

  registerOnChange(fn: any): void {
    this.onChange = fn;
  }

  registerOnTouched(fn: any): void {
    this.onTouched = fn;
  }

  setDisabledState?(isDisabled: boolean): void {
    this.state = isDisabled ? "disabled" : "default";
  }

  markAsTouched() {
    this.onTouched();
  }

  validate(control: AbstractControl): ValidationErrors | null {
    setTimeout(() => {
      let errorMesages: string[] = [];
      if (control.errors && this.initialized) {
        console.log(control.errors);
        errorMesages = Object
          .entries(control.errors)
          .map(([err, errValue]) => {
            switch (err) {
              case "required":
                return `${this.label} is required`;
              case "minlength":
                return `${this.label} must be at least ${errValue.requiredLength} characters long`;
              default:
                return err;
            }
          });
      }
      this.errors.set(errorMesages);
      this.initialized = true;
    });
    return null;
  }

  onBlur() {
    this.onChange(this.value);
    this.markAsTouched();
  }

  onEnterKey() {
    this.onChange(this.value);
    this.markAsTouched();
  }
}
