// global variables for vue
window.API_ENDPOINT = process.env.API_SCHEME + "://" + process.env.API_DOMAIN + process.env.API_PATH;

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