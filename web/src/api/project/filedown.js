import axios from '~/axios';

/**
 * 获取文件列表
 * @param {*} data 
 * @returns 
 */
export function apiGetList(data) {
    return axios.post('/api/v1/project/package/list', data);
}

/**
 * 下载文件
 * @param {*} id 
 * @returns 
 */
export function apiDown(id) {
    return axios({
        url: `/api/v1/project/package/download/${id}`,
        method: 'get',
        responseType: 'blob'
    })
}

/**
 * 删除文件
 * @param {*} id 
 * @returns 
 */
export function apiDel(id) {
    return axios.delete(`/api/v1/project/package/${id}`)
}