/** @type {import('./$types').PageServerLoad} */
export function load() {
    return {
	    gay: ""
    };
}

/** @type {import('./$types').Actions} */
export const actions = {
    search: async ({request}) => {
	    const data = await request.formData();
	    const nombre = data.get("firstname")
	    const apellido = data.get("lastname")
	    const catastro = data.get("catastro")
  	    const direccion_fisica = data.get("direccion_fisica")
	    const direccion_postal = data.get("direccion_postal")
		
	    let filters = []

	    if(nombre){
		filters.push(`CONTACT LIKE '%${nombre}%'`)	
	    }

	    if(apellido){
		filters.push(`CONTACT LIKE '%${apellido}%'`)
	    }

	    if(catastro){
		filters.push(`CATASTRO LIKE '%${catastro}%'`)
	    }

	    if(direccion_fisica){
		filters.push(`DIRECCION_FISICA LIKE '%${direccion_fisica}%'`)
	    }

	    if(direccion_postal){
		filters.push(`DIRECCION_POSTAL LIKE '%${direccion_postal}%'`)
	    }

	    let whereStatement = filters.join(' AND ')

	    const res = await fetch('http://localhost:3000/' + new URLSearchParams(whereStatement))
	    const result = await res.json()

	    return result
    },
};
