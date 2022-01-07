import Apilaravel from "@/services/Api.laravel";

export default{

    get() {
      console.log("ENTRA GET APILARAVEL");
      return Apilaravel().get(`users/`);
    },
    add(params) {
      console.log("apiservice add");
      console.log(params);
      return Apilaravel().post(`users/`, params);
    },
    update(id, params) {
      console.log("apiservice update");
      console.log(id)
      console.log(params)
      return Apilaravel().put(`users/${id}`, params); 
    },
    delete(id) {
      console.log("apiservice delete");
      console.log(id);
      return Apilaravel().delete(`users/${id}`);
    }
  };