import axios from '~/axios';


// 获取用户列表
export function apiGetList(data) {
    return axios.post('/api/v1/project/robot/list', data);
}

// apiAddUser
// 添加用户
export function apiAddData(data) {
    return axios.post('/api/v1/project/robot', data);
}

// apiUpdateUser
// 更新用户
export function apiUpdateData(data) {
    return axios.put('/api/v1/project/robot', data);
}

// apiDeleteUser
// 删除用户
export function apiDeleteById(id) {
    return axios.delete(`/api/v1/project/robot/${id}`);
}