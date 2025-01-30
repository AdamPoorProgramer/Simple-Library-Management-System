export class ValueChangedEvent<T> {
  public readonly oldValue?: T;
  public readonly newValue?: T;
  public constructor(oldV?: T, newV?: T) {
    this.oldValue = oldV;
    this.newValue = newV;
  }
}
