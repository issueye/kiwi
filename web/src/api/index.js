import axios from '~/axios'


export function apiGetHomeCount() {
    return axios.get("/api/v1/admin/homeCount")
}

export function apiGetTrafficStatistics(type) {
    return axios.get(`/api/v1/admin/traffic?type=${type}`)
}

export function apiGetPortTraffic() {
    return axios.post(`/api/v1/proxy/port_statistics`)
}