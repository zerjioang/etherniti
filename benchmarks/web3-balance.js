const { performance } = require('perf_hooks');
const Web3 = require('web3');

/**
 * Measure the time of execution with Date timestamp of a synchronous task
 *
 * @param {function} toMeasure
 * @param {int} repeatTimes
 * @return {Object}
 */
function TimeBenchmark(toMeasure,repeatTimes){
    if(typeof(repeatTimes) != "number"){
        repeatTimes = 1;
    }

    if(typeof(toMeasure) === "function"){
        var start_time = new Date().getTime();
        for (var i = 0; i < repeatTimes; ++i) {
            toMeasure.call();
        }
        var end_time = new Date().getTime();
    }
    
    return {
        start:start_time,
        end: end_time,
        estimatedMilliseconds: end_time - start_time
    };
}


/**
 * Measure the time of execution in milliseconds of a synchronous task
 *
 * @param {function} toMeasure
 * @param {int} repeatTimes
 * @return {Object}
 */
function StandardBenchmark(toMeasure,repeatTimes){
    if(typeof(repeatTimes) != "number"){
        repeatTimes = 1;
    }
    
    if(typeof(toMeasure) === "function"){
        var start_status = performance.now();
        var total_taken = 0;
        for(var i = 0;i < repeatTimes;i++){
            var startTimeSubtask = performance.now();
            toMeasure.call();
            var endTimeSubtask = performance.now();
            total_taken += (endTimeSubtask -startTimeSubtask);
        }
        var final_status = performance.now();
    }

    return { totalMilliseconds: (final_status - start_status), averageMillisecondsPerTask: total_taken / repeatTimes };
}

let getBalanceFunction = function(){
	web3 = new Web3(new Web3.providers.HttpProvider('HTTP://127.0.0.1:7545'));
	//console.log('Getting Ethereum address info.....');
	var addr = ('0x3DEB1894DC2d3e1B4a073f520e516C2DF6f45B88');
	//console.log('Address:', addr);
	web3.eth.getBalance(addr, function (error, result) {
		if (!error)
			console.log('Ether:', web3.utils.fromWei(result,'ether')); 
		else
			console.log('Huston we have a promblem: ', error);
		
	})
}

//const iterations = [1,5,10,50,100,500,1000,2000]
const iterations = [1]
for (var i = 0; i < iterations.length ; i++) {
	let TestResult = new StandardBenchmark(getBalanceFunction,iterations[i]);
	console.log("Iterations ", iterations[i], TestResult);
}