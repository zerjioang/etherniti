//global functions for vue
window.log = function (o){
	if(process.env.NODE_ENV != 'production'){
		console.log(o);
	}
}
window.error = function (o){
	if(process.env.NODE_ENV != 'production'){
		console.error(o);
	}
}

export default {
}