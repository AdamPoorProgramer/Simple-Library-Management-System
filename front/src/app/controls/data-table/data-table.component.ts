import { Component, Input } from '@angular/core';
import { Field } from './data-table.model';

@Component({
  selector: 'app-data-table',
  standalone: false,
  templateUrl: './data-table.component.html',
  styleUrl: './data-table.component.scss',
})
export class DataTableComponent<T> {
  @Input({ required: true })
  Fields?: Field<T>[];

  @Input({ required: true })
  Data?: T[];
}
