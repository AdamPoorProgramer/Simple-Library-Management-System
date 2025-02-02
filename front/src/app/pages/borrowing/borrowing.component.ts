import { Component } from '@angular/core';
import {
  Field,
  FormControls,
} from '../../controls/data-table/data-table.model';
import { Borrowing } from './borrowing.model';
import { AppComponent } from '../../app.component';
import { BookComponent } from '../book/book.component';
import { MemberComponent } from '../member/member.component';
import { Book } from '../book/book.model';
import { FormControl, Validators } from '@angular/forms';

@Component({
  selector: 'app-borrowing',
  standalone: false,
  templateUrl: './borrowing.component.html',
  styleUrl: './borrowing.component.scss',
})
export class BorrowingComponent {
  static readonly BorrowingFormControls: FormControls = {
    member_id: new FormControl(undefined, Validators.required),
    book_id: new FormControl(undefined, Validators.required),
    borrow_date: new FormControl('', Validators.required),
    returned: new FormControl(false),
    return_date: new FormControl('', Validators.required),
  };
  static readonly BorrowingTitleField: Field<Borrowing, string> = {
    Title: '',
    Visible: false,
    Content: (input: Borrowing) =>
      `${MemberComponent.MemberTitleField.Content(input.member)} ' +
      '${BookComponent.BookTitleField.Content(input.book)}`,
  };
  static readonly BorrowingFields: Field<Borrowing, any>[] = [
    ...AppComponent.ItemFields,
    {
      Title: 'Member',
      Visible: true,
      Content: (input: Borrowing) =>
        MemberComponent.MemberTitleField.Content(input.member),
    },
    {
      Title: 'Book',
      Visible: true,
      Content: (input: Borrowing) =>
        BookComponent.BookTitleField.Content(input.book),
    },
    {
      Title: 'Borrow date',
      Visible: false,
      Content: (input: Borrowing) => (input.borrow_date?.split('T') ?? [''])[0],
    },
    {
      Title: 'Is returned',
      Visible: true,
      Content: (input: Borrowing) =>
        input.returned ? 'Returned' : 'Not returned',
    },
    {
      Title: 'Return date',
      Visible: true,
      Content: (input: Borrowing) => (input.return_date?.split('T') ?? [''])[0],
    },
  ];

  protected readonly BorrowingComponent = BorrowingComponent;
  protected readonly MemberComponent = MemberComponent;
  protected readonly BookComponent = BookComponent;
}
