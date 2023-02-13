import http from 'k6/http';
import { check } from 'k6';
export const options = {
 scenarios: {
  constant_request_rate: {
   executor: 'constant-arrival-rate',
   rate: 30,
   timeUnit: '1s',
   duration: '1m',
   preAllocatedVUs: 20,
   maxVUs: 100,
  },
 },
};
export default function () {
 const res =http.get('https://google.com');
 const checkRes = check(res, {
  'status is 200': (r) => r.status === 200,
 });
}
