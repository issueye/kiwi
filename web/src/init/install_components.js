import BasePage from '~/components/base_page/index.vue';
import DynamicTable from '~/components/dynamic_table/index.vue'
import ElIconPicker from '~/components/IconPicker.vue'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

import VueClipboard from 'vue3-clipboard'


export function installComponents(app) {
    app.component('BasePage', BasePage);
    app.component('DTable', DynamicTable);
    app.component('IconPicker', ElIconPicker);
}

export function installElementIcons(app) {
    app.config.globalProperties.$icons = []
    for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
        app.config.globalProperties.$icons.push(key)
        app.component(key, component)
    }
}

export function installColipboard(app) {
    app.use(VueClipboard, {
        autoSetContainer: true, // add this line
        appendToBody: true // add this line
    })
}