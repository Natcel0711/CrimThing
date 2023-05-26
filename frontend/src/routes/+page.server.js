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
		
	    let filters = {
		firstname: "",
		lastname: "",
                catastro: "",
		direccion_fisica: "",
		direccion_postal: ""
	    }

	    if(nombre){
		filters.firstname = nombre.toString().trim()	
	    }

	    if(apellido){
		filters.lastname = apellido.toString().trim()
	    }

	    if(catastro){
		filters.catastro = catastro.toString().trim()
	    }

	    if(direccion_fisica){
		filters.direccion_fisica = direccion_fisica.toString().trim()
	    }

	    if(direccion_postal){
		filters.direccion_postal = direccion_postal.toString().trim()
	    }

	    const res = await fetch('http://localhost:3000/', {
		method: 'POST',
		body: JSON.stringify(filters)
	    })
	    const result = await res.json()
	    
	    return result.features
    },
};
