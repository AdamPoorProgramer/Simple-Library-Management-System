import {
  Component,
  ElementRef,
  EventEmitter,
  forwardRef,
  Input,
  OnInit,
  Output,
  ViewChild,
} from '@angular/core';
import { ControlValueAccessor, NG_VALUE_ACCESSOR } from '@angular/forms';
import { IItem } from '../../app.model';
import { Field } from '../data-table/data-table.model';
import { ValueChangedEvent } from './item-selector.model';
import { ApiService } from '../../app.service';

@Component({
  selector: 'app-item-selector',
  standalone: false,
  templateUrl: './item-selector.component.html',
  styleUrl: './item-selector.component.scss',
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => ItemSelectorComponent),
      multi: true,
    },
  ],
})
export class ItemSelectorComponent<T extends IItem>
  implements ControlValueAccessor, OnInit
{
  @ViewChild('Selector')
  Selector?: ElementRef<HTMLSelectElement>;

  @Input()
  Disabled: boolean = false;

  @Input()
  Value?: number;

  @Input({ required: true })
  TitleField?: Field<T, string>;

  getContent(row: T) {
    return this.TitleField?.Content(row);
  }

  @Input({ required: true })
  TPath: string = '';

  ngOnInit() {
    this.appService.GetAll(this.TPath).subscribe((data) => {
      if (data) this.Data = data;
    });
  }

  @Output()
  ValueChanged: EventEmitter<ValueChangedEvent<number>> = new EventEmitter();

  protected Data: T[] = [];
  protected onChanged?: (value?: number) => void;
  protected onTouched?: () => void;

  constructor(private readonly appService: ApiService<T>) {}

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
    if (this.Selector)
      this.Selector.nativeElement.selectedIndex =
        this.Data.findIndex((row) => row.id == obj) + 1;
  }

  valueChanged() {
    const index = this.Selector?.nativeElement.selectedIndex;
    const val = index ? this.Data[index - 1].id : undefined;
    if (val != this.Value) {
      this.ValueChanged.emit(new ValueChangedEvent(this.Value, val));
      this.Value = val;
      this.onChanged ? this.onChanged(this.Value) : undefined;
    }
  }

  selectorTouched() {
    this.onTouched ? this.onTouched() : undefined;
  }
}
