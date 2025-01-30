import { FormControl } from '@angular/forms';

export interface Field<in T, out V> {
  Title: string;
  Visible: boolean;
  Content(input: T): V | undefined;
}

export type FormControls = { [key: string]: FormControl };
