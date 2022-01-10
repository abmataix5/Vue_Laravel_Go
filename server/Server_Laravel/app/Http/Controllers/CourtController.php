<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

use App\Http\Resources\CourtResource;
use App\Http\Resources\CourtCollection;
use App\Http\Requests\StoreCourtRequest; 

use App\Traits\ApiResponseTrait;
use App\Repositories\CourtRepository;

use App\Models\Court;

class CourtController extends Controller
{

    use ApiResponseTrait;
    public $courtRepository;

    public function __construct(CourtRepository $courtRepository)
    {
        $this->courtRepository = $courtRepository;
        // $this->not_found = Response::HTTP_NOT_FOUND;
    }

    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index(Request $request)
    {
        $data= $this->courtRepository->all();;
       
        ///DEBUG
        $out = new \Symfony\Component\Console\Output\ConsoleOutput();
        $out->writeln("--------------- LIST COURTS -------------------");
        $out->writeln($data);
        /////DEBUG
       
        return self::apiResponseSuccess(new CourtCollection($data), 'Listado Pistas');

    }

    /**
     * Show the form for creating a new resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function create()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function store(StoreCourtRequest $request)
    {
        ///DEBUG
        $out = new \Symfony\Component\Console\Output\ConsoleOutput();
        $out->writeln("---------------STORE COURT CONTROLLER-------------------");
    
        $data = $this->courtRepository->create($request->all());
        return self::apiResponseSuccess($data, 'Pista creada correctamente');
    }

    /**
     * Display the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function show($id)
    {
        // 
    }

    /**
     * Show the form for editing the specified resource.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function edit($id)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     *
     * @param  \Illuminate\Http\Request  $request
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function update(Request $request, $id)
    {
          $data = $this->courtRepository->update($id,$request->all());

          if(is_null($data)){
              return self::apiResponseNotFound($id, 'Pista no encontrada');
          } else {
              return self::apiResponseSuccess($id, 'Pista actualizada correctamente');
          }
    }

    /**
     * Remove the specified resource from storage.
     *
     * @param  int  $id
     * @return \Illuminate\Http\Response
     */
    public function destroy($id)
    {
      $data = $this->courtRepository->delete($id);
      if(is_null($data)){
        return self::apiResponseNotFound($id, 'Pista no encontrada');
    } else {
        return self::apiResponseSuccess($id, 'Pista actualizada correctamente');
    }
      
      // if(Court::where('id', $id)->exists()) {
      //       $court = Court::find($id);
      //       $court->delete();
      //       return self::apiResponseSuccess($id, 'Pista eliminada correctamente');
      //   } else {
      //       return self::apiResponseNotFound($id, 'Pista no encontrado');
      //   }

    }
}
