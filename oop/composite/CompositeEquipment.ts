import { Equipment } from "./Equipment";
import { Watt, Currency } from "./types";

export class CompositeEquipment extends Equipment {
  private equipment: Equipment[] = [];

  constructor(name: string) {
    super(name);
  }

  public add(newEquipment: Equipment): void {
    this.equipment.push(newEquipment);
  }

  public remove(equipmentToBeDeleted: Equipment): void {
    this.equipment = this.equipment.filter(
      (equipmentComponent: Equipment) =>
        equipmentComponent === equipmentToBeDeleted
    );
  }

  public get wattage(): Watt {
    let totalWattage = 0;

    for (const equipmentComponent of this.equipment) {
      totalWattage += equipmentComponent.wattage;
    }

    return totalWattage;
  }

  public get netPrice(): Currency {
    let totalNetPrice = 0;

    for (const equipmentComponent of this.equipment) {
      totalNetPrice += equipmentComponent.netPrice;
    }

    return totalNetPrice;
  }

  public get discountPrice(): Currency {
    let totalDiscountPrice = 0;

    for (const equipmentComponent of this.equipment) {
      totalDiscountPrice += equipmentComponent.discountPrice;
    }
    return totalDiscountPrice;
  }
}
