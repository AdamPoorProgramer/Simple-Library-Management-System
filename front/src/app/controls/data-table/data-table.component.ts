import {
  AfterViewInit,
  Component,
  ElementRef,
  EventEmitter,
  Input,
  Output,
  ViewChild,
} from '@angular/core';
import { Field } from './data-table.model';
import { DialogComponent } from '../dialog/dialog.component';

@Component({
  selector: 'app-data-table',
  standalone: false,
  templateUrl: './data-table.component.html',
  styleUrl: './data-table.component.scss',
})
export class DataTableComponent<T> {
  private fields?: Field<T, any>[];
  IDField?: Field<T, number>;
  @ViewChild('Dialog') Dialog?: DialogComponent;

  @Input({ required: true })
  set Fields(value: Field<T, any>[]) {
    this.fields = value;
    if (value.length > 0) {
      this.IDField = value[0] as Field<T, number>;
    } else {
      this.IDField = undefined;
    }
  }

  get Fields(): Field<T, any>[] {
    return this.fields ?? [];
  }

  @Input({ required: true })
  Data?: T[];

  @Input({ required: true })
  Title?: string;

  @Output()
  Edit = new EventEmitter<number | null>();

  @Output()
  Delete = new EventEmitter<number | null>();

  @Output()
  Reload = new EventEmitter();

  protected readonly Math = Math;

  async deleteItem(row: T) {
    const confirmed = await this.Dialog?.open(this.IDField?.Content(row) ?? 0);
    if (confirmed) {
      this.Delete.emit(this.IDField?.Content(row));
    }
  }
}
