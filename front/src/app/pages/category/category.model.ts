import { IItem } from '../../app.model';
import { Book } from '../book/book.model';

export interface Category extends IItem {
  name: string;
  books: Book[];
}
