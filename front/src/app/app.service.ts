import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { IItem } from './app.model';
import { catchError, map, Observable, of } from 'rxjs';

@Injectable({
  providedIn: 'root',
})
export class ApiService<T extends IItem> {
  constructor(private readonly client: HttpClient) {}
  private readonly basePath = '/api';

  getPath(tPath: string): string {
    return `${this.basePath}/${tPath}/`;
  }

  public Get(tPath: string, id?: number): Observable<T | null> {
    return id
      ? this.client
          .get<T>(this.getPath(tPath), {
            params: { id },
          })
          .pipe(catchError(() => of(null)))
      : of(null);
  }

  public GetAll(tPath: string): Observable<T[] | null> {
    return this.client
      .get<T[]>(this.getPath(tPath))
      .pipe(catchError(() => of(null)));
  }

  public Post(tPath: string, input: T): Observable<boolean> {
    return this.client.post(this.getPath(tPath), JSON.stringify(input)).pipe(
      map(() => true),
      catchError(() => of(false)),
    );
  }

  public Put(tPath: string, input: T, id?: number): Observable<boolean> {
    return id
      ? this.client
          .put(this.getPath(tPath), JSON.stringify(input), {
            params: { id },
          })
          .pipe(
            map(() => true),
            catchError(() => of(false)),
          )
      : this.Post(tPath, input);
  }

  public Delete(tPath: string, id?: number): Observable<boolean> {
    return id
      ? this.client
          .delete(this.getPath(tPath), {
            params: { id },
          })
          .pipe(
            map(() => true),
            catchError(() => of(false)),
          )
      : of(false);
  }
}
