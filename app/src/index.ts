import { getCounter, incrementCounter } from './counter/counter';

const result = document.getElementById('result')!

// load the current value to the view.
getCounter().then(res => {
    result.innerHTML = res.value.toString();
});

// send an increment request and then render the new value to the view.
document.getElementById('incrementBtn')!.onclick = function () {
    incrementCounter().then(res => {
        result.innerHTML = res.value.toString();
    });
}
