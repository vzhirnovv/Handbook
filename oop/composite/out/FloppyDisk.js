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
var FloppyDisk = /** @class */ (function (_super) {
    __extends(FloppyDisk, _super);
    function FloppyDisk(name, _size) {
        var _this = _super.call(this, name) || this;
        _this._size = _size;
        return _this;
    }
    Object.defineProperty(FloppyDisk.prototype, "wattage", {
        get: function () {
            return (this._size * 10);
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(FloppyDisk.prototype, "netPrice", {
        get: function () {
            return (this._size * 100);
        },
        enumerable: true,
        configurable: true
    });
    Object.defineProperty(FloppyDisk.prototype, "discountPrice", {
        get: function () {
            return (this._size * 90);
        },
        enumerable: true,
        configurable: true
    });
    return FloppyDisk;
}(Equipment_1.Equipment));
exports.FloppyDisk = FloppyDisk;
