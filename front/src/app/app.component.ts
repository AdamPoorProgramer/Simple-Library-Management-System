import { Component } from '@angular/core';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { Field } from './controls/data-table/data-table.model';
import { IItem } from './app.model';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  standalone: false,
  styleUrl: './app.component.scss',
})
export class AppComponent {
  static readonly ItemFields: Field<IItem, any>[] = [
    {
      Title: 'ID',
      Visible: false,
      Content: (input: IItem) => input.id,
    },
    {
      Title: 'CreationTime',
      Visible: false,
      Content: (input: IItem) => (input.created_at?.split('T') ?? [''])[0],
    },
    {
      Title: 'LastUpdateTime',
      Visible: false,
      Content: (input: IItem) => (input.updated_at?.split('T') ?? [''])[0],
    },
  ];

  constructor(private modalService: NgbModal) {}

  public open(modal: any): void {
    this.modalService.open(modal);
  }
}
