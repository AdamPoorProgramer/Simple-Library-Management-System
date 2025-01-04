export interface Api<T> {
  GetAll(): T[];
  Get(id: number): T;
  Post(data: T): boolean;
  Put(id: number, data: T): boolean;
  Delete(id: number): boolean;
}

export interface Field<T> {
  Title: string;
  Content(input: T) : any;
}
