// src/api/user.js

import axios from '~/axios';


// 获取用户列表
export function apiGetDictsList(data) {
  return axios.post('/api/v1/dict_mana/list', data);
}

// apiAddUser
// 添加用户
export function apiAddData(data) {
  return axios.post('/api/v1/dict_mana', data);
}

// apiUpdateUser
// 更新用户
export function apiUpdateData(data) {
  return axios.put('/api/v1/dict_mana', data);
}

// apiDeleteUser
// 删除用户
export function apiDeleteById(id) {
  return axios.delete(`/api/v1/dict_mana/${id}`);
}

export function apiSaveDetail(data) {
  return axios.post(`/api/v1/dict_mana/detail`, data);
}

export function apiDeleteDetail(id) {
  return axios.delete(`/api/v1/dict_mana/detail/${id}`);
}

export function apiDetails(data) {
  return axios.post(`/api/v1/dict_mana/details`, data);
}

export function apiGetDictDetails(id) {
  return axios.get(`/api/v1/dict_mana/${id}/details`);
}