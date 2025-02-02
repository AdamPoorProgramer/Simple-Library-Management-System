import { Component } from '@angular/core';
import {
  Field,
  FormControls,
} from '../../controls/data-table/data-table.model';
import { Book } from './book.model';
import { AppComponent } from '../../app.component';
import { FormControl, Validators } from '@angular/forms';
import { CategoryComponent } from '../category/category.component';

@Component({
  selector: 'app-book',
  standalone: false,
  templateUrl: './book.component.html',
  styleUrl: './book.component.scss',
})
export class BookComponent {
  static readonly BookFormControls: FormControls = {
    title: new FormControl('', Validators.required),
    author: new FormControl('', Validators.required),
    publisher: new FormControl('', Validators.required),
    year: new FormControl(new Date().getFullYear(), Validators.required),
    genre: new FormControl('', Validators.required),
    categories: new FormControl([]),
  };
  static readonly BookTitleField: Field<Book, string> = {
    Title: 'Title',
    Visible: true,
    Content: (input: Book) => input.title,
  };
  static readonly BookFields: Field<Book, any>[] = [
    ...AppComponent.ItemFields,
    BookComponent.BookTitleField,
    {
      Title: 'Author',
      Visible: true,
      Content: (input: Book) => input.author,
    },
    {
      Title: 'Publisher',
      Visible: true,
      Content: (input: Book) => input.publisher,
    },
    {
      Title: 'Publish year',
      Visible: false,
      Content: (input: Book) => input.year,
    },
    {
      Title: 'Genre',
      Visible: true,
      Content: (input: Book) => input.genre,
    },
    {
      Title: 'Categories',
      Visible: false,
      Content: (input: Book) => input.categories,
    },
  ];

  protected readonly BookComponent = BookComponent;
  protected readonly CategoryComponent = CategoryComponent;
}
