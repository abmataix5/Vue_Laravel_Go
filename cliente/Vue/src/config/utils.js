/* Guardamos la info del usuario que se logea */
 export const setStore = (name, content) => {
    if (!name) return
    if (typeof content !== 'string') {
      content = JSON.stringify(content)
    }
    return window.localStorage.setItem(name, content)
  }
  
/* Para coger info de localstorage */
  export const getStore = (name) => {
    if (!name) return
    return JSON.parse(window.localStorage.getItem(name))
  }
  
/* Borramos info de localstorage */

  export const removeItem = (name) => {
    if (!name) return
    return window.localStorage.removeItem(name)
  }
  

  
 