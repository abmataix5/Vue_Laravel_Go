<template>
 <div>

  <section class="jumbotron text-center">
    <div class="container">
      <h1>Bienvenido a MyPadelClub</h1>
      <p class="lead text-muted">Inicia sesión en nuestra APP para poder administrar tu club de padel!</p>
      <p>
              <router-link href="#" class="btn btn-primary my-2" to="/login" >Iniciar sesion con GO</router-link>
        <router-link href="#" class="btn btn-secondary my-2" to="/loginLaravel">Iniciar sesion Laravel</router-link>
      </p>
    </div>
  </section>

  <div class="album py-5 bg-light">
    <div class="container">

      <div class="row">
        <div class="col-md-4">
          <div class="card mb-4 shadow-sm">
            <svg class="bd-placeholder-img card-img-top" width="100%" height="225" xmlns="http://www.w3.org/2000/svg" preserveAspectRatio="xMidYMid slice" focusable="false" role="img" aria-label="Placeholder: Thumbnail"><image xlink:href="https://reservatupadel.es/wp-content/uploads/2021/06/cropped-Logo-Reserva-Tu-Padel-e1623667908373.png"></image></svg>
            <div class="card-body">
              <p class="card-text">Consulta tus reservas!</p>
              <div class="d-flex justify-content-between align-items-center">
               <p>
              <router-link href="#" class="btn btn-primary my-2" to="/rent" >Ir a reservas</router-link>
       
               </p>
     
            
              </div>
            </div>
          </div>
        </div>
   

    <div class="col-md-4">
          <div class="card mb-4 shadow-sm">
            <svg class="bd-placeholder-img card-img-top" width="100%" height="225" xmlns="http://www.w3.org/2000/svg" preserveAspectRatio="xMidYMid slice" focusable="false" role="img" aria-label="Placeholder: Thumbnail"><image xlink:href="https://www.padelzaragoza.es/img/personalizacion/padelzgz/Paginas/flyers-socios-1web.jpg"></image></svg>
            <div class="card-body">
              <p class="card-text">Socios</p>
              <div class="d-flex justify-content-between align-items-center">

                <p>
              <router-link href="#" class="btn btn-primary my-2" to="/partner" >Mas información</router-link>
       
               </p>
     
              </div>
            </div>
          </div>
        </div>


         <div class="col-md-4">
          <div class="card mb-4 shadow-sm">
            <svg class="bd-placeholder-img card-img-top" width="100%" height="225" xmlns="http://www.w3.org/2000/svg" preserveAspectRatio="xMidYMid slice" focusable="false" role="img" aria-label="Placeholder: Thumbnail"><image xlink:href="https://indusmetaltorres.es/wp-content/uploads/2018/04/Pista-padel-Carteya-verde-6005-1024x704.jpg" width="350"></image></svg>
            <div class="card-body">
              <p class="card-text">Consultar pistas disponibles</p>
              <div class="d-flex justify-content-between align-items-center">
                  <p>
              <router-link href="#" class="btn btn-primary my-2" to="/court" >Consultar pistas</router-link>
       
               </p>
              
              </div>
            </div>
          </div>
        </div>

         <div class="col-md-4">
          <div class="card mb-4 shadow-sm">
            <svg class="bd-placeholder-img card-img-top" width="100%" height="225" xmlns="http://www.w3.org/2000/svg" preserveAspectRatio="xMidYMid slice" focusable="false" role="img" aria-label="Placeholder: Thumbnail"><image xlink:href="https://www.streetpadel.com/blog/wp-content/uploads/sreetpadel-clases-de-padel.jpg" width="350"></image></svg>
            <div class="card-body">
              <p class="card-text">Clases de padel</p>
              <div class="d-flex justify-content-between align-items-center">
               
                <p>

                 <router-link href="#" class="btn btn-primary my-2" to="/court" >Mas información</router-link>
       
               </p>

              </div>
            </div>
          </div>
        </div>

         <div class="col-md-4">
          <div class="card mb-4 shadow-sm">
            <svg class="bd-placeholder-img card-img-top" width="100%" height="225" xmlns="http://www.w3.org/2000/svg" preserveAspectRatio="xMidYMid slice" focusable="false" role="img" aria-label="Placeholder: Thumbnail"><image xlink:href="https://www.padeladdict.com/wp-content/uploads/2019/10/partes-pala-de-padel-1020x580.jpg" width="400"></image></svg>
            <div class="card-body">
              <p class="card-text">Consulta listado trabajadores</p>
              <div class="d-flex justify-content-between align-items-center">
                <p>

                 <router-link href="#" class="btn btn-primary my-2" to="/workers" >Consultar trabajadores</router-link>
       
               </p>
          
              </div>
            </div>
          </div>
        </div>

         <div class="col-md-4">
          <div class="card mb-4 shadow-sm">
            <svg class="bd-placeholder-img card-img-top" width="100%" height="225" xmlns="http://www.w3.org/2000/svg" preserveAspectRatio="xMidYMid slice" focusable="false" role="img" aria-label="Placeholder: Thumbnail"><image xlink:href="https://as01.epimg.net/opinion/imagenes/2020/06/09/blogs/1591692680_508101_1591693087_noticia_normal_recorte1.jpg" width="500"></image></svg>
            <div class="card-body">
              <p class="card-text">Nuestras instalaciones</p>
              <div class="d-flex justify-content-between align-items-center">
                
                 <p>

                   <router-link href="#" class="btn btn-primary my-2" to="/workers" >Mas información</router-link>
       
                 </p>
              
              </div>
            </div>
          </div>
        </div>
      

     
      
   

  
      
      </div>
    </div>
  </div>

</div>
</template>

<script>
  export default {
    data() {
      return {
        slide: 0,
        sliding: null
      }
    },
    methods: {
      onSlideStart(slide) {
        console.log(slide)
        this.sliding = true
      },
      onSlideEnd(slide) {
           console.log(slide)
        this.sliding = false
      }
    }
  }
</script>