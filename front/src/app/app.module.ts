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
import { DialogComponent } from './controls/dialog/dialog.component';
import { CategoryComponent } from './pages/category/category.component';
import { ReactiveFormsModule } from '@angular/forms';
import {
  provideHttpClient,
  withInterceptorsFromDi,
} from '@angular/common/http';
import { ListViewComponent } from './controls/list-view/list-view.component';
import { ItemSelectorComponent } from './controls/item-selector/item-selector.component';
import { DatePickerComponent } from './controls/date-picker/date-picker.component';

@NgModule({
  declarations: [
    AppComponent,
    MenuComponent,
    BookComponent,
    MemberComponent,
    BorrowingComponent,
    PageComponent,
    DataTableComponent,
    DialogComponent,
    CategoryComponent,
    ListViewComponent,
    ItemSelectorComponent,
    DatePickerComponent,
  ],
  imports: [BrowserModule, NgbModule, AppRoutingModule, ReactiveFormsModule],
  providers: [provideHttpClient(withInterceptorsFromDi())],
  bootstrap: [AppComponent],
})
export class AppModule {}
