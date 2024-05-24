import http from 'k6/http'

export let options = {
    // config k6 run -> don't need to use options in command line
    vus: 5,
    duration: '5s',
}

export default function () {
    // http.get('http://localhost:8000/hello')
    http.get('http://host.docker.internal:8000/products') // k6 in container
}

