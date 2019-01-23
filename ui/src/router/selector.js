import soon from '@/router/soon';
import maintenance from '@/router/maintenance';
import full from '@/router/full';

function getRouter() {
	if(process.env.COOMING_SOON){
	  //get cooming soon only router
	  log("Running app in COOMING_SOON mode")
	  return soon;
	} else if (process.env.MAINTENANCE_MODE) {
		log("Running app in MAINTENANCE mode")
	  	return maintenance;
	}
	else {
	  log("Running app in FULL mode")
	  return full;
	}
}

export default {
	getRouter
}