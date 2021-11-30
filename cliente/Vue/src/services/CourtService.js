import { ApiServiceLaravel } from "./Api.laravel";


export const RentingsService = {

    get(id) {
      return ApiServiceLaravel.get("courts/", id);
    },
    create(params) {
      return ApiServiceLaravel.post("courts/", { court: params });
    },
    update(id, params) {
      return ApiServiceLaravel.update("courts/", id, { court: params });
    },
    destroy(id) {
      return ApiServiceLaravel.delete(`courts/${id}`);
    }
  };
  