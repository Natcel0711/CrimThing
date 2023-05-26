<script>
	import StackedList from "../components/StackedList.svelte";
	import { applyAction, enhance } from "$app/forms";

	let searching = false
	export let data;
	export let form;
	
	let loading = false
	let props
</script>
{#if !searching }
<div class="container mx-auto px-4 py-4">
<form action="?/search" use:enhance={() => {
searching = true
loading = true
return async ({result}) => {
 props = result.data
 loading = false
 await applyAction(result)
};
}}
method="post" class="w-full center max-w-lg">
  <div class="flex flex-wrap -mx-3 mb-6">
    <div class="w-full md:w-1/2 px-3 mb-6 md:mb-0">
      <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="FIRST_NAME">
        Nombre
      </label>
      <input class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="FirstnameInput" type="text" placeholder="Juan" name="firstname">
    </div>
    <div class="w-full md:w-1/2 px-3">
      <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="LAST_NAME">
        Apellido
      </label>
      <input class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="ApellidoInput" type="text" placeholder="Del Pueblo" name="lastname">
    </div>
  </div>
  <div class="flex flex-wrap -mx-3 mb-6">
    <div class="w-full px-3">
      <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="NUMERO_CATASTRO">
        Numero de catastro
      </label>
      <input class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="NumCatastroInput" placeholder="Numero de catastro" name="catastro">
    </div>
  </div>
  <div class="flex flex-wrap -mx-3 mb-6">
    <div class="w-full px-3">
      <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="DIRECCION_FISICA">
        Direcion fisica
      </label>
      <input class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="DFisicaInput" placeholder="Direccion fisica" name="direccion_fisica">
    </div>
  </div>
  <div class="flex flex-wrap -mx-3 mb-6">
    <div class="w-full px-3">
      <label class="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2" for="DIRECCION_POSTAL">
        Direcion postal
      </label>
      <input class="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white focus:border-gray-500" id="DFisicaInput" placeholder="Direccion Postal" name="direccion_postal">
    </div>
  </div>
  <button type="submit" class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded">
  Search
</button>
</form>
</div>
{:else}
{#if loading}
<div class="absolute right-1/2 bottom-1/2  transform translate-x-1/2 translate-y-1/2 ">
    <div class="border-t-transparent border-solid animate-spin  rounded-full border-blue-400 border-8 h-64 w-64"></div>
</div>
{:else}

<div class="grid grid-cols-2 gap-4">
  <div class="m-2">
    <a on:click={()=>{ searching = false }} href="/" class="rounded-md bg-gray-800 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-gray-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gray-600">Back</a>
  </div>
  <div>	
<label class="relative block">
	<span class="sr-only">Search</span>
  	<span class="absolute inset-y-0 left-0 flex items-center pl-2">
    		<svg class="h-5 w-5 fill-slate-300" viewBox="0 0 20 20"><!-- ... --></svg>
	</span>
	  <input class="placeholder:italic placeholder:text-slate-400 block bg-white w-full border border-slate-300 rounded-md py-2 pl-9 pr-3 shadow-sm focus:outline-none focus:border-sky-500 focus:ring-sky-500 focus:ring-1 sm:text-sm" placeholder="Search for anything..." type="text" name="search"/>
	</label>

  </div>
</div>

<ul role="list" class="divide-y divide-gray-100 m-5">
{#each props as property}
	<StackedList 
	CONTACT={property.attributes.CONTACT} 
	NUMERO_CATASTRO={property.attributes.NUM_CATASTRO} 
	DIRECCION_POSTAL={property.attributes.DIRECCION_POSTAL} 
	DIRECCION_FISICA={property.attributes.DIRECCION_FISICA} 
	MUNICIPIO={property.attributes.MUNICIPIO} />
{/each}
</ul>
{/if}
{/if}

<style>
.center {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  padding: 10px;
}
</style>
