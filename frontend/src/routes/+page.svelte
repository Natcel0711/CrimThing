<script>
	import data from "../lib/dummydata/properties.json"
</script>

{#each data.features as property}
	<h1>{property.attributes.CONTACT}</h1>
{/each}
