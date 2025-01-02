import { Component, Input } from '@angular/core';

@Component({
  selector: 'app-page',
  standalone: false,

  templateUrl: './page.component.html',
  styleUrl: './page.component.scss',
})
export class PageComponent {
  @Input({ required: true })
  Title?: string;
}
