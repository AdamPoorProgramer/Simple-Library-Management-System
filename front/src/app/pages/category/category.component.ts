import { Component } from '@angular/core';
import {
  Field,
  FormControls,
} from '../../controls/data-table/data-table.model';
import { Category } from './category.model';
import { AppComponent } from '../../app.component';
import { FormControl, Validators } from '@angular/forms';
import { BookComponent } from '../book/book.component';

@Component({
  selector: 'app-category',
  standalone: false,
  templateUrl: './category.component.html',
  styleUrl: './category.component.scss',
})
export class CategoryComponent {
  static readonly CategoryFormControls: FormControls = {
    name: new FormControl('', Validators.required),
    books: new FormControl([]),
  };
  static readonly CategoryTitleField: Field<Category, string> = {
    Title: 'Name',
    Visible: true,
    Content: (input: Category) => input.name,
  };
  static readonly CategoryFields: Field<Category, any>[] = [
    ...AppComponent.ItemFields,
    CategoryComponent.CategoryTitleField,
    {
      Title: 'Books',
      Visible: false,
      Content: (input: Category) => input.books,
    },
  ];

  protected readonly CategoryComponent = CategoryComponent;
  protected readonly BookComponent = BookComponent;
}
