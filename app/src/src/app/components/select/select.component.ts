import { Component, forwardRef, Input } from "@angular/core";
import {
  ControlValueAccessor,
  FormsModule,
  NG_VALUE_ACCESSOR,
} from "@angular/forms";

@Component({
    selector: "fs-select",
    imports: [FormsModule],
    providers: [
        {
            provide: NG_VALUE_ACCESSOR,
            useExisting: forwardRef(() => SelectComponent),
            multi: true,
        },
    ],
    templateUrl: "./select.component.html"
})
export class SelectComponent implements ControlValueAccessor {
  @Input()
  entries!: Record<string, any>;

  idDisabled = false;
  onChange = (v: any) => {};
  onTouched = () => {};

  _value: number = 0;

  public set value(v: number) {
    if (v !== this._value) { // To prevent redundant updates
      this._value = v;
      this.onChange(v);
    }
  }

  public get value(): number {
    return this._value;
  }

  writeValue(value: any): void {
    this.value = value;
  }

  registerOnChange(fn: any): void {
    this.onChange = fn;
  }

  registerOnTouched(fn: any): void {
    this.onTouched = fn;
  }

  setDisabledState?(isDisabled: boolean): void {
    this.idDisabled = isDisabled;
  }
}
