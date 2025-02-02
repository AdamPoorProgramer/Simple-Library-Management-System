import { IItem } from '../../app.model';

export interface Member extends IItem {
  first_name: string;
  last_name: string;
  phone_number: string;
  email: string;
  join_date: string;
}
