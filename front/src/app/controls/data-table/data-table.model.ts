export interface Field<in T, out V> {
  Title: string;
  Content(input: T): V;
}
