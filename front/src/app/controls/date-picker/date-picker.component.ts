import { Component, forwardRef } from '@angular/core';
import { ControlValueAccessor, NG_VALUE_ACCESSOR } from '@angular/forms';
import { NgbDateStruct } from '@ng-bootstrap/ng-bootstrap';

@Component({
  selector: 'app-date-picker',
  standalone: false,
  templateUrl: './date-picker.component.html',
  styleUrl: './date-picker.component.scss',
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => DatePickerComponent),
      multi: true,
    },
  ],
})
export class DatePickerComponent implements ControlValueAccessor {
  value: string = '';
  disabled: boolean = false;
  private onChange = (date: string) => {};
  private onTouched = () => {};

  writeValue(date: string): void {
    this.value = date.split('T')[0];
  }

  registerOnChange(fn: (date: string) => void): void {
    this.onChange = fn;
  }

  registerOnTouched(fn: () => void): void {
    this.onTouched = fn;
  }

  setDisabledState?(isDisabled: boolean): void {
    this.disabled = isDisabled;
  }

  selectorTouched() {
    this.onTouched ? this.onTouched() : undefined;
  }

  valueChanged(date: NgbDateStruct) {
    const year = date.year;
    const month = `${date.month < 10 ? '0' : ''}${date.month}`;
    const day = `${date.day < 10 ? '0' : ''}${date.day}`;
    this.onChange
      ? this.onChange(`${year}-${month}-${day}T00:00:00Z`)
      : undefined;
  }
}
