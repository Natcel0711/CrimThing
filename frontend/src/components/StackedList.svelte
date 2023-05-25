<script lang="ts">
	export let CONTACT
	export let MUNICIPIO
	export let DIRECCION_FISICA
	export let DIRECCION_POSTAL
	export let NUMERO_CATASTRO
	
	function ShowDetails(prop){
		console.log(prop)
	}
</script>
  <li on:click={ShowDetails({
  Name: CONTACT,
  City:MUNICIPIO,
  NumeroCatastro: NUMERO_CATASTRO,
  })} class="flex justify-between gap-x-6 py-5">
    <div class="flex gap-x-4">
      <div class="min-w-0 flex-auto">
        <p class="text-sm font-semibold leading-6 text-gray-900">{CONTACT}</p>
        <p class="mt-1 truncate text-xs leading-5 text-gray-500">{MUNICIPIO}</p>
      </div>
    </div>
    <div class="hidden sm:flex sm:flex-col sm:items-end">
      <p class="text-sm leading-6 text-gray-900">{DIRECCION_FISICA} - {DIRECCION_POSTAL}</p>
      <div class="mt-1 flex items-center gap-x-1.5">
        <p class="text-xs leading-5 text-gray-500">{NUMERO_CATASTRO}</p>
      </div>
    </div>
  </li>

