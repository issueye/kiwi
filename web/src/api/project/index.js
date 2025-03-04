// src/api/user.js

import axios from '~/axios';


// 获取用户列表
export function apiGetList(data) {
  return axios.post('/api/v1/project/list', data);
}

// apiAddUser
// 添加用户
export function apiAddData(data) {
  return axios.post('/api/v1/project', data);
}

// apiUpdateUser
// 更新用户
export function apiUpdateData(data) {
  return axios.put('/api/v1/project', data);
}

// apiDeleteUser
// 删除用户
export function apiDeleteById(id) {
  return axios.delete(`/api/v1/project/${id}`);
}

// apiSyncData
// 同步数据
export function apiSyncData(id) {
  return axios.post(`/api/v1/project/sync/${id}`);
}

// apiSaveCompileConfig
// 保存编译配置
export function apiSaveCompileConfig(data) {
  return axios.post(`/api/v1/project/compile/config`, data);
}

// apiCompile
// 编译
export function apiCompile(id, version_id) {
  return axios.get(`/api/v1/project/compile/run/${id}?version_id=${version_id}`);
}

// apiGetByProjectId
// 根据项目id获取项目
export function apiGetByProjectId(id) {
  return axios.get(`/api/v1/project/compile/config/${id}`);
}
