import {
  Component,
  EventEmitter,
  Input,
  Output,
  AfterViewInit,
  ViewChild,
  ElementRef,
} from '@angular/core';
import * as bootstrap from 'bootstrap';

@Component({
  selector: 'app-dialog',
  standalone: false,
  templateUrl: './dialog.component.html',
  styleUrl: './dialog.component.scss',
})
export class DialogComponent implements AfterViewInit {
  @Input() Title?: string;
  @Input() Message?: string;
  @ViewChild('Modal') Modal?: ElementRef<HTMLDivElement>;

  ID?: number;

  private modalInstance: any;
  private resolve?(value: boolean): void;

  ngAfterViewInit() {
    this.modalInstance = new bootstrap.Modal(this.Modal?.nativeElement ?? '');
    this.Modal?.nativeElement?.addEventListener(
      'hidden.bs.modal',
      this.closed.bind(this),
    );
  }

  open(id: number) {
    this.ID = id;
    this.modalInstance.show();
    return new Promise<boolean>((resolve) => (this.resolve = resolve));
  }

  confirm() {
    if (this.resolve) {
      this.resolve(true);
      this.resolve = undefined;
    }
    this.modalInstance.hide();
  }

  closed() {
    if (this.resolve) {
      this.resolve(false);
      this.resolve = undefined;
    }
  }
}
