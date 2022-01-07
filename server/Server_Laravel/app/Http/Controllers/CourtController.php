<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

use App\Http\Resources\CourtResource;
use App\Http\Resources\CourtCollection;
use App\Http\Requests\StoreCourtRequest; 

use App\Traits\ApiResponseTrait;

use App\Models\Court;

class CourtController extends Controller
{

  use ApiResponseTrait;

    /**
     * Display a listing of the resource.
     *
     * @return \Illuminate\Http\Response
     */
    public function index(Request $request)
    {
      $res= Court::get();
        ///DEBUG
        $out = new \Symfony\Component\Console\Output\ConsoleOutput();
        $out->writeln("--------------- LIST COURT -------------------");
        $out->writeln($res);
        /////DEBUG

        // $page = $request->has('page') ? $request->get('page') : 1;
        // $limit = $request->has('limit') ? $request->get('limit') : 3;
        // return CourtResource::collection(Court::limit($limit)->offset(($page - 1) * $limit)->get());
       
        return self::apiResponseSuccess(new CourtCollection(Court::get()), 'Listado Pistas');
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
         /////DEBUG
         $court = new Court;
         $court->name = $request->name;
         $court->schedule = $request->schedule;
         $court->date = $request->date;
         $court->A_drive = $request->A_drive;
         $court->A_reves = $request->A_reves;
         $court->B_drive = $request->B_drive;
         $court->B_reves = $request->B_reves;
         $court->save();

         return self::apiResponseSuccess($data, 'Pista creada correctamente');

       /*  return CourtResource::make(Court::create($request->validated())); */

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
        
        if (Court::where('id', $id)->exists()) {
            $court = Court::find($id);
            $court->name = $request->name;
            $court->schedule = $request->schedule;
            $court->date = $request->date;
            $court->A_drive = $request->A_drive;
            $court->A_reves = $request->A_reves;
            $court->B_drive = $request->B_drive;
            $court->B_reves = $request->B_reves;
            $court->save();
            return self::apiResponseSuccess($id, 'Pista actualizada correctamente');

          } else {
            return self::apiResponseNotFound($id, 'Pista no encontrada');
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
        
        if(Court::where('id', $id)->exists()) {
            $court = Court::find($id);
            $court->delete();
            return self::apiResponseSuccess($id, 'Pista eliminada correctamente');
        } else {
            return self::apiResponseNotFound($id, 'Pista no encontrado');
        }

    }
}
