import { IItem } from '../../app.model';
import { Category } from '../category/category.model';

export interface Book extends IItem {
  title: string;
  author: string;
  publisher: string;
  year: number;
  genre: string;
  categories: Category[];
}
