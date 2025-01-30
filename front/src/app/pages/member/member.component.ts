import { Component } from '@angular/core';
import {
  Field,
  FormControls,
} from '../../controls/data-table/data-table.model';
import { Member } from './member.model';
import { AppComponent } from '../../app.component';
import { FormControl, Validators } from '@angular/forms';

@Component({
  selector: 'app-member',
  standalone: false,
  templateUrl: './member.component.html',
  styleUrl: './member.component.scss',
})
export class MemberComponent {
  static readonly MemberFormControls: FormControls = {
    first_name: new FormControl('', Validators.required),
    last_name: new FormControl('', Validators.required),
    phone_number: new FormControl('', [
      Validators.required,
      Validators.pattern(/^09[0-9]{9}$/),
    ]),
    email: new FormControl('', [Validators.required, Validators.email]),
    join_date: new FormControl('', Validators.required),
  };
  static readonly MemberTitleField: Field<Member, string> = {
    Title: 'Name',
    Visible: true,
    Content: (input: Member) => {
      return `${input.first_name} ${input.last_name}`;
    },
  };
  static readonly MemberFields: Field<Member, any>[] = [
    ...AppComponent.ItemFields,
    MemberComponent.MemberTitleField,
    {
      Title: 'Phone number',
      Visible: false,
      Content: (input: Member) => input.phone_number,
    },
    {
      Title: 'Email',
      Visible: true,
      Content: (input: Member) => input.email,
    },
    {
      Title: 'Join date',
      Visible: true,
      Content: (input: Member) => (input.join_date?.split('T') ?? [''])[0],
    },
  ];

  protected readonly MemberComponent = MemberComponent;
}
