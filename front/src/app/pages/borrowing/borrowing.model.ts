import { IItem } from '../../app.model';
import { Member } from '../member/member.model';
import { Book } from '../book/book.model';

export interface Borrowing extends IItem {
  book_id: number;
  member_id: number;
  member: Member;
  book: Book;
  borrow_date: string;
  return_date?: string | null;
  returned: boolean;
}
