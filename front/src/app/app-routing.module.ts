import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { MenuComponent } from './pages/menu/menu.component';
import { MemberComponent } from './pages/member/member.component';
import { BookComponent } from './pages/book/book.component';
import { BorrowingComponent } from './pages/borrowing/borrowing.component';
import { CategoryComponent } from './pages/category/category.component';

const routes: Routes = [
  { path: '', component: MenuComponent },
  { path: 'member', component: MemberComponent },
  { path: 'book', component: BookComponent },
  { path: 'category', component: CategoryComponent },
  { path: 'borrowing', component: BorrowingComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
