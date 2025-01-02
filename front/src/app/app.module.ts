import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { MenuComponent } from './pages/menu/menu.component';
import { BookComponent } from './pages/book/book.component';
import { MemberComponent } from './pages/member/member.component';
import { BorrowingComponent } from './pages/borrowing/borrowing.component';
import { PageComponent } from './controls/page/page.component';
import { DataTableComponent } from './controls/data-table/data-table.component';

@NgModule({
  declarations: [
    AppComponent,
    MenuComponent,
    BookComponent,
    MemberComponent,
    BorrowingComponent,
    PageComponent,
    DataTableComponent,
  ],
  imports: [BrowserModule, NgbModule, AppRoutingModule],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
