const {res, status} = await req()

// if (status == 404){
// 	err = "not found"
// 	return
// }

if (status == 401){
	err = "unathorized"
	return
}

if (status == 200){
	// kek
}

err = res
