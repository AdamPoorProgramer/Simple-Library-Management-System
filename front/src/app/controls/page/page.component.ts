import { Component, Input, TemplateRef } from '@angular/core';

@Component({
  selector: 'app-page',
  standalone: false,
  templateUrl: './page.component.html',
  styleUrl: './page.component.scss',
})
export class PageComponent {
  @Input({ required: true })
  Title?: string;

  @Input({ required: false })
  TitleTemplate?: TemplateRef<any>;
}
