<script>
	import data from "../lib/dummydata/properties.json"
	import StackedList from "../components/StackedList.svelte";
	
	let propertyFilter = {
		Name: "",
		Municipio: "",
		NumeroCatastro: "",
	}

	let searching = false
</script>
{#if !searching }
<div class="container mx-auto px-4">
  <div>Bitch</div>
</div>
{:else}

<div class="grid grid-cols-2 gap-4">
  <div class="m-2">
  <a href="/" class="rounded-md bg-gray-800 px-3.5 py-2.5 text-sm font-semibold text-white shadow-sm hover:bg-gray-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-gray-600">Back</a>
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
{#each data.features as property}
	<StackedList 
	CONTACT={property.attributes.CONTACT} 
	NUMERO_CATASTRO={property.attributes.NUM_CATASTRO} 
	DIRECCION_POSTAL={property.attributes.DIRECCION_POSTAL} 
	DIRECCION_FISICA={property.attributes.DIRECCION_FISICA} 
	MUNICIPIO={property.attributes.MUNICIPIO} />
{/each}
</ul>
{/if}
