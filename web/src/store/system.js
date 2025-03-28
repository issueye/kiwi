import { defineStore } from 'pinia'
import { reactive, ref } from 'vue'

import { apiGetRoleList } from "~/api/role";

export const useSystemStore = defineStore(
    'system',
    () => {
        /**
         * 登录用户
         */
        const loginUser = reactive({

        })

        /**
         * 角色列表
         */
        const roleList = ref([]);

        /**
         * 获取角色列表
         */
        const getRoleList = async (params) => {
            const data = await apiGetRoleList(params);
            roleList.value = data.list;
        }

        return {
            // 属性
            loginUser,
            roleList,

            // 方法
            getRoleList
        }
    },
    { 
        persist: true
    }
)