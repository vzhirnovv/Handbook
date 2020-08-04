import { Watt, Currency } from "./types";

export abstract class Equipment {
  constructor(private _name: string) {}

  public add(equipment: Equipment): void {
    throw new Error(
      "Equipment can not be added: This equipment does not have any components."
    );
  }

  public remove(equipment: Equipment): void {
    throw new Error(
      "Equipment can not be deleted: This equipment does not have any components."
    );
  }

  public get name(): string {
    return this._name;
  }

  public abstract get wattage(): Watt;
  public abstract get netPrice(): Currency;
  public abstract get discountPrice(): Currency;
}
