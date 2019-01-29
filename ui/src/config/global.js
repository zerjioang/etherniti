//global functions for vue
window.log = function (o){
	if(!process.env.PRODUCTION){
		console.log(o);
	}
}
window.error = function (o){
	if(!process.env.PRODUCTION){
		console.error(o);
	}
}

export default {
}