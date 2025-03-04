import axios from '~/axios';

// apiGetRuleList
// 获取规则信息列表
export function apiGetRuleList(data) {
    return axios.post('/api/v1/rule/list', data);
}

// apiAddRule
// 添加规则信息
export function apiAddRule(data) {
    return axios.post('/api/v1/rule', data);
}

// apiUpdateRule
// 更新规则信息
export function apiUpdateRule(data) {
    return axios.put('/api/v1/rule', data);
}

// apiDeleteRule
// 删除规则信息
export function apiDeleteRule(id) {
    return axios.delete(`/api/v1/rule/${id}`);
}

export function apiModifyStatusRule(id) {
    return axios.put(`/api/v1/rule/status/${id}`)
}