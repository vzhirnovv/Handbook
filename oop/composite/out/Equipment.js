"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
var Equipment = /** @class */ (function () {
    function Equipment(_name) {
        this._name = _name;
    }
    Equipment.prototype.add = function (equipment) {
        throw new Error("Equipment can not be added: This equipment does not have any components.");
    };
    Equipment.prototype.remove = function (equipment) {
        throw new Error("Equipment can not be deleted: This equipment does not have any components.");
    };
    Object.defineProperty(Equipment.prototype, "name", {
        get: function () {
            return this._name;
        },
        enumerable: true,
        configurable: true
    });
    return Equipment;
}());
exports.Equipment = Equipment;
