import http from 'k6/http';
import {SharedArray} from 'k6/data';
import {Rate} from 'k6/metrics';
import {check} from 'k6';
import papaparse from 'https://jslib.k6.io/papaparse/5.1.1/index.js';
import {uuidv4} from 'https://jslib.k6.io/k6-utils/1.0.0/index.js';
import jsonpath from 'https://jslib.k6.io/jsonpath/1.0.2/index.js';

export let options = {
    discardResponseBodies: true,
    scenarios: {
        time_measurement_arm_selector: {
            executor: 'ramping-arrival-rate',
            exec: 'arm_selector',
            startRate: 20,
            preAllocatedVUs: 10,
            timeUnit: '1m',
            maxVUs: 100,
            stages: [
                { target: 800, duration: '10m' },
                { target: 0, duration: '1m' },
            ],
        }
    },
    thresholds: {
        http_req_failed: ['rate<0.05'],
        http_req_duration: ['p(99)<500'],
    }
};

//Metrics
const failRate = new Rate('errorRate');

export function metrics(resp)
{
    failRate.add(resp.status !== 200)
    check(resp, {"status was 200": (r) => r.status == 200})
}

const ENV = `${__ENV.ENV}`
const URL_ARM_SELECTOR = (ENV === 'docker')? 'http://host.docker.internal:8090' : 'http://localhost:8090'

const rewardInputs = new SharedArray("reward_inputs", function() {return papaparse.parse(open('reward_predictor_test_data.csv'), { header: true }).data;});

export function arm_selector() {

    let randomI = Math.floor(Math.random() * rewardInputs.length)
    let experimentId = rewardInputs[randomI]["experiment_id"]
    let contextAge = rewardInputs[randomI]["age"]
    let contextIsFemale = rewardInputs[randomI]["is_female"]

    const data = {
        age: contextAge,
        is_female: contextIsFemale
    };

    let response = http.post(`${URL_ARM_SELECTOR}/api/v1/arm/selection/${experimentId}`, JSON.stringify(data), {
        headers: {'Content-Type': 'application/json', 'accept': 'application/json'},
    });

    failRate.add(response.status !== 200 && response.status !== 404)
    check(response, {"status was ok": (r) => r.status === 200 || r.status === 404})
}