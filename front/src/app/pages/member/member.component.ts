import { Component } from '@angular/core';

@Component({
  selector: 'app-member',
  standalone: false,
  templateUrl: './member.component.html',
  styleUrl: './member.component.scss',
})
export class MemberComponent {
  fields = [
    {
      Title: 'Value1111111111111111111111',
      Content(v: string) {
        return v;
      },
    },
  ];

  data = ['hi', 'bye'];
}
