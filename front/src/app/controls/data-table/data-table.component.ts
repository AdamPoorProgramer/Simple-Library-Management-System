import { Component, Input, OnInit, ViewChild } from '@angular/core';
import { Field, FormControls } from './data-table.model';
import { DialogComponent } from '../dialog/dialog.component';
import { ApiService } from '../../app.service';
import { IItem } from '../../app.model';
import { FormBuilder, FormGroup } from '@angular/forms';
import { BookComponent } from '../../pages/book/book.component';
import { CategoryComponent } from '../../pages/category/category.component';

@Component({
  selector: 'app-data-table',
  standalone: false,
  templateUrl: './data-table.component.html',
  styleUrl: './data-table.component.scss',
})
export class DataTableComponent<T extends IItem> implements OnInit {
  @ViewChild('DeleteDialog')
  DeleteDialog?: DialogComponent;

  @ViewChild('EditDialog')
  EditDialog?: DialogComponent;

  @ViewChild('ViewDialog')
  ViewDialog?: DialogComponent;

  @Input({ required: true })
  set Fields(value: Field<T, any>[]) {
    this._fields = value;
    this.width = 100 / (value.filter((f) => f.Visible).length + 1);
  }

  get Fields(): Field<T, any>[] {
    return this._fields;
  }

  @Input({ required: true })
  set FormControls(value: FormControls) {
    Object.entries(value).forEach((entry) => {
      this.defaultValue[entry[0]] = entry[1].value;
      this.allTrueValue[entry[0]] = true;
    });
    this.formGroup = new FormGroup(value);
  }

  @Input({ required: true })
  TitleField?: Field<T, string>;

  @Input({ required: true })
  Title?: string;

  @Input({ required: true })
  TPath: string = '';

  protected defaultValue: { [key: string]: any } = {};
  protected allTrueValue: { [key: string]: any } = {};
  protected selectedRow?: T;
  protected formGroup: FormGroup;
  protected Data: T[] = [];
  protected SortField?: Field<T, any>;
  protected reversed: boolean = false;
  protected _fields: Field<T, any>[] = [];
  protected width: number = 100;

  constructor(
    private readonly appService: ApiService<T>,
    readonly formBuilder: FormBuilder,
  ) {
    this.formGroup = new FormGroup({});
  }

  ngOnInit() {
    this.refreshDate();
  }

  setSort(field: Field<T, any>) {
    if (this.SortField != field) {
      this.SortField = field;
      this.reversed = false;
    } else if (this.reversed) {
      this.SortField = undefined;
    } else {
      this.reversed = true;
    }
  }

  refreshDate() {
    this.appService.GetAll(this.TPath).subscribe((data) => {
      if (data) this.Data = data;
    });
  }

  compare(a: T, b: T) {
    const ac = this.SortField?.Content(a);
    const bc = this.SortField?.Content(b);
    if (typeof ac == 'string') {
      if (ac < bc) return -1;
      if (ac > bc) return 1;
      return 0;
    }
    if (typeof ac == 'boolean') return ac ? 1 : -1;
    if (typeof ac == 'number') return ac - bc;
    return 0;
  }

  getSortedDate() {
    if (!this.SortField) {
      return this.Data;
    }
    return [...this.Data].sort(
      (a, b) => (this.reversed ? -1 : 1) * this.compare(a, b),
    );
  }

  async deleteItem(row: T) {
    this.selectedRow = row;
    const confirmed = await this.DeleteDialog?.open();
    if (confirmed) {
      this.appService.Delete(this.TPath, row.id).subscribe((res) => {
        if (res) this.refreshDate();
      });
    }
  }

  async openEditDialog(): Promise<boolean> {
    const confirmed = await this.EditDialog?.open();
    if (!confirmed) {
      return false;
    }
    if (this.formGroup.valid) {
      return true;
    }
    return await new Promise<boolean>((resolve) => {
      setTimeout(async () => {
        resolve(await this.openEditDialog());
      }, 1000);
    });
  }

  async newItem() {
    this.formGroup.setValue(this.defaultValue);
    this.selectedRow = undefined;
    if (await this.openEditDialog()) {
      this.appService
        .Post(this.TPath, this.formGroup.value as T)
        .subscribe((res) => {
          if (res) this.refreshDate();
        });
    }
  }

  async editItem(row: T) {
    const val: { [key: string]: any } = {};
    Object.entries(row).forEach((entry) => {
      if (this.allTrueValue[entry[0]]) {
        val[entry[0]] = entry[1];
      }
    });
    this.formGroup.setValue(val);
    this.selectedRow = row;
    if (await this.openEditDialog()) {
      this.appService
        .Put(this.TPath, this.formGroup.value, row.id)
        .subscribe((res) => {
          if (res) this.refreshDate();
        });
    }
  }

  viewItem(row: T) {
    this.selectedRow = row;
    this.ViewDialog?.open();
  }

  protected readonly BookComponent = BookComponent;
  protected readonly CategoryComponent = CategoryComponent;
}
