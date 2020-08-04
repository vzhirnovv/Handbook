import { CompositeEquipment } from "./CompositeEquipment";
import { FloppyDisk } from "./FloppyDisk";

const smallFloppy = new FloppyDisk("Small Floppy Disk", 10);
console.log(smallFloppy);
const largeFloppy = new FloppyDisk("Large Floppy Disk", 1000);
console.log(largeFloppy);

const floppyReader = new CompositeEquipment("Floppy Reader");

floppyReader.add(smallFloppy);
floppyReader.add(largeFloppy);

const chassis = new CompositeEquipment("Chassis");
chassis.add(floppyReader);

console.log(chassis.netPrice, chassis.wattage, chassis.discountPrice);

floppyReader.remove(smallFloppy);

console.log(chassis.netPrice, chassis.wattage, chassis.discountPrice);
