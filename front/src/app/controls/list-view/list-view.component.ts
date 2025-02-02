import { Component, forwardRef, Input } from '@angular/core';
import { ControlValueAccessor, NG_VALUE_ACCESSOR } from '@angular/forms';
import { Field } from '../data-table/data-table.model';
import { ValueChangedEvent } from '../item-selector/item-selector.model';
import { ApiService } from '../../app.service';
import { IItem } from '../../app.model';

@Component({
  selector: 'app-list-view',
  standalone: false,
  templateUrl: './list-view.component.html',
  styleUrl: './list-view.component.scss',
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => ListViewComponent),
      multi: true,
    },
  ],
})
export class ListViewComponent<T extends IItem>
  implements ControlValueAccessor
{
  @Input()
  Disabled: boolean = false;

  @Input()
  Value: T[] = [];

  @Input({ required: true })
  TitleField?: Field<T, string>;

  @Input({ required: true })
  TPath: string = '';

  @Input()
  ReadOnly: boolean = false;

  protected selectedItem: T | null = null;
  protected onChanged?: (value?: T[]) => void;
  protected onTouched?: () => void;

  constructor(private readonly appService: ApiService<T>) {}

  getSortedValue() {
    return this.Value.sort((a, b) => {
      const ac = this.TitleField?.Content(a) ?? '';
      const bc = this.TitleField?.Content(b) ?? '';
      if (ac < bc) return -1;
      if (ac > bc) return 1;
      return 0;
    });
  }

  registerOnChange(fn: any): void {
    this.onChanged = fn;
  }

  registerOnTouched(fn: any): void {
    this.onTouched = fn;
  }

  setDisabledState(isDisabled: boolean): void {
    this.Disabled = isDisabled;
  }

  writeValue(obj: any): void {
    this.Value = obj;
  }

  selectItem($event: ValueChangedEvent<number>) {
    if ($event.newValue) {
      this.appService.Get(this.TPath, $event.newValue).subscribe((value) => {
        this.selectedItem = value;
      });
    }
  }

  addItem() {
    if (
      this.selectedItem &&
      this.Value.findIndex((i) => i.id == this.selectedItem?.id) == -1
    ) {
      this.Value.push(this.selectedItem);
      this.onChanged ? this.onChanged(this.Value) : undefined;
    }
  }

  deleteItem(row: T) {
    const i = this.Value.indexOf(row);
    if (i != -1) {
      this.Value.splice(i, 1);
      this.onChanged ? this.onChanged(this.Value) : undefined;
    }
  }

  selectorTouched() {
    this.onTouched ? this.onTouched() : undefined;
  }
}
