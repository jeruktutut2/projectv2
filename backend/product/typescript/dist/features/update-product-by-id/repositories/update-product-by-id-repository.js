"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.UpdateProductByIdRepository = void 0;
class UpdateProductByIdRepository {
    static getById(poolConnection, id) {
        return __awaiter(this, void 0, void 0, function* () {
            const query = `SELECT id, user_id, name, description FROM products WHERE id = ?`;
            const result = yield poolConnection.execute(query, [id]);
            return result;
        });
    }
    static updateNameAndDescriptionById(poolConnection, product) {
        return __awaiter(this, void 0, void 0, function* () {
            const query = `UPDATE products SET name = ?, description = ? WHERE id = ?`;
            const result = yield poolConnection.execute(query, [product.name, product.description, product.id]);
            return result;
        });
    }
}
exports.UpdateProductByIdRepository = UpdateProductByIdRepository;
