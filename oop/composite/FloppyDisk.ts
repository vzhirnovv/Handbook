import { Equipment } from "./Equipment";
import { Watt, Currency } from "./types";

export class FloppyDisk extends Equipment {
  constructor(name: string, private _size: number) {
    super(name);
  }

  public get wattage(): Watt {
    return (this._size * 10) as Watt;
  }

  public get netPrice(): Currency {
    return (this._size * 100) as Currency;
  }

  public get discountPrice(): Currency {
    return (this._size * 90) as Currency;
  }
}
