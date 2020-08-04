"use strict";
var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
Object.defineProperty(exports, "__esModule", { value: true });
var Equipment_1 = require("./Equipment");
var CompositeEquipment = /** @class */ (function (_super) {
    __extends(CompositeEquipment, _super);
    function CompositeEquipment(name) {
        var _this = _super.call(this, name) || this;
        _this.equipment = [];
        return _this;
    }
    CompositeEquipment.prototype.add = function (newEquipment) {
        this.equipment.push(newEquipment);
    };
    CompositeEquipment.prototype.remove = function (equipmentToBeDeleted) {
        this.equipment = this.equipment.filter(function (equipmentComponent) {
            return equipmentComponent === equipmentToBeDeleted;
        });
    };
    Object.defineProperty(CompositeEquipment.prototype, "wattage", {
        get: function () {
            var totalWattage = 0;
            for (var _i = 0, _a = this.equipment; _i < _a.length; _i++) {
                var equipmentComponent = _a[_i];
                totalWattage += equipmentComponent.wattage;
            }
            return totalWattage;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(CompositeEquipment.prototype, "netPrice", {
        get: function () {
            var totalNetPrice = 0;
            for (var _i = 0, _a = this.equipment; _i < _a.length; _i++) {
                var equipmentComponent = _a[_i];
                totalNetPrice += equipmentComponent.netPrice;
            }
            return totalNetPrice;
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(CompositeEquipment.prototype, "discountPrice", {
        get: function () {
            var totalDiscountPrice = 0;
            for (var _i = 0, _a = this.equipment; _i < _a.length; _i++) {
                var equipmentComponent = _a[_i];
                totalDiscountPrice += equipmentComponent.discountPrice;
            }
            return totalDiscountPrice;
        },
        enumerable: true,
        configurable: true
    });
    return CompositeEquipment;
}(Equipment_1.Equipment));
exports.CompositeEquipment = CompositeEquipment;
